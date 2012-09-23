package pkcs11

// Generated with:
// grep CKM *t.h|grep '#define' | sed 's/^#define //' | awk ' { print $1 "=" $2 } '

const (
	CKM_RSA_PKCS_KEY_PAIR_GEN          = 0x00000000
	CKM_RSA_PKCS                       = 0x00000001
	CKM_RSA_9796                       = 0x00000002
	CKM_RSA_X_509                      = 0x00000003
	CKM_MD2_RSA_PKCS                   = 0x00000004
	CKM_MD5_RSA_PKCS                   = 0x00000005
	CKM_SHA1_RSA_PKCS                  = 0x00000006
	CKM_RIPEMD128_RSA_PKCS             = 0x00000007
	CKM_RIPEMD160_RSA_PKCS             = 0x00000008
	CKM_RSA_PKCS_OAEP                  = 0x00000009
	CKM_RSA_X9_31_KEY_PAIR_GEN         = 0x0000000A
	CKM_RSA_X9_31                      = 0x0000000B
	CKM_SHA1_RSA_X9_31                 = 0x0000000C
	CKM_RSA_PKCS_PSS                   = 0x0000000D
	CKM_SHA1_RSA_PKCS_PSS              = 0x0000000E
	CKM_DSA_KEY_PAIR_GEN               = 0x00000010
	CKM_DSA                            = 0x00000011
	CKM_DSA_SHA1                       = 0x00000012
	CKM_DH_PKCS_KEY_PAIR_GEN           = 0x00000020
	CKM_DH_PKCS_DERIVE                 = 0x00000021
	CKM_X9_42_DH_KEY_PAIR_GEN          = 0x00000030
	CKM_X9_42_DH_DERIVE                = 0x00000031
	CKM_X9_42_DH_HYBRID_DERIVE         = 0x00000032
	CKM_X9_42_MQV_DERIVE               = 0x00000033
	CKM_SHA256_RSA_PKCS                = 0x00000040
	CKM_SHA384_RSA_PKCS                = 0x00000041
	CKM_SHA512_RSA_PKCS                = 0x00000042
	CKM_SHA256_RSA_PKCS_PSS            = 0x00000043
	CKM_SHA384_RSA_PKCS_PSS            = 0x00000044
	CKM_SHA512_RSA_PKCS_PSS            = 0x00000045
	CKM_SHA224_RSA_PKCS                = 0x00000046
	CKM_SHA224_RSA_PKCS_PSS            = 0x00000047
	CKM_RC2_KEY_GEN                    = 0x00000100
	CKM_RC2_ECB                        = 0x00000101
	CKM_RC2_CBC                        = 0x00000102
	CKM_RC2_MAC                        = 0x00000103
	CKM_RC2_MAC_GENERAL                = 0x00000104
	CKM_RC2_CBC_PAD                    = 0x00000105
	CKM_RC4_KEY_GEN                    = 0x00000110
	CKM_RC4                            = 0x00000111
	CKM_DES_KEY_GEN                    = 0x00000120
	CKM_DES_ECB                        = 0x00000121
	CKM_DES_CBC                        = 0x00000122
	CKM_DES_MAC                        = 0x00000123
	CKM_DES_MAC_GENERAL                = 0x00000124
	CKM_DES_CBC_PAD                    = 0x00000125
	CKM_DES2_KEY_GEN                   = 0x00000130
	CKM_DES3_KEY_GEN                   = 0x00000131
	CKM_DES3_ECB                       = 0x00000132
	CKM_DES3_CBC                       = 0x00000133
	CKM_DES3_MAC                       = 0x00000134
	CKM_DES3_MAC_GENERAL               = 0x00000135
	CKM_DES3_CBC_PAD                   = 0x00000136
	CKM_CDMF_KEY_GEN                   = 0x00000140
	CKM_CDMF_ECB                       = 0x00000141
	CKM_CDMF_CBC                       = 0x00000142
	CKM_CDMF_MAC                       = 0x00000143
	CKM_CDMF_MAC_GENERAL               = 0x00000144
	CKM_CDMF_CBC_PAD                   = 0x00000145
	CKM_DES_OFB64                      = 0x00000150
	CKM_DES_OFB8                       = 0x00000151
	CKM_DES_CFB64                      = 0x00000152
	CKM_DES_CFB8                       = 0x00000153
	CKM_MD2                            = 0x00000200
	CKM_MD2_HMAC                       = 0x00000201
	CKM_MD2_HMAC_GENERAL               = 0x00000202
	CKM_MD5                            = 0x00000210
	CKM_MD5_HMAC                       = 0x00000211
	CKM_MD5_HMAC_GENERAL               = 0x00000212
	CKM_SHA_1                          = 0x00000220
	CKM_SHA_1_HMAC                     = 0x00000221
	CKM_SHA_1_HMAC_GENERAL             = 0x00000222
	CKM_RIPEMD128                      = 0x00000230
	CKM_RIPEMD128_HMAC                 = 0x00000231
	CKM_RIPEMD128_HMAC_GENERAL         = 0x00000232
	CKM_RIPEMD160                      = 0x00000240
	CKM_RIPEMD160_HMAC                 = 0x00000241
	CKM_RIPEMD160_HMAC_GENERAL         = 0x00000242
	CKM_SHA256                         = 0x00000250
	CKM_SHA256_HMAC                    = 0x00000251
	CKM_SHA256_HMAC_GENERAL            = 0x00000252
	CKM_SHA224                         = 0x00000255
	CKM_SHA224_HMAC                    = 0x00000256
	CKM_SHA224_HMAC_GENERAL            = 0x00000257
	CKM_SHA384                         = 0x00000260
	CKM_SHA384_HMAC                    = 0x00000261
	CKM_SHA384_HMAC_GENERAL            = 0x00000262
	CKM_SHA512                         = 0x00000270
	CKM_SHA512_HMAC                    = 0x00000271
	CKM_SHA512_HMAC_GENERAL            = 0x00000272
	CKM_SECURID_KEY_GEN                = 0x00000280
	CKM_SECURID                        = 0x00000282
	CKM_HOTP_KEY_GEN                   = 0x00000290
	CKM_HOTP                           = 0x00000291
	CKM_ACTI                           = 0x000002A0
	CKM_ACTI_KEY_GEN                   = 0x000002A1
	CKM_CAST_KEY_GEN                   = 0x00000300
	CKM_CAST_ECB                       = 0x00000301
	CKM_CAST_CBC                       = 0x00000302
	CKM_CAST_MAC                       = 0x00000303
	CKM_CAST_MAC_GENERAL               = 0x00000304
	CKM_CAST_CBC_PAD                   = 0x00000305
	CKM_CAST3_KEY_GEN                  = 0x00000310
	CKM_CAST3_ECB                      = 0x00000311
	CKM_CAST3_CBC                      = 0x00000312
	CKM_CAST3_MAC                      = 0x00000313
	CKM_CAST3_MAC_GENERAL              = 0x00000314
	CKM_CAST3_CBC_PAD                  = 0x00000315
	CKM_CAST5_KEY_GEN                  = 0x00000320
	CKM_CAST128_KEY_GEN                = 0x00000320
	CKM_CAST5_ECB                      = 0x00000321
	CKM_CAST128_ECB                    = 0x00000321
	CKM_CAST5_CBC                      = 0x00000322
	CKM_CAST128_CBC                    = 0x00000322
	CKM_CAST5_MAC                      = 0x00000323
	CKM_CAST128_MAC                    = 0x00000323
	CKM_CAST5_MAC_GENERAL              = 0x00000324
	CKM_CAST128_MAC_GENERAL            = 0x00000324
	CKM_CAST5_CBC_PAD                  = 0x00000325
	CKM_CAST128_CBC_PAD                = 0x00000325
	CKM_RC5_KEY_GEN                    = 0x00000330
	CKM_RC5_ECB                        = 0x00000331
	CKM_RC5_CBC                        = 0x00000332
	CKM_RC5_MAC                        = 0x00000333
	CKM_RC5_MAC_GENERAL                = 0x00000334
	CKM_RC5_CBC_PAD                    = 0x00000335
	CKM_IDEA_KEY_GEN                   = 0x00000340
	CKM_IDEA_ECB                       = 0x00000341
	CKM_IDEA_CBC                       = 0x00000342
	CKM_IDEA_MAC                       = 0x00000343
	CKM_IDEA_MAC_GENERAL               = 0x00000344
	CKM_IDEA_CBC_PAD                   = 0x00000345
	CKM_GENERIC_SECRET_KEY_GEN         = 0x00000350
	CKM_CONCATENATE_BASE_AND_KEY       = 0x00000360
	CKM_CONCATENATE_BASE_AND_DATA      = 0x00000362
	CKM_CONCATENATE_DATA_AND_BASE      = 0x00000363
	CKM_XOR_BASE_AND_DATA              = 0x00000364
	CKM_EXTRACT_KEY_FROM_KEY           = 0x00000365
	CKM_SSL3_PRE_MASTER_KEY_GEN        = 0x00000370
	CKM_SSL3_MASTER_KEY_DERIVE         = 0x00000371
	CKM_SSL3_KEY_AND_MAC_DERIVE        = 0x00000372
	CKM_SSL3_MASTER_KEY_DERIVE_DH      = 0x00000373
	CKM_TLS_PRE_MASTER_KEY_GEN         = 0x00000374
	CKM_TLS_MASTER_KEY_DERIVE          = 0x00000375
	CKM_TLS_KEY_AND_MAC_DERIVE         = 0x00000376
	CKM_TLS_MASTER_KEY_DERIVE_DH       = 0x00000377
	CKM_TLS_PRF                        = 0x00000378
	CKM_SSL3_MD5_MAC                   = 0x00000380
	CKM_SSL3_SHA1_MAC                  = 0x00000381
	CKM_MD5_KEY_DERIVATION             = 0x00000390
	CKM_MD2_KEY_DERIVATION             = 0x00000391
	CKM_SHA1_KEY_DERIVATION            = 0x00000392
	CKM_SHA256_KEY_DERIVATION          = 0x00000393
	CKM_SHA384_KEY_DERIVATION          = 0x00000394
	CKM_SHA512_KEY_DERIVATION          = 0x00000395
	CKM_SHA224_KEY_DERIVATION          = 0x00000396
	CKM_PBE_MD2_DES_CBC                = 0x000003A0
	CKM_PBE_MD5_DES_CBC                = 0x000003A1
	CKM_PBE_MD5_CAST_CBC               = 0x000003A2
	CKM_PBE_MD5_CAST3_CBC              = 0x000003A3
	CKM_PBE_MD5_CAST5_CBC              = 0x000003A4
	CKM_PBE_MD5_CAST128_CBC            = 0x000003A4
	CKM_PBE_SHA1_CAST5_CBC             = 0x000003A5
	CKM_PBE_SHA1_CAST128_CBC           = 0x000003A5
	CKM_PBE_SHA1_RC4_128               = 0x000003A6
	CKM_PBE_SHA1_RC4_40                = 0x000003A7
	CKM_PBE_SHA1_DES3_EDE_CBC          = 0x000003A8
	CKM_PBE_SHA1_DES2_EDE_CBC          = 0x000003A9
	CKM_PBE_SHA1_RC2_128_CBC           = 0x000003AA
	CKM_PBE_SHA1_RC2_40_CBC            = 0x000003AB
	CKM_PKCS5_PBKD2                    = 0x000003B0
	CKM_PBA_SHA1_WITH_SHA1_HMAC        = 0x000003C0
	CKM_WTLS_PRE_MASTER_KEY_GEN        = 0x000003D0
	CKM_WTLS_MASTER_KEY_DERIVE         = 0x000003D1
	CKM_WTLS_MASTER_KEY_DERIVE_DH_ECC  = 0x000003D2
	CKM_WTLS_PRF                       = 0x000003D3
	CKM_WTLS_SERVER_KEY_AND_MAC_DERIVE = 0x000003D4
	CKM_WTLS_CLIENT_KEY_AND_MAC_DERIVE = 0x000003D5
	CKM_KEY_WRAP_LYNKS                 = 0x00000400
	CKM_KEY_WRAP_SET_OAEP              = 0x00000401
	CKM_CMS_SIG                        = 0x00000500
	CKM_KIP_DERIVE                     = 0x00000510
	CKM_KIP_WRAP                       = 0x00000511
	CKM_KIP_MAC                        = 0x00000512
	CKM_CAMELLIA_KEY_GEN               = 0x00000550
	CKM_CAMELLIA_ECB                   = 0x00000551
	CKM_CAMELLIA_CBC                   = 0x00000552
	CKM_CAMELLIA_MAC                   = 0x00000553
	CKM_CAMELLIA_MAC_GENERAL           = 0x00000554
	CKM_CAMELLIA_CBC_PAD               = 0x00000555
	CKM_CAMELLIA_ECB_ENCRYPT_DATA      = 0x00000556
	CKM_CAMELLIA_CBC_ENCRYPT_DATA      = 0x00000557
	CKM_CAMELLIA_CTR                   = 0x00000558
	CKM_ARIA_KEY_GEN                   = 0x00000560
	CKM_ARIA_ECB                       = 0x00000561
	CKM_ARIA_CBC                       = 0x00000562
	CKM_ARIA_MAC                       = 0x00000563
	CKM_ARIA_MAC_GENERAL               = 0x00000564
	CKM_ARIA_CBC_PAD                   = 0x00000565
	CKM_ARIA_ECB_ENCRYPT_DATA          = 0x00000566
	CKM_ARIA_CBC_ENCRYPT_DATA          = 0x00000567
	CKM_SKIPJACK_KEY_GEN               = 0x00001000
	CKM_SKIPJACK_ECB64                 = 0x00001001
	CKM_SKIPJACK_CBC64                 = 0x00001002
	CKM_SKIPJACK_OFB64                 = 0x00001003
	CKM_SKIPJACK_CFB64                 = 0x00001004
	CKM_SKIPJACK_CFB32                 = 0x00001005
	CKM_SKIPJACK_CFB16                 = 0x00001006
	CKM_SKIPJACK_CFB8                  = 0x00001007
	CKM_SKIPJACK_WRAP                  = 0x00001008
	CKM_SKIPJACK_PRIVATE_WRAP          = 0x00001009
	CKM_SKIPJACK_RELAYX                = 0x0000100a
	CKM_KEA_KEY_PAIR_GEN               = 0x00001010
	CKM_KEA_KEY_DERIVE                 = 0x00001011
	CKM_FORTEZZA_TIMESTAMP             = 0x00001020
	CKM_BATON_KEY_GEN                  = 0x00001030
	CKM_BATON_ECB128                   = 0x00001031
	CKM_BATON_ECB96                    = 0x00001032
	CKM_BATON_CBC128                   = 0x00001033
	CKM_BATON_COUNTER                  = 0x00001034
	CKM_BATON_SHUFFLE                  = 0x00001035
	CKM_BATON_WRAP                     = 0x00001036
	CKM_ECDSA_KEY_PAIR_GEN             = 0x00001040
	CKM_EC_KEY_PAIR_GEN                = 0x00001040
	CKM_ECDSA                          = 0x00001041
	CKM_ECDSA_SHA1                     = 0x00001042
	CKM_ECDH1_DERIVE                   = 0x00001050
	CKM_ECDH1_COFACTOR_DERIVE          = 0x00001051
	CKM_ECMQV_DERIVE                   = 0x00001052
	CKM_JUNIPER_KEY_GEN                = 0x00001060
	CKM_JUNIPER_ECB128                 = 0x00001061
	CKM_JUNIPER_CBC128                 = 0x00001062
	CKM_JUNIPER_COUNTER                = 0x00001063
	CKM_JUNIPER_SHUFFLE                = 0x00001064
	CKM_JUNIPER_WRAP                   = 0x00001065
	CKM_FASTHASH                       = 0x00001070
	CKM_AES_KEY_GEN                    = 0x00001080
	CKM_AES_ECB                        = 0x00001081
	CKM_AES_CBC                        = 0x00001082
	CKM_AES_MAC                        = 0x00001083
	CKM_AES_MAC_GENERAL                = 0x00001084
	CKM_AES_CBC_PAD                    = 0x00001085
	CKM_AES_CTR                        = 0x00001086
	CKM_BLOWFISH_KEY_GEN               = 0x00001090
	CKM_BLOWFISH_CBC                   = 0x00001091
	CKM_TWOFISH_KEY_GEN                = 0x00001092
	CKM_TWOFISH_CBC                    = 0x00001093
	CKM_DES_ECB_ENCRYPT_DATA           = 0x00001100
	CKM_DES_CBC_ENCRYPT_DATA           = 0x00001101
	CKM_DES3_ECB_ENCRYPT_DATA          = 0x00001102
	CKM_DES3_CBC_ENCRYPT_DATA          = 0x00001103
	CKM_AES_ECB_ENCRYPT_DATA           = 0x00001104
	CKM_AES_CBC_ENCRYPT_DATA           = 0x00001105
	CKM_DSA_PARAMETER_GEN              = 0x00002000
	CKM_DH_PKCS_PARAMETER_GEN          = 0x00002001
	CKM_X9_42_DH_PARAMETER_GEN         = 0x00002002
	CKM_VENDOR_DEFINED                 = 0x80000000
)
