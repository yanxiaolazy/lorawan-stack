// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bleve

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	errCorruptedIndex = errors.DefineCorruption("corrupted_index", "corrupted index file")
)

// GetBrands lists available end device vendors.
func (s *bleveStore) GetBrands(req store.GetBrandsRequest) (*store.GetBrandsResponse, error) {
	documentTypeQuery := bleve.NewTermQuery(brandDocumentType)
	documentTypeQuery.SetField("Type")
	queries := []query.Query{documentTypeQuery}

	if q := req.Search; q != "" {
		queries = append(queries, bleve.NewQueryStringQuery(q))
	}
	if q := req.BrandID; q != "" {
		query := bleve.NewTermQuery(q)
		query.SetField("BrandID")
		queries = append(queries, query)
	}

	searchRequest := bleve.NewSearchRequest(bleve.NewConjunctionQuery(queries...))
	if req.Limit > 0 {
		searchRequest.Size = int(req.Limit)
	}
	if req.Page == 0 {
		req.Page = 1
	}
	searchRequest.From = int((req.Page - 1) * req.Limit)

	switch req.OrderBy {
	case "brand_id":
		searchRequest.SortBy([]string{"BrandID"})
	case "-brand_id":
		searchRequest.SortBy([]string{"-BrandID"})
	case "name":
		searchRequest.SortBy([]string{"BrandName"})
	case "-name":
		searchRequest.SortBy([]string{"-BrandName"})
	}

	result, err := s.index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	brands := make([]*ttnpb.EndDeviceBrand, 0, len(result.Hits))
	for _, hit := range result.Hits {
		brand, err := s.dr.getBrand(hit.ID)
		if err != nil {
			return nil, err
		}
		pb := &ttnpb.EndDeviceBrand{}
		if err := pb.SetFields(brand.Brand, req.Paths...); err != nil {
			return nil, err
		}
		brands = append(brands, pb)
	}
	return &store.GetBrandsResponse{
		Count:  uint32(len(result.Hits)),
		Total:  uint32(result.Total),
		Offset: uint32(searchRequest.From),
		Brands: brands,
	}, nil
}

// GetModels lists available end device definitions.
func (s *bleveStore) GetModels(req store.GetModelsRequest) (*store.GetModelsResponse, error) {
	documentTypeQuery := bleve.NewTermQuery(modelDocumentType)
	documentTypeQuery.SetField("Type")
	queries := []query.Query{documentTypeQuery}

	if q := req.Search; q != "" {
		queries = append(queries, bleve.NewQueryStringQuery(q))
	}
	if q := req.BrandID; q != "" {
		query := bleve.NewTermQuery(q)
		query.SetField("BrandID")
		queries = append(queries, query)
	}
	if q := req.ModelID; q != "" {
		query := bleve.NewTermQuery(q)
		query.SetField("ModelID")
		queries = append(queries, query)
	}

	searchRequest := bleve.NewSearchRequest(bleve.NewConjunctionQuery(queries...))
	if req.Limit > 0 {
		searchRequest.Size = int(req.Limit)
	}
	if req.Page == 0 {
		req.Page = 1
	}
	searchRequest.From = int((req.Page - 1) * req.Limit)

	switch req.OrderBy {
	case "brand_id":
		searchRequest.SortBy([]string{"BrandID"})
	case "-brand_id":
		searchRequest.SortBy([]string{"-BrandID"})
	case "model_id":
		searchRequest.SortBy([]string{"ModelID"})
	case "-model_id":
		searchRequest.SortBy([]string{"-ModelID"})
	}

	result, err := s.index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	models := make([]*ttnpb.EndDeviceModel, 0, len(result.Hits))
	for _, hit := range result.Hits {
		brandID, modelID := modelIDFromKey(hit.ID)
		model, err := s.dr.getModel(brandID, modelID)
		if err != nil {
			return nil, err
		}
		pb := &ttnpb.EndDeviceModel{}
		if err := pb.SetFields(model.Model, req.Paths...); err != nil {
			return nil, err
		}
		models = append(models, pb)
	}
	return &store.GetModelsResponse{
		Count:  uint32(len(result.Hits)),
		Total:  uint32(result.Total),
		Offset: uint32(searchRequest.From),
		Models: models,
	}, nil
}

// GetTemplate retrieves an end device template for an end device definition.
func (s *bleveStore) GetTemplate(ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	if ids == nil {
		return nil, errBrandNotFound.WithAttributes("brand_id", "")
	}
	model, err := s.dr.getModel(ids.BrandID, ids.ModelID)
	if err != nil {
		return nil, err
	}

	tmpl, ok := model.Templates[versionKey(*ids)]
	if !ok {
		return nil, errTemplateNotFound.WithAttributes(
			"brand_id", ids.BrandID,
			"model_id", ids.ModelID,
			"firmware_version", ids.FirmwareVersion,
			"band_id", ids.BandID,
		)
	}
	return tmpl, nil
}

func (s *bleveStore) getCodec(ids *ttnpb.EndDeviceVersionIdentifiers, f func(c storedCodec) *ttnpb.MessagePayloadFormatter) (*ttnpb.MessagePayloadFormatter, error) {
	if ids == nil {
		return nil, errBrandNotFound.WithAttributes("brand_id", "")
	}
	brand, err := s.dr.getBrand(ids.BrandID)
	if err != nil {
		return nil, err
	}
	codecID, ok := brand.CodecIDs[versionKey(*ids)]
	if !ok {
		return nil, errCodecNotFound.WithAttributes(
			"brand_id", ids.BrandID,
			"model_id", ids.ModelID,
			"firmware_version", ids.FirmwareVersion,
			"band_id", ids.BandID,
		)
	}
	codec, ok := brand.Codecs[codecID]
	if !ok {
		return nil, errCorruptedIndex.New()
	}
	formatter := f(codec)
	if formatter == nil {
		return nil, errCodecNotFound.WithAttributes(
			"brand_id", ids.BrandID,
			"model_id", ids.ModelID,
			"firmware_version", ids.FirmwareVersion,
			"band_id", ids.BandID,
		)
	}
	return formatter, nil
}

// GetUplinkDecoder retrieves the codec for decoding uplink messages.
func (s *bleveStore) GetUplinkDecoder(ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	return s.getCodec(ids, func(c storedCodec) *ttnpb.MessagePayloadFormatter { return c.UplinkDecoder })
}

// GetDownlinkDecoder retrieves the codec for decoding downlink messages.
func (s *bleveStore) GetDownlinkDecoder(ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	return s.getCodec(ids, func(c storedCodec) *ttnpb.MessagePayloadFormatter { return c.DownlinkDecoder })
}

// GetDownlinkEncoder retrieves the codec for encoding downlink messages.
func (s *bleveStore) GetDownlinkEncoder(ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	return s.getCodec(ids, func(c storedCodec) *ttnpb.MessagePayloadFormatter { return c.DownlinkEncoder })
}

// Close closes the store.
func (s *bleveStore) Close() error {
	return s.index.Close()
}
