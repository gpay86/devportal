package com.gpay.gpaysample.GpaySampleApplication.enumeration;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public enum VaApiUri {
    TOKEN("/authentication/token/create"),
    CREATE("/virtual-account/create"),
    UPDATE("/virtual-account/update"),
    CLOSE("/virtual-account/close"),
    DETAIL("/virtual-account/detail"),
    REOPEN("/virtual-account/re-open"),
    VIETQR("/virtual-account/viet-qr/");
    private final String uri;
}

