package com.gpay.gpaysample.GpaySampleApplication.enumeration;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public enum InternationalPaymentApiUri {
    INIT("/intlpayment/order/create"),
    TOKEN("/authentication/token/create"),
    DETAIL("/intlpayment/order/query");
    private final String uri;
}

