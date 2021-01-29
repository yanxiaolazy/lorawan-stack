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
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

func newPopulatedDeviceRepo(numBrands, numModels, numCodecs, numTemplates int) deviceRepository {
	repo := deviceRepository{
		Brands: make(map[string]storedBrand, numBrands),
	}
	r := test.Randy

	for len(repo.Brands) < numBrands {
		brandPB := ttnpb.NewPopulatedEndDeviceBrand(r, true)
		if _, ok := repo.Brands[brandPB.BrandID]; ok {
			continue
		}
		if brandPB.BrandID == "" {
			continue
		}

		brand := storedBrand{
			Brand:    brandPB,
			Models:   make(map[string]storedModel, numModels),
			Codecs:   make(map[string]storedCodec, numCodecs),
			CodecIDs: make(map[string]string, numCodecs),
		}
		repo.Brands[brandPB.BrandID] = brand

		for i := 0; i < numCodecs; i++ {
			codecID := fmt.Sprintf("codec-%d", i)
			brand.Codecs[codecID] = storedCodec{
				UplinkDecoder:   ttnpb.NewPopulatedMessagePayloadFormatter(r, true),
				DownlinkDecoder: ttnpb.NewPopulatedMessagePayloadFormatter(r, true),
				DownlinkEncoder: ttnpb.NewPopulatedMessagePayloadFormatter(r, true),
			}
			brand.CodecIDs[codecID] = codecID
		}

		for len(brand.Models) < numModels {
			modelPB := ttnpb.NewPopulatedEndDeviceModel(r, true)

			if modelPB.ModelID == "" {
				continue
			}
			if _, ok := brand.Models[modelPB.ModelID]; ok {
				continue
			}
			model := storedModel{
				Model:     modelPB,
				Templates: make(map[string]*ttnpb.EndDeviceTemplate, numTemplates),
			}

			for i := 0; i < numTemplates; i++ {
				tmpl := ttnpb.NewPopulatedEndDeviceTemplate(r, true)
				tmpl.EndDevice.MACState = nil
				tmpl.EndDevice.PendingMACState = nil
				tmpl.EndDevice.Session = nil
				tmpl.EndDevice.PendingSession = nil
				tmpl.EndDevice.ProvisioningData = nil

				model.Templates[fmt.Sprintf("tmpl-%d", i)] = tmpl
			}

			brand.Models[modelPB.ModelID] = model
		}
	}

	return repo
}
