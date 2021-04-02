# Input objects

### About input objects

[Input objects](https://graphql.github.io/graphql-spec/June2018/#sec-Input-Objects) can be described as "composable objects" because they include a set of input fields that define the object.

### AttributeInput

<p>Specifies the input fields required for an attribute.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>key</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>Key or name of the attribute.</p></td>
  </tr>
  <tr>
    <td><strong>value</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>Value of the attribute.</p></td>
  </tr>
</table>

---

### CheckoutAttributesUpdateInput

<p>Specifies the fields required to update a checkout&rsquo;s attributes.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allowPartialAddresses</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Allows setting partial addresses on a Checkout, skipping the full validation of attributes.
The required attributes are city, province, and country.
Full validation of the addresses is still done at complete time.</p></td>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="input_objects.md#attributeinput">[AttributeInput!]</a>)</td>
    <td><p>A list of extra information that is added to the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>note</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The text of an optional note that a shop owner can attach to the checkout.</p></td>
  </tr>
</table>

---

### CheckoutAttributesUpdateV2Input

<p>Specifies the fields required to update a checkout&rsquo;s attributes.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allowPartialAddresses</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Allows setting partial addresses on a Checkout, skipping the full validation of attributes.
The required attributes are city, province, and country.
Full validation of the addresses is still done at complete time.</p></td>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="input_objects.md#attributeinput">[AttributeInput!]</a>)</td>
    <td><p>A list of extra information that is added to the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>note</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The text of an optional note that a shop owner can attach to the checkout.</p></td>
  </tr>
</table>

---

### CheckoutCreateInput

<p>Specifies the fields required to create a checkout.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allowPartialAddresses</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Allows setting partial addresses on a Checkout, skipping the full validation of attributes.
The required attributes are city, province, and country.
Full validation of addresses is still done at complete time.</p></td>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="input_objects.md#attributeinput">[AttributeInput!]</a>)</td>
    <td><p>A list of extra information that is added to the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The email with which the customer wants to checkout.</p></td>
  </tr>
  <tr>
    <td><strong>lineItems</strong> (<a href="input_objects.md#checkoutlineiteminput">[CheckoutLineItemInput!]</a>)</td>
    <td><p>A list of line item objects, each one containing information about an item in the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>note</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The text of an optional note that a shop owner can attach to the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>presentmentCurrencyCode</strong> (<a href="enums.md#currencycode">CurrencyCode</a>)</td>
    <td><p>The three-letter currency code of one of the shop&rsquo;s enabled presentment currencies.
Including this field creates a checkout in the specified currency. By default, new
checkouts are created in the shop&rsquo;s primary currency.</p></td>
  </tr>
  <tr>
    <td><strong>shippingAddress</strong> (<a href="input_objects.md#mailingaddressinput">MailingAddressInput</a>)</td>
    <td><p>The shipping address to where the line items will be shipped.</p></td>
  </tr>
</table>

---

### CheckoutLineItemInput

<p>Specifies the input fields to create a line item on a checkout.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="input_objects.md#attributeinput">[AttributeInput!]</a>)</td>
    <td><p>Extra information in the form of an array of Key-Value pairs about the line item.</p></td>
  </tr>
  <tr>
    <td><strong>quantity</strong> (<a href="scalars.md#int">Int!</a>)</td>
    <td><p>The quantity of the line item.</p></td>
  </tr>
  <tr>
    <td><strong>variantId</strong> (<a href="scalars.md#id">ID!</a>)</td>
    <td><p>The identifier of the product variant for the line item.</p></td>
  </tr>
</table>

---

### CheckoutLineItemUpdateInput

<p>Specifies the input fields to update a line item on the checkout.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="input_objects.md#attributeinput">[AttributeInput!]</a>)</td>
    <td><p>Extra information in the form of an array of Key-Value pairs about the line item.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID</a>)</td>
    <td><p>The identifier of the line item.</p></td>
  </tr>
  <tr>
    <td><strong>quantity</strong> (<a href="scalars.md#int">Int</a>)</td>
    <td><p>The quantity of the line item.</p></td>
  </tr>
  <tr>
    <td><strong>variantId</strong> (<a href="scalars.md#id">ID</a>)</td>
    <td><p>The variant identifier of the line item.</p></td>
  </tr>
</table>

---

### CreditCardPaymentInput

<p>Specifies the fields required to complete a checkout with
a Shopify vaulted credit card payment.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>amount</strong> (<a href="scalars.md#money">Money!</a>)</td>
    <td><p>The amount of the payment.</p></td>
  </tr>
  <tr>
    <td><strong>billingAddress</strong> (<a href="input_objects.md#mailingaddressinput">MailingAddressInput!</a>)</td>
    <td><p>The billing address for the payment.</p></td>
  </tr>
  <tr>
    <td><strong>idempotencyKey</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A unique client generated key used to avoid duplicate charges. When a
duplicate payment is found, the original is returned instead of creating a new
one. For more information, refer to <a href="https://shopify.dev/concepts/about-apis/idempotent-requests">Idempotent
requests</a>.</p></td>
  </tr>
  <tr>
    <td><strong>test</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Executes the payment in test mode if possible. Defaults to <code>false</code>.</p></td>
  </tr>
  <tr>
    <td><strong>vaultId</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The ID returned by Shopify&rsquo;s Card Vault.</p></td>
  </tr>
</table>

---

### CreditCardPaymentInputV2

<p>Specifies the fields required to complete a checkout with
a Shopify vaulted credit card payment.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>billingAddress</strong> (<a href="input_objects.md#mailingaddressinput">MailingAddressInput!</a>)</td>
    <td><p>The billing address for the payment.</p></td>
  </tr>
  <tr>
    <td><strong>idempotencyKey</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A unique client generated key used to avoid duplicate charges. When a
duplicate payment is found, the original is returned instead of creating a new
one. For more information, refer to <a href="https://shopify.dev/concepts/about-apis/idempotent-requests">Idempotent
requests</a>.</p></td>
  </tr>
  <tr>
    <td><strong>paymentAmount</strong> (<a href="input_objects.md#moneyinput">MoneyInput!</a>)</td>
    <td><p>The amount and currency of the payment.</p></td>
  </tr>
  <tr>
    <td><strong>test</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Executes the payment in test mode if possible. Defaults to <code>false</code>.</p></td>
  </tr>
  <tr>
    <td><strong>vaultId</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The ID returned by Shopify&rsquo;s Card Vault.</p></td>
  </tr>
</table>

---

### CustomerAccessTokenCreateInput

<p>Specifies the input fields required to create a customer access token.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The email associated to the customer.</p></td>
  </tr>
  <tr>
    <td><strong>password</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The login password to be used by the customer.</p></td>
  </tr>
</table>

---

### CustomerActivateInput

<p>Specifies the input fields required to activate a customer.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>activationToken</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The activation token required to activate the customer.</p></td>
  </tr>
  <tr>
    <td><strong>password</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>New password that will be set during activation.</p></td>
  </tr>
</table>

---

### CustomerCreateInput

<p>Specifies the fields required to create a new customer.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>acceptsMarketing</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Indicates whether the customer has consented to be sent marketing material via email.</p></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The customer’s email.</p></td>
  </tr>
  <tr>
    <td><strong>firstName</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The customer’s first name.</p></td>
  </tr>
  <tr>
    <td><strong>lastName</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The customer’s last name.</p></td>
  </tr>
  <tr>
    <td><strong>password</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The login password used by the customer.</p></td>
  </tr>
  <tr>
    <td><strong>phone</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>A unique phone number for the customer.</p>

<p>Formatted using E.164 standard. For example, <em>+16135551111</em>.</p></td>
  </tr>
</table>

---

### CustomerResetInput

<p>Specifies the fields required to reset a customer’s password.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>password</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>New password that will be set as part of the reset password process.</p></td>
  </tr>
  <tr>
    <td><strong>resetToken</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The reset token required to reset the customer’s password.</p></td>
  </tr>
</table>

---

### CustomerUpdateInput

<p>Specifies the fields required to update the Customer information.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>acceptsMarketing</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Indicates whether the customer has consented to be sent marketing material via email.</p></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The customer’s email.</p></td>
  </tr>
  <tr>
    <td><strong>firstName</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The customer’s first name.</p></td>
  </tr>
  <tr>
    <td><strong>lastName</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The customer’s last name.</p></td>
  </tr>
  <tr>
    <td><strong>password</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The login password used by the customer.</p></td>
  </tr>
  <tr>
    <td><strong>phone</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>A unique phone number for the customer.</p>

<p>Formatted using E.164 standard. For example, <em>+16135551111</em>. To remove the phone number, specify <code>null</code>.</p></td>
  </tr>
</table>

---

### MailingAddressInput

<p>Specifies the fields accepted to create or update a mailing address.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>address1</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The first line of the address. Typically the street address or PO Box number.</p></td>
  </tr>
  <tr>
    <td><strong>address2</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The second line of the address. Typically the number of the apartment, suite, or unit.</p></td>
  </tr>
  <tr>
    <td><strong>city</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The name of the city, district, village, or town.</p></td>
  </tr>
  <tr>
    <td><strong>company</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The name of the customer&rsquo;s company or organization.</p></td>
  </tr>
  <tr>
    <td><strong>country</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The name of the country.</p></td>
  </tr>
  <tr>
    <td><strong>firstName</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The first name of the customer.</p></td>
  </tr>
  <tr>
    <td><strong>lastName</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The last name of the customer.</p></td>
  </tr>
  <tr>
    <td><strong>phone</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>A unique phone number for the customer.</p>

<p>Formatted using E.164 standard. For example, <em>+16135551111</em>.</p></td>
  </tr>
  <tr>
    <td><strong>province</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The region of the address, such as the province, state, or district.</p></td>
  </tr>
  <tr>
    <td><strong>zip</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>The zip or postal code of the address.</p></td>
  </tr>
</table>

---

### MoneyInput

<p>Specifies the fields for a monetary value with currency.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>amount</strong> (<a href="scalars.md#decimal">Decimal!</a>)</td>
    <td><p>Decimal money amount.</p></td>
  </tr>
  <tr>
    <td><strong>currencyCode</strong> (<a href="enums.md#currencycode">CurrencyCode!</a>)</td>
    <td><p>Currency of the money.</p></td>
  </tr>
</table>

---

### SelectedOptionInput

<p>Specifies the input fields required for a selected option.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The product option’s name.</p></td>
  </tr>
  <tr>
    <td><strong>value</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The product option’s value.</p></td>
  </tr>
</table>

---

### TokenizedPaymentInput

<p>Specifies the fields required to complete a checkout with
a tokenized payment.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>amount</strong> (<a href="scalars.md#money">Money!</a>)</td>
    <td><p>The amount of the payment.</p></td>
  </tr>
  <tr>
    <td><strong>billingAddress</strong> (<a href="input_objects.md#mailingaddressinput">MailingAddressInput!</a>)</td>
    <td><p>The billing address for the payment.</p></td>
  </tr>
  <tr>
    <td><strong>idempotencyKey</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A unique client generated key used to avoid duplicate charges. When a
duplicate payment is found, the original is returned instead of creating a new
one. For more information, refer to <a href="https://shopify.dev/concepts/about-apis/idempotent-requests">Idempotent
requests</a>.</p></td>
  </tr>
  <tr>
    <td><strong>identifier</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>Public Hash Key used for AndroidPay payments only.</p></td>
  </tr>
  <tr>
    <td><strong>paymentData</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A simple string or JSON containing the required payment data for the tokenized payment.</p></td>
  </tr>
  <tr>
    <td><strong>test</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Executes the payment in test mode if possible. Defaults to <code>false</code>.</p></td>
  </tr>
  <tr>
    <td><strong>type</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The type of payment token.</p></td>
  </tr>
</table>

---

### TokenizedPaymentInputV2

<p>Specifies the fields required to complete a checkout with
a tokenized payment.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>billingAddress</strong> (<a href="input_objects.md#mailingaddressinput">MailingAddressInput!</a>)</td>
    <td><p>The billing address for the payment.</p></td>
  </tr>
  <tr>
    <td><strong>idempotencyKey</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A unique client generated key used to avoid duplicate charges. When a
duplicate payment is found, the original is returned instead of creating a new
one. For more information, refer to <a href="https://shopify.dev/concepts/about-apis/idempotent-requests">Idempotent
requests</a>.</p></td>
  </tr>
  <tr>
    <td><strong>identifier</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>Public Hash Key used for AndroidPay payments only.</p></td>
  </tr>
  <tr>
    <td><strong>paymentAmount</strong> (<a href="input_objects.md#moneyinput">MoneyInput!</a>)</td>
    <td><p>The amount and currency of the payment.</p></td>
  </tr>
  <tr>
    <td><strong>paymentData</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A simple string or JSON containing the required payment data for the tokenized payment.</p></td>
  </tr>
  <tr>
    <td><strong>test</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Whether to execute the payment in test mode, if possible. Test mode is not
supported in production stores. Defaults to <code>false</code>.</p></td>
  </tr>
  <tr>
    <td><strong>type</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>The type of payment token.</p></td>
  </tr>
</table>

---

### TokenizedPaymentInputV3

<p>Specifies the fields required to complete a checkout with
a tokenized payment.</p>


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>billingAddress</strong> (<a href="input_objects.md#mailingaddressinput">MailingAddressInput!</a>)</td>
    <td><p>The billing address for the payment.</p></td>
  </tr>
  <tr>
    <td><strong>idempotencyKey</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A unique client generated key used to avoid duplicate charges. When a
duplicate payment is found, the original is returned instead of creating a new
one. For more information, refer to <a href="https://shopify.dev/concepts/about-apis/idempotent-requests">Idempotent
requests</a>.</p></td>
  </tr>
  <tr>
    <td><strong>identifier</strong> (<a href="scalars.md#string">String</a>)</td>
    <td><p>Public Hash Key used for AndroidPay payments only.</p></td>
  </tr>
  <tr>
    <td><strong>paymentAmount</strong> (<a href="input_objects.md#moneyinput">MoneyInput!</a>)</td>
    <td><p>The amount and currency of the payment.</p></td>
  </tr>
  <tr>
    <td><strong>paymentData</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td><p>A simple string or JSON containing the required payment data for the tokenized payment.</p></td>
  </tr>
  <tr>
    <td><strong>test</strong> (<a href="scalars.md#boolean">Boolean</a>)</td>
    <td><p>Whether to execute the payment in test mode, if possible. Test mode is not
supported in production stores. Defaults to <code>false</code>.</p></td>
  </tr>
  <tr>
    <td><strong>type</strong> (<a href="enums.md#paymenttokentype">PaymentTokenType!</a>)</td>
    <td><p>The type of payment token.</p></td>
  </tr>
</table>

---