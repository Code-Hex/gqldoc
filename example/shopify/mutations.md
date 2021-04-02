# Mutations

### About mutations

The schema’s entry-point for mutations. This acts as the public, top-level API from which all mutation queries must start.

### checkoutAttributesUpdate

<p>Updates the attributes of a checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- input ([CheckoutAttributesUpdateInput!](input_objects.md#checkoutattributesupdateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutAttributesUpdateV2

<p>Updates the attributes of a checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- input ([CheckoutAttributesUpdateV2Input!](input_objects.md#checkoutattributesupdatev2input))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCompleteFree

<p>Completes a checkout without providing payment information. You can use this
mutation for free items or items whose purchase price is covered by a gift card.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCompleteWithCreditCard

<p>Completes a checkout using a credit card token from Shopify&rsquo;s Vault.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- payment ([CreditCardPaymentInput!](input_objects.md#creditcardpaymentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The checkout on which the payment was applied.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| payment ([Payment](objects.md#payment)) | <p>A representation of the attempted payment.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCompleteWithCreditCardV2

<p>Completes a checkout using a credit card token from Shopify&rsquo;s card vault.
Before you can complete checkouts using CheckoutCompleteWithCreditCardV2, you
need to  <a href="https://help.shopify.com/api/guides/sales-channel-sdk/getting-started#request-payment-processing"><em>request payment processing</em></a>.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- payment ([CreditCardPaymentInputV2!](input_objects.md#creditcardpaymentinputv2))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The checkout on which the payment was applied.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| payment ([Payment](objects.md#payment)) | <p>A representation of the attempted payment.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCompleteWithTokenizedPayment

<p>Completes a checkout with a tokenized payment.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- payment ([TokenizedPaymentInput!](input_objects.md#tokenizedpaymentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The checkout on which the payment was applied.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| payment ([Payment](objects.md#payment)) | <p>A representation of the attempted payment.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCompleteWithTokenizedPaymentV2

<p>Completes a checkout with a tokenized payment.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- payment ([TokenizedPaymentInputV2!](input_objects.md#tokenizedpaymentinputv2))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The checkout on which the payment was applied.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| payment ([Payment](objects.md#payment)) | <p>A representation of the attempted payment.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCompleteWithTokenizedPaymentV3

<p>Completes a checkout with a tokenized payment.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- payment ([TokenizedPaymentInputV3!](input_objects.md#tokenizedpaymentinputv3))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The checkout on which the payment was applied.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| payment ([Payment](objects.md#payment)) | <p>A representation of the attempted payment.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCreate

<p>Creates a new checkout.</p>

#### Input fields

- input ([CheckoutCreateInput!](input_objects.md#checkoutcreateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The new checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCustomerAssociate

<p>Associates a customer to the checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- customerAccessToken ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The updated checkout object.</p> |
| customer ([Customer](objects.md#customer)) | <p>The associated customer object.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCustomerAssociateV2

<p>Associates a customer to the checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- customerAccessToken ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| customer ([Customer](objects.md#customer)) | <p>The associated customer object.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCustomerDisassociate

<p>Disassociates the current checkout customer from the checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutCustomerDisassociateV2

<p>Disassociates the current checkout customer from the checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutDiscountCodeApply

<p>Applies a discount to an existing checkout using a discount code.</p>

#### Input fields

- discountCode ([String!](scalars.md#string))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutDiscountCodeApplyV2

<p>Applies a discount to an existing checkout using a discount code.</p>

#### Input fields

- discountCode ([String!](scalars.md#string))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutDiscountCodeRemove

<p>Removes the applied discount from an existing checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutEmailUpdate

<p>Updates the email on an existing checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- email ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The checkout object with the updated email.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutEmailUpdateV2

<p>Updates the email on an existing checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- email ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The checkout object with the updated email.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutGiftCardApply

<p>Applies a gift card to an existing checkout using a gift card code. This will replace all currently applied gift cards.</p>

#### Input fields

- giftCardCode ([String!](scalars.md#string))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutGiftCardRemove

<p>Removes an applied gift card from the checkout.</p>

#### Input fields

- appliedGiftCardId ([ID!](scalars.md#id))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutGiftCardRemoveV2

<p>Removes an applied gift card from the checkout.</p>

#### Input fields

- appliedGiftCardId ([ID!](scalars.md#id))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutGiftCardsAppend

<p>Appends gift cards to an existing checkout.</p>

#### Input fields

- giftCardCodes ([[String!]!](scalars.md#string))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutLineItemsAdd

<p>Adds a list of line items to a checkout.</p>

#### Input fields

- lineItems ([[CheckoutLineItemInput!]!](input_objects.md#checkoutlineiteminput))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutLineItemsRemove

<p>Removes line items from an existing checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- lineItemIds ([[ID!]!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutLineItemsReplace

<p>Sets a list of line items to a checkout.</p>

#### Input fields

- lineItems ([[CheckoutLineItemInput!]!](input_objects.md#checkoutlineiteminput))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| userErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutLineItemsUpdate

<p>Updates line items on a checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- lineItems ([[CheckoutLineItemUpdateInput!]!](input_objects.md#checkoutlineitemupdateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutShippingAddressUpdate

<p>Updates the shipping address of an existing checkout.</p>

#### Input fields

- shippingAddress ([MailingAddressInput!](input_objects.md#mailingaddressinput))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout!](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutShippingAddressUpdateV2

<p>Updates the shipping address of an existing checkout.</p>

#### Input fields

- shippingAddress ([MailingAddressInput!](input_objects.md#mailingaddressinput))

- checkoutId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### checkoutShippingLineUpdate

<p>Updates the shipping lines on an existing checkout.</p>

#### Input fields

- checkoutId ([ID!](scalars.md#id))

- shippingRateHandle ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| checkout ([Checkout](objects.md#checkout)) | <p>The updated checkout object.</p> |
| checkoutUserErrors ([[CheckoutUserError!]!](objects.md#checkoutusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerAccessTokenCreate

<p>Creates a customer access token.
The customer access token is required to modify the customer object in any way.</p>

#### Input fields

- input ([CustomerAccessTokenCreateInput!](input_objects.md#customeraccesstokencreateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>The newly created customer access token object.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerAccessTokenCreateWithMultipass

<p>Creates a customer access token using a multipass token instead of email and password.
A customer record is created if customer does not exist. If a customer record already
exists but the record is disabled, then it&rsquo;s enabled.</p>

#### Input fields

- multipassToken ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>An access token object associated with the customer.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerAccessTokenDelete

<p>Permanently destroys a customer access token.</p>

#### Input fields

- customerAccessToken ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| deletedAccessToken ([String](scalars.md#string)) | <p>The destroyed access token.</p> |
| deletedCustomerAccessTokenId ([String](scalars.md#string)) | <p>ID of the destroyed customer access token.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerAccessTokenRenew

<p>Renews a customer access token.</p>

<p>Access token renewal must happen <em>before</em> a token expires.
If a token has already expired, a new one should be created instead via <code>customerAccessTokenCreate</code>.</p>

#### Input fields

- customerAccessToken ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>The renewed customer access token object.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerActivate

<p>Activates a customer.</p>

#### Input fields

- id ([ID!](scalars.md#id))

- input ([CustomerActivateInput!](input_objects.md#customeractivateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| customer ([Customer](objects.md#customer)) | <p>The customer object.</p> |
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>A newly created customer access token object for the customer.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerActivateByUrl

<p>Activates a customer with the activation url received from <code>customerCreate</code>.</p>

#### Input fields

- activationUrl ([URL!](scalars.md#url))

- password ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| customer ([Customer](objects.md#customer)) | <p>The customer that was activated.</p> |
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>A new customer access token for the customer.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerAddressCreate

<p>Creates a new address for a customer.</p>

#### Input fields

- customerAccessToken ([String!](scalars.md#string))

- address ([MailingAddressInput!](input_objects.md#mailingaddressinput))
 

#### Returns

| Name | Description |
|------|-------------|
| customerAddress ([MailingAddress](objects.md#mailingaddress)) | <p>The new customer address object.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerAddressDelete

<p>Permanently deletes the address of an existing customer.</p>

#### Input fields

- id ([ID!](scalars.md#id))

- customerAccessToken ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| deletedCustomerAddressId ([String](scalars.md#string)) | <p>ID of the deleted customer address.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerAddressUpdate

<p>Updates the address of an existing customer.</p>

#### Input fields

- customerAccessToken ([String!](scalars.md#string))

- id ([ID!](scalars.md#id))

- address ([MailingAddressInput!](input_objects.md#mailingaddressinput))
 

#### Returns

| Name | Description |
|------|-------------|
| customerAddress ([MailingAddress](objects.md#mailingaddress)) | <p>The customer’s updated mailing address.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerCreate

<p>Creates a new customer.</p>

#### Input fields

- input ([CustomerCreateInput!](input_objects.md#customercreateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| customer ([Customer](objects.md#customer)) | <p>The created customer object.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerDefaultAddressUpdate

<p>Updates the default address of an existing customer.</p>

#### Input fields

- customerAccessToken ([String!](scalars.md#string))

- addressId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| customer ([Customer](objects.md#customer)) | <p>The updated customer object.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerRecover

<p>Sends a reset password email to the customer, as the first step in the reset password process.</p>

#### Input fields

- email ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerReset

<p>Resets a customer’s password with a token received from <code>CustomerRecover</code>.</p>

#### Input fields

- id ([ID!](scalars.md#id))

- input ([CustomerResetInput!](input_objects.md#customerresetinput))
 

#### Returns

| Name | Description |
|------|-------------|
| customer ([Customer](objects.md#customer)) | <p>The customer object which was reset.</p> |
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>A newly created customer access token object for the customer.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerResetByUrl

<p>Resets a customer’s password with the reset password url received from <code>CustomerRecover</code>.</p>

#### Input fields

- resetUrl ([URL!](scalars.md#url))

- password ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| customer ([Customer](objects.md#customer)) | <p>The customer object which was reset.</p> |
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>A newly created customer access token object for the customer.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---

### customerUpdate

<p>Updates an existing customer.</p>

#### Input fields

- customerAccessToken ([String!](scalars.md#string))

- customer ([CustomerUpdateInput!](input_objects.md#customerupdateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| customer ([Customer](objects.md#customer)) | <p>The updated customer object.</p> |
| customerAccessToken ([CustomerAccessToken](objects.md#customeraccesstoken)) | <p>The newly created customer access token. If the customer&rsquo;s password is updated, all previous access tokens
(including the one used to perform this mutation) become invalid, and a new token is generated.</p> |
| customerUserErrors ([[CustomerUserError!]!](objects.md#customerusererror)) | <p>List of errors that occurred executing the mutation.</p> |
| userErrors ([[UserError!]!](objects.md#usererror)) | <p>List of errors that occurred executing the mutation.</p> |

---