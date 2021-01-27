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
	"bytes"
	"compress/gzip"
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/blevesearch/bleve"
	"go.thethings.network/lorawan-stack/v3/pkg/devicerepository/store"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
)

// bleveStore wraps a store.Store adding support for searching/sorting results using a bleve index.
type bleveStore struct {
	ctx context.Context

	index bleve.Index

	dr deviceRepository
}

// NewStore returns a new Device Repository store with indexing capabilities (using bleve).
func (c Config) NewStore(ctx context.Context) (store.Store, error) {
	wd, err := getWorkingDirectory(c.SearchPaths)
	if err != nil {
		return nil, err
	}
	fileName := filepath.Join(wd, packageFile)
	log.FromContext(ctx).WithField("file", fileName).Debug("Loading Device Repository package")
	compressed, err := ioutil.ReadFile(filepath.Join(wd, packageFile))
	if err != nil {
		return nil, err
	}
	z, err := gzip.NewReader(bytes.NewBuffer(compressed))
	if err != nil {
		return nil, err
	}
	defer z.Close()
	b, err := ioutil.ReadAll(z)
	if err != nil {
		return nil, err
	}
	s := &bleveStore{
		ctx: ctx,
		dr:  deviceRepository{},
	}
	if err := jsonpb.TTN().Unmarshal(b, &s.dr); err != nil {
		return nil, err
	}

	if s.index, err = s.dr.makeIndex(ctx); err != nil {
		return nil, err
	}
	return s, nil
}
