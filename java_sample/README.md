# Gpay sample code

Kết nối với hệ thống của Gpay
*(Connect to server of Gpay)*
## Bắt đầu(Start)

Controller.java	=> Ví dụ về chuyển tiền, tạo virtual account, khởi tạo giao dịch cổng
*(IbftController.java	=> Ibft example, VAN example, payment gateway init)*


GpayWebhookController.java	=> Nhận kết quả giao dịch VA, cổng.
*(GpayWebhookController.java	=> Receive VA transaction results, payment gateway translation)*

## Tạo RSA key (mở git bash) 
*Create RSA key(open git bash)*

```
openssl genrsa -out mykey.pem 2048

openssl req -new -x509 -key mykey.pem -out publickey.pem -days 3650

```
Dùng  **private_key.pem** để tạo chữ ký. Và gửi **publickey.pem** cho Gpay

*(Use **private_key.pem** to generate a signature. And send **publickey.pem** to Gpay.)*

