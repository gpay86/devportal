package com.gpay.gpaysample.GpaySampleApplication.exception;


import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

import java.io.Serializable;

@Getter
@NoArgsConstructor
@AllArgsConstructor
public enum GpayErrorCode implements Serializable {

    INVALID_SIGNATURE("1009", "Invalid Signature"),
    SEVER_EXCEPTION("1010", "Server Exception"),
    INVALID_MERCHANT("1011", "Invalid Merchant"),
    PUBLIC_KEY_IS_NOT_CONFIG("1100", "public key is not config"),
    INVALID_REQUEST("1000", "Invalid request");


    private String code;
    private String message;

}
