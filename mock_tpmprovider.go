// +build unit_test

/*
 * Copyright (C) 2020 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
package tpmprovider

// #include "tpm.h"
import "C"

import (
	"github.com/stretchr/testify/mock"
)

type MockedTpmFactory struct {
	TpmProvider *MockedTpmProvider
}

func (mockedTpmFactory MockedTpmFactory) NewTpmProvider() (TpmProvider, error) {
	return mockedTpmFactory.TpmProvider, nil
}

//-------------------------------------------------------------------------------------------------
// Mocked TpmProvider interface
//-------------------------------------------------------------------------------------------------
type MockedTpmProvider struct {
	mock.Mock
}

func (mockedTpm MockedTpmProvider) Close() {
	_ = mockedTpm.Called()
	return
}

func (mockedTpm MockedTpmProvider) Version() C.TPM_VERSION {
	args := mockedTpm.Called()
	return args.Get(0).(C.TPM_VERSION)
}

func (mockedTpm MockedTpmProvider) TakeOwnership(ownerSecretKey string) error {
	args := mockedTpm.Called(ownerSecretKey)
	return args.Error(0)
}

func (mockedTpm MockedTpmProvider) IsOwnedWithAuth(ownerSecretKey string) (bool, error) {
	args := mockedTpm.Called(ownerSecretKey)
	return args.Bool(0), args.Error(1)
}

func (mockedTpm MockedTpmProvider) CreateAik(ownerSecretKey string, aikSecretKey string) error {
	args := mockedTpm.Called(ownerSecretKey, aikSecretKey)
	return args.Error(0)
}

func (mockedTpm MockedTpmProvider) GetAikBytes() ([]byte, error) {
	args := mockedTpm.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) GetAikName() ([]byte, error) {
	args := mockedTpm.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) ActivateCredential(ownerSecretKey string, aikSecretKey string, credentialBytes []byte, secretBytes []byte) ([]byte, error) {
	args := mockedTpm.Called(ownerSecretKey, aikSecretKey, credentialBytes, secretBytes)
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) GetTpmQuote(aikSecretKey string, nonce []byte, pcrBanks []string, pcrs []int) ([]byte, error) {
	args := mockedTpm.Called(aikSecretKey, nonce, pcrBanks, pcrs)
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) NvIndexExists(nvIndex uint32) (bool, error) {
	args := mockedTpm.Called(nvIndex)
	return args.Bool(0), args.Error(1)
}

func (mockedTpm MockedTpmProvider) NvDefine(ownerSecretKey string, nvIndex uint32, indexSize uint16) error {
	args := mockedTpm.Called(ownerSecretKey, nvIndex, indexSize)
	return args.Error(0)
}

func (mockedTpm MockedTpmProvider) NvRelease(ownerSecretKey string, nvIndex uint32) error {
	args := mockedTpm.Called(ownerSecretKey, nvIndex)
	return args.Error(0)
}

func (mockedTpm MockedTpmProvider) NvRead(ownerSecretKey string, nvIndex uint32) ([]byte, error) {
	args := mockedTpm.Called(ownerSecretKey, nvIndex)
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) NvWrite(ownerSecretKey string, nvIndex uint32, data []byte) error {
	args := mockedTpm.Called(ownerSecretKey, nvIndex, data)
	return args.Error(0)
}

func (mockedTpm MockedTpmProvider) CreatePrimaryHandle(ownerSecretKey string, handle uint32) error {
	args := mockedTpm.Called(ownerSecretKey, handle)
	return args.Error(0)
}

func (mockedTpm MockedTpmProvider) CreateSigningKey(signingSecretKey string, aikSecretKey string) (*CertifiedKey, error) {
	args := mockedTpm.Called(signingSecretKey, aikSecretKey)
	return args.Get(0).(*CertifiedKey), args.Error(1)
}

func (mockedTpm MockedTpmProvider) CreateBindingKey(bindingSecretKey string, aikSecretKey string) (*CertifiedKey, error) {
	args := mockedTpm.Called(bindingSecretKey, aikSecretKey)
	return args.Get(0).(*CertifiedKey), args.Error(1)
}

func (mockedTpm MockedTpmProvider) Unbind(certifiedKey *CertifiedKey, ownerSecretKey string, encryptedData []byte) ([]byte, error) {
	args := mockedTpm.Called(certifiedKey, ownerSecretKey, encryptedData)
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) Sign(certifiedKey *CertifiedKey, ownerSecretKey string, hash []byte) ([]byte, error) {
	args := mockedTpm.Called(certifiedKey, ownerSecretKey, hash)
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) PublicKeyExists(handle uint32) (bool, error) {
	args := mockedTpm.Called(handle)
	return args.Bool(0), args.Error(1)
}

func (mockedTpm MockedTpmProvider) ReadPublic(handle uint32) ([]byte, error) {
	args := mockedTpm.Called(handle)
	return args.Get(0).([]byte), args.Error(1)
}

func (mockedTpm MockedTpmProvider) CreateEk(ownerSecretKey string, handle uint32) error {
	args := mockedTpm.Called(ownerSecretKey, handle)
	return args.Error(0)
}

func (mockedTpm MockedTpmProvider) IsValidEk(ownerSecretKey string, handle uint32, nvIndex uint32) (bool, error) {
	args := mockedTpm.Called(ownerSecretKey, handle, nvIndex)
	return args.Bool(0), args.Error(1)
}

func (mockedTpm MockedTpmProvider) IsPcrBankActive(pcrBank string) (bool, error) {
	args := mockedTpm.Called(pcrBank)
	return args.Bool(0), args.Error(1)
}
