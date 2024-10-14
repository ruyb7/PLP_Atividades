class StripePayment:
    def pay(self, amount):
        print(f"Pagamento de ${amount} via Stripe Processado.")

class PaypalPayment:
    def send_payment(self, amount):
        print(f"pagamento de ${amount} via Paypal Processado.")

class PaypalAdapter(StripePayment):
    def __init__ (self, paypal_payment):
        self.paypal_payment = paypal_payment

    def pay(self, amount):
        self.paypal_payment.send_payment(amount)

def process_payment(payment_system, amount):
    payment_system.pay(amount)

stripe_payment = StripePayment()
paypal_payment = PaypalAdapter(PaypalPayment())

process_payment(stripe_payment, 300)
process_payment(paypal_payment, 200)
