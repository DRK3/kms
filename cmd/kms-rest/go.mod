// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/kms/cmd/kms-rest

go 1.15

require (
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/containerd/continuity v0.0.0-20200928162600-f2cc35102c2a // indirect
	github.com/gorilla/mux v1.8.0
	github.com/hyperledger/aries-framework-go v0.1.6-0.20210306003819-79bbc2abaa5f
	github.com/hyperledger/aries-framework-go-ext/component/storage/couchdb v0.0.0-20210305233053-d3d22c802c12
	github.com/hyperledger/aries-framework-go-ext/component/vdr/sidetree v0.0.0-20210305233053-d3d22c802c12 // indirect
	github.com/hyperledger/aries-framework-go-ext/component/vdr/trustbloc v0.0.0-20210305233053-d3d22c802c12
	github.com/hyperledger/aries-framework-go/component/storageutil v0.0.0-20210306003819-79bbc2abaa5f
	github.com/hyperledger/aries-framework-go/spi v0.0.0-20210306003819-79bbc2abaa5f
	github.com/kilic/bls12-381 v0.1.0 // indirect
	github.com/magefile/mage v1.11.0 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/pquerna/cachecontrol v0.0.0-20201205024021-ac21108117ac // indirect
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.8.0 // indirect
	github.com/spf13/cobra v1.1.2
	github.com/stretchr/testify v1.7.0
	github.com/trustbloc/edge-core v0.1.6-0.20210305000733-14a89fe44ae8
	github.com/trustbloc/kms v0.0.0-00010101000000-000000000000
	github.com/trustbloc/sidetree-core-go v0.6.0 // indirect
	go.opentelemetry.io/otel v0.16.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.16.0
	go.opentelemetry.io/otel/sdk v0.16.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
)

replace github.com/trustbloc/kms => ../..
