/*
 * Copyright (C) 2019 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
 package tpmprovider

// #include "tpm.h"
import "C"

import (
	"fmt"
)

type MockTpm struct {
}

func NewMockTpm() (TpmProvider, error) {
	tpmMock := MockTpm {}
	return &tpmMock, nil
}

func (t *MockTpm) Close() {
}

func (t *MockTpm) Version() C.TPM_VERSION {
	return C.TPM_VERSION_UNKNOWN
}

func (t *MockTpm) CreateSigningKey(secretKey []byte, aikSecretKey []byte) (*CertifiedKey, error) {
	return nil, fmt.Errorf("MockTpm.CreateSigningKey is not implemented")
}

func (t *MockTpm) CreateBindingKey(secretKey []byte, aikSecretKey []byte) (*CertifiedKey, error) {
	return nil, fmt.Errorf("MockTpm.CreateBindingKey is not implemented")
}

func (t *MockTpm) Unbind(ck *CertifiedKey, keyAuth []byte, encData []byte) ([]byte, error) {
	return nil, nil
}

func (t *MockTpm) Sign(ck *CertifiedKey, keyAuth []byte, hashed []byte) ([]byte, error) {
	return nil, nil
}

func (t *MockTpm) TakeOwnership(newOwnerAuth string) error {
	return nil
}

func (t *MockTpm) IsOwnedWithAuth(ownerAuth string) (bool, error) {
	return true, nil
}

func (t *MockTpm) GetAikBytes(tpmSecretKey string) ([]byte, error) {
	return nil, fmt.Errorf("MockTpm.GetAikBytes is not implemented")
}

func (t *MockTpm) GetAikName(tpmSecretKey string) ([]byte, error) {
	return nil, fmt.Errorf("MockTpm.GetAikName is not implemented")
}

// func (t *MockTpm) IsAikPresent(tpmSecretKey string) (bool, error) {
// 	return false, fmt.Errorf("MockTpm.IsAikPresent is not implemented")
// }

func (t *MockTpm) CreateAik(tpmSecretKey string, aikSecretKey string) error {
	return fmt.Errorf("MockTpm.CreateAik is not implemented")
}

// func (t *MockTpm) FinalizeAik(aikSecretKey string) error {
// 	return fmt.Errorf("MockTpm.FinalizeAik is not implemented")
// }

func (t *MockTpm) GetTpmQuote(aikSecretKey string, nonce []byte, pcrBanks []string, pcrs []int)([]byte, error) {
	return nil, fmt.Errorf("MockTpm.GetTpmQuote is not implemented")
}

func (t *MockTpm) ActivateCredential(tpmSecretKey string, aikSecretKey string, credentialBytes []byte, secretBytes []byte) ([]byte, error) {
	return nil, fmt.Errorf("MockTpm.ActivateIdentity is not implemented")
}

// func (t *MockTpm) CreateEndorsementKey(tpmSecretKey string) error {
// 	return fmt.Errorf("MockTpm.CreateEndorsementKey is not implemented")
// }

func (t *MockTpm) NvIndexExists(handle uint32) (bool, error) {
	return false, fmt.Errorf("MockTpm.NvIndexExists is not implemented")
}

func (t *MockTpm) NvRelease(tpmOwnerSecretKey string, nvIndex uint32) error {
	return fmt.Errorf("MockTpm.NvRelease is not implemented")
}
func (t *MockTpm) NvDefine(tpmOwnerSecretKey string, nvIndex uint32, indexSize uint16) error {
	return fmt.Errorf("MockTpm.NvDefine is not implemented")
}

func (t *MockTpm) NvRead(tpmOwnerSecretKey string, handle uint32) ([]byte, error) {
	return nil, fmt.Errorf("MockTpm.NvRead is not implemented")
}

func (t *MockTpm) NvWrite(tpmOwnerSecretKey string, handle uint32, data []byte) error {
	return fmt.Errorf("MockTpm.NvWrite is not implemented")
}

func (tpm *MockTpm) CreatePrimaryHandle(ownerSecret []byte, handle uint32) error {
	return fmt.Errorf("MockTpm.CreatePrimaryHandle is not implemented")
}


func (tpm *MockTpm) PublicKeyExists(handle uint32) (bool, error) {
	return false, fmt.Errorf("MockTpm.PublicKeyExists is not implemented")
}

func (tpm *MockTpm) ReadPublic(secretKey string, handle uint32) ([]byte, error) {
	return nil, fmt.Errorf("MockTpm.ReadPublic is not implemented")
}


func (t *MockTpm) SetCredential(authHandle uint, ownerAuth []byte, credentialBlob []byte) error {
	return nil
}

func (t *MockTpm) GetCredential(authHandle uint) ([]byte, error) {
	var b[] byte
	b = make([]byte, 20, 20)
	return b, nil
}

func (t *MockTpm) GetAssetTag(authHandle uint) ([]byte, error) {
	var b[] byte
	b = make([]byte, 20, 20)
	return b, nil
}

func (t *MockTpm) GetAssetTagIndex() (uint, error) {
	return 0, nil
}