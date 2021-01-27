// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/index/scorch"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store/remote"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

const (
	packageFile = "lorawan-devices-db.json.gz"
)

type deviceRepository struct {
	Brands map[string]storedBrand
}

type storedBrand struct {
	Brand    *ttnpb.EndDeviceBrand
	Models   map[string]storedModel
	Codecs   map[string]storedCodec
	CodecIDs map[string]string // map version identifiers to codec IDs
}

type storedModel struct {
	Model     *ttnpb.EndDeviceModel
	Templates map[string]*ttnpb.EndDeviceTemplate
}

type storedCodec struct {
	UplinkDecoder,
	DownlinkEncoder,
	DownlinkDecoder *ttnpb.MessagePayloadFormatter
}

// Build generates a package file containing all definitions from the Device Repository.
func (c Config) Build(ctx context.Context, f fetch.Interface) error {
	s := remote.NewRemoteStore(f)

	wd, err := getWorkingDirectory(c.SearchPaths)
	if err != nil {
		return err
	}

	log.FromContext(ctx).Info("Building Device Repository package")

	brands, err := s.GetBrands(store.GetBrandsRequest{
		Paths: ttnpb.EndDeviceBrandFieldPathsNested,
	})
	if err != nil {
		return err
	}

	d := deviceRepository{
		Brands: make(map[string]storedBrand, len(brands.Brands)),
	}

	for _, brandPB := range brands.Brands {
		models, err := s.GetModels(store.GetModelsRequest{
			Paths:   ttnpb.EndDeviceModelFieldPathsNested,
			BrandID: brandPB.BrandID,
		})
		if err != nil {
			if errors.IsNotFound(err) {
				// Skip vendors without any models
				continue
			} else {
				return err
			}
		}

		brand := storedBrand{
			Brand:    brandPB,
			Models:   make(map[string]storedModel, len(models.Models)),
			Codecs:   make(map[string]storedCodec),
			CodecIDs: make(map[string]string),
		}
		for _, modelPB := range models.Models {
			model := storedModel{
				Model:     modelPB,
				Templates: make(map[string]*ttnpb.EndDeviceTemplate),
			}
			for _, fwVersion := range modelPB.GetFirmwareVersions() {
				for bandID, profile := range fwVersion.Profiles {
					version := ttnpb.EndDeviceVersionIdentifiers{
						BrandID:         modelPB.BrandID,
						ModelID:         modelPB.ModelID,
						FirmwareVersion: fwVersion.Version,
						BandID:          bandID,
					}
					tmpl, err := s.GetTemplate(&version)
					if err != nil {
						return err
					}

					key := versionKey(version)
					model.Templates[key] = tmpl

					if codecID := profile.CodecID; codecID != "" {
						brand.CodecIDs[key] = codecID
						if _, ok := brand.Codecs[codecID]; !ok {
							codec := storedCodec{}
							codec.UplinkDecoder, _ = s.GetUplinkDecoder(&version)
							codec.DownlinkDecoder, _ = s.GetDownlinkDecoder(&version)
							codec.DownlinkEncoder, _ = s.GetDownlinkEncoder(&version)
							brand.Codecs[codecID] = codec
						}
					}

				}
			}

			brand.Models[model.Model.ModelID] = model
		}
		d.Brands[brand.Brand.BrandID] = brand
	}

	b, err := jsonpb.TTN().Marshal(d)
	if err != nil {
		return err
	}
	fout, err := os.Create(filepath.Join(wd, packageFile))
	if err != nil {
		return err
	}
	defer fout.Close()
	z := gzip.NewWriter(fout)
	defer z.Close()
	if _, err := z.Write(b); err != nil {
		return err
	}
	return nil
}

func versionKey(version ttnpb.EndDeviceVersionIdentifiers) string {
	return fmt.Sprintf("%s:%s:%s:%s", version.BrandID, version.ModelID, version.FirmwareVersion, version.BandID)
}

func modelIDKey(brandID, modelID string) string {
	return fmt.Sprintf("%s:%s", brandID, modelID)
}

func modelIDFromKey(key string) (brandID string, modelID string) {
	parts := strings.SplitN(key, ":", 2)
	if len(parts) != 2 {
		panic(fmt.Sprintf("Invalid model ID key %s", key))
	}
	return parts[0], parts[1]
}

func (dr *deviceRepository) makeIndex(ctx context.Context) (bleve.Index, error) {

	mapping := bleve.NewIndexMapping()
	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = keyword.Name
	for _, keyword := range []string{"BrandID", "ModelID", "Type"} {
		mapping.DefaultMapping.AddFieldMappingsAt(keyword, keywordFieldMapping)
	}

	index, err := bleve.NewUsing("", mapping, scorch.Name, bleve.Config.DefaultMemKVStore, nil)
	if err != nil {
		return nil, err
	}

	batch := index.NewBatch()
	for brandID, brand := range dr.Brands {
		models := make([]*ttnpb.EndDeviceModel, 0, len(brand.Models))
		for modelID, model := range brand.Models {
			modelPB := model.Model
			if err := batch.Index(modelIDKey(brandID, modelID), indexableModel{
				Brand:   brand.Brand,
				Model:   model.Model,
				BrandID: brandID,
				ModelID: modelID,
				Type:    modelDocumentType,
			}); err != nil {
				return nil, err
			}
			models = append(models, modelPB)
		}
		if err := batch.Index(brandID, indexableBrand{
			Type:      brandDocumentType,
			Brand:     brand.Brand,
			Models:    models,
			BrandID:   brandID,
			BrandName: brand.Brand.Name,
		}); err != nil {
			return nil, err
		}
	}

	if err := index.Batch(batch); err != nil {
		return nil, err
	}
	return index, nil
}

func (dr *deviceRepository) getBrand(brandID string) (storedBrand, error) {
	brand, ok := dr.Brands[brandID]
	if !ok {
		return storedBrand{}, errBrandNotFound.WithAttributes("brand_id", brandID)
	}
	return brand, nil
}

func (dr *deviceRepository) getModel(brandID, modelID string) (storedModel, error) {
	brand, err := dr.getBrand(brandID)
	if err != nil {
		return storedModel{}, err
	}
	model, ok := brand.Models[modelID]
	if !ok {
		return storedModel{}, errModelNotFound.WithAttributes("brand_id", brandID, "model_id", modelID)
	}
	return model, nil
}
