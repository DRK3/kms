/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

// swagger:route POST /kms/keystores/{keystoreID}/keys/{keyID}/easy crypto-box easyReq
//
// Encrypts (anonymously) a payload.
//
// Responses:
//        200: easyResp
//    default: errorResp
func (o *Operation) easyHandler(rw http.ResponseWriter, req *http.Request) { //nolint:funlen // TODO refactor
	ctx, span := o.tracer.Start(req.Context(), "easyHandler")
	defer span.End()

	o.logger.Debugf("handling request: %s", req.URL.String())

	start := time.Now()

	k, err := o.kmsService.ResolveKeystore(req.WithContext(ctx))
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, resolveKeystoreFailure, err)

		return
	}

	span.AddEvent("ResolveKeystore completed",
		trace.WithAttributes(label.String("duration", time.Since(start).String())))

	var request easyReq
	if ok := o.parseRequest(&request, rw, req); !ok {
		return
	}

	keystoreID := mux.Vars(req)[keystoreIDQueryParam]
	keyID := mux.Vars(req)[keyIDQueryParam]

	span.SetAttributes(label.String("keystoreID", keystoreID))
	span.SetAttributes(label.String("keyID", keyID))

	payload, err := base64.URLEncoding.DecodeString(request.Payload)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	nonce, err := base64.URLEncoding.DecodeString(request.Nonce)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	theirPub, err := base64.URLEncoding.DecodeString(request.TheirPub)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	cryptoBox, err := o.cryptoBoxCreator(k.KeyManager())
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, easyMessageFailure, err)

		return
	}

	cipherText, err := cryptoBox.Easy(payload, nonce, theirPub, keyID)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, easyMessageFailure, err)

		return
	}

	o.writeResponse(rw, easyResp{
		CipherText: base64.URLEncoding.EncodeToString(cipherText),
	})
}

// swagger:route POST /kms/keystores/{keystoreID}/easyopen crypto-box easyOpenReq
//
// Decrypts ("easy open") a payload.
//
// Responses:
//        200: easyOpenResp
//    default: errorResp
func (o *Operation) easyOpenHandler(rw http.ResponseWriter, req *http.Request) { //nolint:funlen // TODO refactor
	ctx, span := o.tracer.Start(req.Context(), "easyOpenHandler")
	defer span.End()

	o.logger.Debugf("handling request: %s", req.URL.String())

	start := time.Now()

	k, err := o.kmsService.ResolveKeystore(req.WithContext(ctx))
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, resolveKeystoreFailure, err)

		return
	}

	span.AddEvent("ResolveKeystore completed",
		trace.WithAttributes(label.String("duration", time.Since(start).String())))

	var request easyOpenReq
	if ok := o.parseRequest(&request, rw, req); !ok {
		return
	}

	keystoreID := mux.Vars(req)[keystoreIDQueryParam]

	span.SetAttributes(label.String("keystoreID", keystoreID))

	cipherText, err := base64.URLEncoding.DecodeString(request.CipherText)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	nonce, err := base64.URLEncoding.DecodeString(request.Nonce)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	theirPub, err := base64.URLEncoding.DecodeString(request.TheirPub)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	myPub, err := base64.URLEncoding.DecodeString(request.MyPub)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	cryptoBox, err := o.cryptoBoxCreator(k.KeyManager())
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, easyOpenMessageFailure, err)

		return
	}

	plainText, err := cryptoBox.EasyOpen(cipherText, nonce, theirPub, myPub)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, easyOpenMessageFailure, err)

		return
	}

	o.writeResponse(rw, easyOpenResp{
		PlainText: base64.URLEncoding.EncodeToString(plainText),
	})
}

// swagger:route POST /kms/keystores/{keystoreID}/sealopen crypto-box sealOpenReq
//
// Decrypts ("seal open") a payload.
//
// Responses:
//        200: sealOpenResp
//    default: errorResp
func (o *Operation) sealOpenHandler(rw http.ResponseWriter, req *http.Request) {
	ctx, span := o.tracer.Start(req.Context(), "sealOpenHandler")
	defer span.End()

	o.logger.Debugf("handling request: %s", req.URL.String())

	start := time.Now()

	k, err := o.kmsService.ResolveKeystore(req.WithContext(ctx))
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, resolveKeystoreFailure, err)

		return
	}

	span.AddEvent("ResolveKeystore completed",
		trace.WithAttributes(label.String("duration", time.Since(start).String())))

	var request sealOpenReq
	if ok := o.parseRequest(&request, rw, req); !ok {
		return
	}

	keystoreID := mux.Vars(req)[keystoreIDQueryParam]

	span.SetAttributes(label.String("keystoreID", keystoreID))

	cipherText, err := base64.URLEncoding.DecodeString(request.CipherText)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	myPub, err := base64.URLEncoding.DecodeString(request.MyPub)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusBadRequest, receivedBadRequest, err)

		return
	}

	cryptoBox, err := o.cryptoBoxCreator(k.KeyManager())
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, sealOpenPayloadFailure, err)

		return
	}

	plainText, err := cryptoBox.SealOpen(cipherText, myPub)
	if err != nil {
		o.writeErrorResponse(rw, http.StatusInternalServerError, sealOpenPayloadFailure, err)

		return
	}

	o.writeResponse(rw, sealOpenResp{
		PlainText: base64.URLEncoding.EncodeToString(plainText),
	})
}
