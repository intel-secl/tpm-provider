/*
 * Copyright (C) 2020 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */
#include "tpm20linux.h"

static int change_auth(TSS2_SYS_CONTEXT* sys,
                       TPM2B_AUTH* newSecretKey, 
                       TPM2B_AUTH* oldSecretKey, 
                       const char* desc,
                       TPMI_RH_HIERARCHY_AUTH auth_handle) 
{

    TSS2_RC                 rval;
    TSS2L_SYS_AUTH_COMMAND  sessionsData = {0};

    if (newSecretKey == NULL) 
    {
        ERROR("The new secret key must be provided");
        return -1;
    }

    if (oldSecretKey == NULL)
    {
        ERROR("The old secret key must be provided");
        return -1;
    }

    sessionsData.count = 1;
    sessionsData.auths[0].sessionHandle = TPM2_RS_PW;
    memcpy(&sessionsData.auths[0].hmac, oldSecretKey, sizeof(TPM2B_AUTH));
    sessionsData.auths[0].sessionAttributes = 0;

    rval = Tss2_Sys_HierarchyChangeAuth(sys, auth_handle, &sessionsData, newSecretKey, 0);
    if (rval != TPM2_RC_SUCCESS) 
    {
        ERROR("Could not change hierarchy for %s: 0x%x", desc, rval);
    }

    return rval;
}

static int take_ownership(TSS2_SYS_CONTEXT* sys, 
                          TPM2B_AUTH* newSecretKey, 
                          TPM2B_AUTH* oldSecretKey)
{
    TSS2_RC rc;

    rc = change_auth(sys, newSecretKey, oldSecretKey, "Owner", TPM2_RH_OWNER);
    if(rc != TPM2_RC_SUCCESS)
    {
        return rc;
    }

    rc = change_auth(sys, newSecretKey, oldSecretKey, "Endorsement", TPM2_RH_ENDORSEMENT);
    if(rc != TPM2_RC_SUCCESS)
    {
        return rc;
    }

    rc = change_auth(sys, newSecretKey, oldSecretKey, "Lockout", TPM2_RH_LOCKOUT);
    if(rc != TPM2_RC_SUCCESS)
    {
        return rc;
    }

    return rc;
}


//-------------------------------------------------------------------------------------------------
// 'TakeOwnership' wraps three tpm2-tools commands: tpm2_takeownership, tpm2_createprimary 
// and tpm2_evictcontrol
//-------------------------------------------------------------------------------------------------
int TakeOwnership(const tpmCtx* ctx, 
                  const char* tpmSecretKey, 
                  size_t keyLength) 
{
    TSS2_RC rval = 0;
    // TPM2_HANDLE handle2048rsa = 0;
    TPM2B_AUTH newSecretKey = {0};
    TPM2B_AUTH oldSecretKey = {0}; // create an empty TPM2B_AUTH when provisioning the TPM
                                   // note:  We assume that this function is only called when the 
                                   // trust agent does not have a password configured AND WHEN
                                   // THE TPM IS CLEARED.  Changing the password is a feature
                                   // enhancement.


    rval = str2Tpm2bAuth(tpmSecretKey, keyLength, &newSecretKey);
    if(rval != 0)
    {
        ERROR("There was an error creating the new TPM2B_AUTH");
        return rval;
    }

    //
    // TakeOwnership of 'owner', 'endorsement' and 'lockout' similar to running...
    // tpm2_takeownership -o hex:c758af994ac60743fdf1ad5d8186ca216657f99f -e hex:c758af994ac60743fdf1ad5d8186ca216657f99f -l hex:c758af994ac60743fdf1ad5d8186ca216657f99f
    //
    rval = take_ownership(ctx->sys, &newSecretKey, &oldSecretKey);
    if(rval != TPM2_RC_SUCCESS)
    {
        return rval;
    }

    return TPM2_RC_SUCCESS;
}

//
// This function operates similar to the TpmLinuxV20.java implementation:  if 'change_auth' is successfull 
// when applying the same password for new/old keys, then consider the TPM owned with password 'secretKey'.
// 
// Returns zero (true) if the secretKey works against the TPM, -1 if not owned.  All other values non-zero
// values are error codes.
// 
int IsOwnedWithAuth(const tpmCtx* ctx, 
                    const char* ownerSecretKey, 
                    size_t keyLength)
{
    int rval;
    TPM2B_AUTH newSecretKey = {0};
    TPM2B_AUTH oldSecretKey = {0};

    rval = str2Tpm2bAuth(ownerSecretKey, keyLength, &newSecretKey);
    if(rval != 0)
    {
        ERROR("There was an error creating the new TPM2B_AUTH");
        return -2;
    }

    rval = str2Tpm2bAuth(ownerSecretKey, keyLength, &oldSecretKey);
    if(rval != 0)
    {
        ERROR("There was an error creating the old TPM2B_AUTH");
        return -2;
    }

    rval = change_auth(ctx->sys, &newSecretKey, &oldSecretKey, "Owner", TPM2_RH_OWNER);
    if (rval == 0x9a2) 
    {
        rval = -1;
    }

    return rval;    
}