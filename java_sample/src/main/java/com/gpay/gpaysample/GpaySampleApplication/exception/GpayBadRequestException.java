package com.gpay.gpaysample.GpaySampleApplication.exception;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class GpayBadRequestException extends Exception {
    private String code;
    private String message;


    public GpayBadRequestException(GpayErrorCode gpayErrorCode) {
        this.code = gpayErrorCode.getCode();
        this.message = gpayErrorCode.getMessage();
    }
}
