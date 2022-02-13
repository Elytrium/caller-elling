<img src="https://elytrium.net/src/img/elytrium.webp" alt="Elytrium" align="right">

# Phone verification module for Elling - Elytrium Billing
[![Join our Discord](https://img.shields.io/discord/775778822334709780.svg?logo=discord&label=Discord)](https://ely.su/discord)

Allows verifying client's phone number<br>
Available routing methods:
- ``/caller/get`` - Get the user's phone number
- ``/caller/call?number=79001231212&method=ucaller`` - Send the verification code to the user's phone number
- ``/caller/verify?code=1234`` - Verify the user's phone number



## Example verify method

(Put in the "caller" directory)

```yaml
name: UCaller
length: 4
verify-request:
    url: https://api.ucaller.ru/v1.0/initCall?service_id=12345&key=test&phone=79001231212&code={code}
    method: GET
    response-type: NONE
```

## Donation

Your donations are really appreciated. Donations wallets/links/cards:

- MasterCard Debit Card (Tinkoff Bank): ``5536 9140 0599 1975``
- Qiwi Wallet: ``PFORG`` or [this link](https://my.qiwi.com/form/Petr-YSpyiLt9c6)
- YooMoney Wallet: ``4100 1721 8467 044`` or [this link](https://yoomoney.ru/quickpay/shop-widget?writer=seller&targets=Donation&targets-hint=&default-sum=&button-text=11&payment-type-choice=on&mobile-payment-type-choice=on&hint=&successURL=&quickpay=shop&account=410017218467044)
- PayPal: ``ogurec332@mail.ru``
