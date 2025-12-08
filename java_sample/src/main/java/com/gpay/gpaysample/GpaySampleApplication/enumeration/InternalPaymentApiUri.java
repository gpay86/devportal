package com.gpay.gpaysample.GpaySampleApplication.enumeration;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public enum InternalPaymentApiUri {
    INIT("/order/init"),
    TOKEN("/authentication/token/create"),
    DETAIL("/order/detail"),
    REMOVE_CARD_TOKEN("/order/remove-token"),
    BANK_LIST("/order/bank-list");
    private final String uri;
}

