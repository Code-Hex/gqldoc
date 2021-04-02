# Enums

### About enums

[Enums](https://graphql.github.io/graphql-spec/June2018/#sec-Enums) represent possible sets of values for a field.

### ArticleSortKeys

<p>The set of valid sort keys for the Article query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AUTHOR</strong></td>
    <td><p>Sort by the <code>author</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>BLOG_TITLE</strong></td>
    <td><p>Sort by the <code>blog_title</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLISHED_AT</strong></td>
    <td><p>Sort by the <code>published_at</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>TITLE</strong></td>
    <td><p>Sort by the <code>title</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Sort by the <code>updated_at</code> value.</p></td>
  </tr>
</table>

---

### BlogSortKeys

<p>The set of valid sort keys for the Blog query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>HANDLE</strong></td>
    <td><p>Sort by the <code>handle</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>TITLE</strong></td>
    <td><p>Sort by the <code>title</code> value.</p></td>
  </tr>
</table>

---

### CardBrand

<p>Card brand, such as Visa or Mastercard, which can be used for payments.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AMERICAN_EXPRESS</strong></td>
    <td><p>American Express.</p></td>
  </tr>
  <tr>
    <td><strong>DINERS_CLUB</strong></td>
    <td><p>Diners Club.</p></td>
  </tr>
  <tr>
    <td><strong>DISCOVER</strong></td>
    <td><p>Discover.</p></td>
  </tr>
  <tr>
    <td><strong>JCB</strong></td>
    <td><p>JCB.</p></td>
  </tr>
  <tr>
    <td><strong>MASTERCARD</strong></td>
    <td><p>Mastercard.</p></td>
  </tr>
  <tr>
    <td><strong>VISA</strong></td>
    <td><p>Visa.</p></td>
  </tr>
</table>

---

### CheckoutErrorCode

<p>Possible error codes that could be returned by CheckoutUserError.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALREADY_COMPLETED</strong></td>
    <td><p>Checkout is already completed.</p></td>
  </tr>
  <tr>
    <td><strong>BAD_DOMAIN</strong></td>
    <td><p>Input email contains an invalid domain name.</p></td>
  </tr>
  <tr>
    <td><strong>BLANK</strong></td>
    <td><p>Input value is blank.</p></td>
  </tr>
  <tr>
    <td><strong>CART_DOES_NOT_MEET_DISCOUNT_REQUIREMENTS_NOTICE</strong></td>
    <td><p>Cart does not meet discount requirements notice.</p></td>
  </tr>
  <tr>
    <td><strong>CUSTOMER_ALREADY_USED_ONCE_PER_CUSTOMER_DISCOUNT_NOTICE</strong></td>
    <td><p>Customer already used once per customer discount notice.</p></td>
  </tr>
  <tr>
    <td><strong>DISCOUNT_ALREADY_APPLIED</strong></td>
    <td><p>Discount already applied.</p></td>
  </tr>
  <tr>
    <td><strong>DISCOUNT_DISABLED</strong></td>
    <td><p>Discount disabled.</p></td>
  </tr>
  <tr>
    <td><strong>DISCOUNT_EXPIRED</strong></td>
    <td><p>Discount expired.</p></td>
  </tr>
  <tr>
    <td><strong>DISCOUNT_LIMIT_REACHED</strong></td>
    <td><p>Discount limit reached.</p></td>
  </tr>
  <tr>
    <td><strong>DISCOUNT_NOT_FOUND</strong></td>
    <td><p>Discount not found.</p></td>
  </tr>
  <tr>
    <td><strong>EMPTY</strong></td>
    <td><p>Checkout is already completed.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_ALREADY_APPLIED</strong></td>
    <td><p>Gift card has already been applied.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_CODE_INVALID</strong></td>
    <td><p>Gift card code is invalid.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_CURRENCY_MISMATCH</strong></td>
    <td><p>Gift card currency does not match checkout currency.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_DEPLETED</strong></td>
    <td><p>Gift card has no funds left.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_DISABLED</strong></td>
    <td><p>Gift card is disabled.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_EXPIRED</strong></td>
    <td><p>Gift card is expired.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_NOT_FOUND</strong></td>
    <td><p>Gift card was not found.</p></td>
  </tr>
  <tr>
    <td><strong>GIFT_CARD_UNUSABLE</strong></td>
    <td><p>Gift card cannot be applied to a checkout that contains a gift card.</p></td>
  </tr>
  <tr>
    <td><strong>GREATER_THAN_OR_EQUAL_TO</strong></td>
    <td><p>Input value should be greater than or equal to minimum allowed value.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID</strong></td>
    <td><p>Input value is invalid.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID_FOR_COUNTRY</strong></td>
    <td><p>Input Zip is invalid for country provided.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID_FOR_COUNTRY_AND_PROVINCE</strong></td>
    <td><p>Input Zip is invalid for country and province provided.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID_PROVINCE_IN_COUNTRY</strong></td>
    <td><p>Invalid province in country.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID_REGION_IN_COUNTRY</strong></td>
    <td><p>Invalid region in country.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID_STATE_IN_COUNTRY</strong></td>
    <td><p>Invalid state in country.</p></td>
  </tr>
  <tr>
    <td><strong>LESS_THAN</strong></td>
    <td><p>Input value should be less than maximum allowed value.</p></td>
  </tr>
  <tr>
    <td><strong>LESS_THAN_OR_EQUAL_TO</strong></td>
    <td><p>Input value should be less or equal to maximum allowed value.</p></td>
  </tr>
  <tr>
    <td><strong>LINE_ITEM_NOT_FOUND</strong></td>
    <td><p>Line item was not found in checkout.</p></td>
  </tr>
  <tr>
    <td><strong>LOCKED</strong></td>
    <td><p>Checkout is locked.</p></td>
  </tr>
  <tr>
    <td><strong>MISSING_PAYMENT_INPUT</strong></td>
    <td><p>Missing payment input.</p></td>
  </tr>
  <tr>
    <td><strong>NOT_ENOUGH_IN_STOCK</strong></td>
    <td><p>Not enough in stock.</p></td>
  </tr>
  <tr>
    <td><strong>NOT_SUPPORTED</strong></td>
    <td><p>Input value is not supported.</p></td>
  </tr>
  <tr>
    <td><strong>PRESENT</strong></td>
    <td><p>Input value is not present.</p></td>
  </tr>
  <tr>
    <td><strong>SHIPPING_RATE_EXPIRED</strong></td>
    <td><p>Shipping rate expired.</p></td>
  </tr>
  <tr>
    <td><strong>TOO_LONG</strong></td>
    <td><p>Input value is too long.</p></td>
  </tr>
  <tr>
    <td><strong>TOTAL_PRICE_MISMATCH</strong></td>
    <td><p>The amount of the payment does not match the value to be paid.</p></td>
  </tr>
  <tr>
    <td><strong>UNABLE_TO_APPLY</strong></td>
    <td><p>Unable to apply discount.</p></td>
  </tr>
</table>

---

### CollectionSortKeys

<p>The set of valid sort keys for the Collection query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>TITLE</strong></td>
    <td><p>Sort by the <code>title</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Sort by the <code>updated_at</code> value.</p></td>
  </tr>
</table>

---

### CountryCode

<p>ISO 3166-1 alpha-2 country codes with some differences.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AC</strong></td>
    <td><p>Ascension Island.</p></td>
  </tr>
  <tr>
    <td><strong>AD</strong></td>
    <td><p>Andorra.</p></td>
  </tr>
  <tr>
    <td><strong>AE</strong></td>
    <td><p>United Arab Emirates.</p></td>
  </tr>
  <tr>
    <td><strong>AF</strong></td>
    <td><p>Afghanistan.</p></td>
  </tr>
  <tr>
    <td><strong>AG</strong></td>
    <td><p>Antigua &amp; Barbuda.</p></td>
  </tr>
  <tr>
    <td><strong>AI</strong></td>
    <td><p>Anguilla.</p></td>
  </tr>
  <tr>
    <td><strong>AL</strong></td>
    <td><p>Albania.</p></td>
  </tr>
  <tr>
    <td><strong>AM</strong></td>
    <td><p>Armenia.</p></td>
  </tr>
  <tr>
    <td><strong>AN</strong></td>
    <td><p>Netherlands Antilles.</p></td>
  </tr>
  <tr>
    <td><strong>AO</strong></td>
    <td><p>Angola.</p></td>
  </tr>
  <tr>
    <td><strong>AR</strong></td>
    <td><p>Argentina.</p></td>
  </tr>
  <tr>
    <td><strong>AT</strong></td>
    <td><p>Austria.</p></td>
  </tr>
  <tr>
    <td><strong>AU</strong></td>
    <td><p>Australia.</p></td>
  </tr>
  <tr>
    <td><strong>AW</strong></td>
    <td><p>Aruba.</p></td>
  </tr>
  <tr>
    <td><strong>AX</strong></td>
    <td><p>Åland Islands.</p></td>
  </tr>
  <tr>
    <td><strong>AZ</strong></td>
    <td><p>Azerbaijan.</p></td>
  </tr>
  <tr>
    <td><strong>BA</strong></td>
    <td><p>Bosnia &amp; Herzegovina.</p></td>
  </tr>
  <tr>
    <td><strong>BB</strong></td>
    <td><p>Barbados.</p></td>
  </tr>
  <tr>
    <td><strong>BD</strong></td>
    <td><p>Bangladesh.</p></td>
  </tr>
  <tr>
    <td><strong>BE</strong></td>
    <td><p>Belgium.</p></td>
  </tr>
  <tr>
    <td><strong>BF</strong></td>
    <td><p>Burkina Faso.</p></td>
  </tr>
  <tr>
    <td><strong>BG</strong></td>
    <td><p>Bulgaria.</p></td>
  </tr>
  <tr>
    <td><strong>BH</strong></td>
    <td><p>Bahrain.</p></td>
  </tr>
  <tr>
    <td><strong>BI</strong></td>
    <td><p>Burundi.</p></td>
  </tr>
  <tr>
    <td><strong>BJ</strong></td>
    <td><p>Benin.</p></td>
  </tr>
  <tr>
    <td><strong>BL</strong></td>
    <td><p>St. Barthélemy.</p></td>
  </tr>
  <tr>
    <td><strong>BM</strong></td>
    <td><p>Bermuda.</p></td>
  </tr>
  <tr>
    <td><strong>BN</strong></td>
    <td><p>Brunei.</p></td>
  </tr>
  <tr>
    <td><strong>BO</strong></td>
    <td><p>Bolivia.</p></td>
  </tr>
  <tr>
    <td><strong>BQ</strong></td>
    <td><p>Caribbean Netherlands.</p></td>
  </tr>
  <tr>
    <td><strong>BR</strong></td>
    <td><p>Brazil.</p></td>
  </tr>
  <tr>
    <td><strong>BS</strong></td>
    <td><p>Bahamas.</p></td>
  </tr>
  <tr>
    <td><strong>BT</strong></td>
    <td><p>Bhutan.</p></td>
  </tr>
  <tr>
    <td><strong>BV</strong></td>
    <td><p>Bouvet Island.</p></td>
  </tr>
  <tr>
    <td><strong>BW</strong></td>
    <td><p>Botswana.</p></td>
  </tr>
  <tr>
    <td><strong>BY</strong></td>
    <td><p>Belarus.</p></td>
  </tr>
  <tr>
    <td><strong>BZ</strong></td>
    <td><p>Belize.</p></td>
  </tr>
  <tr>
    <td><strong>CA</strong></td>
    <td><p>Canada.</p></td>
  </tr>
  <tr>
    <td><strong>CC</strong></td>
    <td><p>Cocos (Keeling) Islands.</p></td>
  </tr>
  <tr>
    <td><strong>CD</strong></td>
    <td><p>Congo - Kinshasa.</p></td>
  </tr>
  <tr>
    <td><strong>CF</strong></td>
    <td><p>Central African Republic.</p></td>
  </tr>
  <tr>
    <td><strong>CG</strong></td>
    <td><p>Congo - Brazzaville.</p></td>
  </tr>
  <tr>
    <td><strong>CH</strong></td>
    <td><p>Switzerland.</p></td>
  </tr>
  <tr>
    <td><strong>CI</strong></td>
    <td><p>Côte d’Ivoire.</p></td>
  </tr>
  <tr>
    <td><strong>CK</strong></td>
    <td><p>Cook Islands.</p></td>
  </tr>
  <tr>
    <td><strong>CL</strong></td>
    <td><p>Chile.</p></td>
  </tr>
  <tr>
    <td><strong>CM</strong></td>
    <td><p>Cameroon.</p></td>
  </tr>
  <tr>
    <td><strong>CN</strong></td>
    <td><p>China.</p></td>
  </tr>
  <tr>
    <td><strong>CO</strong></td>
    <td><p>Colombia.</p></td>
  </tr>
  <tr>
    <td><strong>CR</strong></td>
    <td><p>Costa Rica.</p></td>
  </tr>
  <tr>
    <td><strong>CU</strong></td>
    <td><p>Cuba.</p></td>
  </tr>
  <tr>
    <td><strong>CV</strong></td>
    <td><p>Cape Verde.</p></td>
  </tr>
  <tr>
    <td><strong>CW</strong></td>
    <td><p>Curaçao.</p></td>
  </tr>
  <tr>
    <td><strong>CX</strong></td>
    <td><p>Christmas Island.</p></td>
  </tr>
  <tr>
    <td><strong>CY</strong></td>
    <td><p>Cyprus.</p></td>
  </tr>
  <tr>
    <td><strong>CZ</strong></td>
    <td><p>Czechia.</p></td>
  </tr>
  <tr>
    <td><strong>DE</strong></td>
    <td><p>Germany.</p></td>
  </tr>
  <tr>
    <td><strong>DJ</strong></td>
    <td><p>Djibouti.</p></td>
  </tr>
  <tr>
    <td><strong>DK</strong></td>
    <td><p>Denmark.</p></td>
  </tr>
  <tr>
    <td><strong>DM</strong></td>
    <td><p>Dominica.</p></td>
  </tr>
  <tr>
    <td><strong>DO</strong></td>
    <td><p>Dominican Republic.</p></td>
  </tr>
  <tr>
    <td><strong>DZ</strong></td>
    <td><p>Algeria.</p></td>
  </tr>
  <tr>
    <td><strong>EC</strong></td>
    <td><p>Ecuador.</p></td>
  </tr>
  <tr>
    <td><strong>EE</strong></td>
    <td><p>Estonia.</p></td>
  </tr>
  <tr>
    <td><strong>EG</strong></td>
    <td><p>Egypt.</p></td>
  </tr>
  <tr>
    <td><strong>EH</strong></td>
    <td><p>Western Sahara.</p></td>
  </tr>
  <tr>
    <td><strong>ER</strong></td>
    <td><p>Eritrea.</p></td>
  </tr>
  <tr>
    <td><strong>ES</strong></td>
    <td><p>Spain.</p></td>
  </tr>
  <tr>
    <td><strong>ET</strong></td>
    <td><p>Ethiopia.</p></td>
  </tr>
  <tr>
    <td><strong>FI</strong></td>
    <td><p>Finland.</p></td>
  </tr>
  <tr>
    <td><strong>FJ</strong></td>
    <td><p>Fiji.</p></td>
  </tr>
  <tr>
    <td><strong>FK</strong></td>
    <td><p>Falkland Islands.</p></td>
  </tr>
  <tr>
    <td><strong>FO</strong></td>
    <td><p>Faroe Islands.</p></td>
  </tr>
  <tr>
    <td><strong>FR</strong></td>
    <td><p>France.</p></td>
  </tr>
  <tr>
    <td><strong>GA</strong></td>
    <td><p>Gabon.</p></td>
  </tr>
  <tr>
    <td><strong>GB</strong></td>
    <td><p>United Kingdom.</p></td>
  </tr>
  <tr>
    <td><strong>GD</strong></td>
    <td><p>Grenada.</p></td>
  </tr>
  <tr>
    <td><strong>GE</strong></td>
    <td><p>Georgia.</p></td>
  </tr>
  <tr>
    <td><strong>GF</strong></td>
    <td><p>French Guiana.</p></td>
  </tr>
  <tr>
    <td><strong>GG</strong></td>
    <td><p>Guernsey.</p></td>
  </tr>
  <tr>
    <td><strong>GH</strong></td>
    <td><p>Ghana.</p></td>
  </tr>
  <tr>
    <td><strong>GI</strong></td>
    <td><p>Gibraltar.</p></td>
  </tr>
  <tr>
    <td><strong>GL</strong></td>
    <td><p>Greenland.</p></td>
  </tr>
  <tr>
    <td><strong>GM</strong></td>
    <td><p>Gambia.</p></td>
  </tr>
  <tr>
    <td><strong>GN</strong></td>
    <td><p>Guinea.</p></td>
  </tr>
  <tr>
    <td><strong>GP</strong></td>
    <td><p>Guadeloupe.</p></td>
  </tr>
  <tr>
    <td><strong>GQ</strong></td>
    <td><p>Equatorial Guinea.</p></td>
  </tr>
  <tr>
    <td><strong>GR</strong></td>
    <td><p>Greece.</p></td>
  </tr>
  <tr>
    <td><strong>GS</strong></td>
    <td><p>South Georgia &amp; South Sandwich Islands.</p></td>
  </tr>
  <tr>
    <td><strong>GT</strong></td>
    <td><p>Guatemala.</p></td>
  </tr>
  <tr>
    <td><strong>GW</strong></td>
    <td><p>Guinea-Bissau.</p></td>
  </tr>
  <tr>
    <td><strong>GY</strong></td>
    <td><p>Guyana.</p></td>
  </tr>
  <tr>
    <td><strong>HK</strong></td>
    <td><p>Hong Kong SAR.</p></td>
  </tr>
  <tr>
    <td><strong>HM</strong></td>
    <td><p>Heard &amp; McDonald Islands.</p></td>
  </tr>
  <tr>
    <td><strong>HN</strong></td>
    <td><p>Honduras.</p></td>
  </tr>
  <tr>
    <td><strong>HR</strong></td>
    <td><p>Croatia.</p></td>
  </tr>
  <tr>
    <td><strong>HT</strong></td>
    <td><p>Haiti.</p></td>
  </tr>
  <tr>
    <td><strong>HU</strong></td>
    <td><p>Hungary.</p></td>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Indonesia.</p></td>
  </tr>
  <tr>
    <td><strong>IE</strong></td>
    <td><p>Ireland.</p></td>
  </tr>
  <tr>
    <td><strong>IL</strong></td>
    <td><p>Israel.</p></td>
  </tr>
  <tr>
    <td><strong>IM</strong></td>
    <td><p>Isle of Man.</p></td>
  </tr>
  <tr>
    <td><strong>IN</strong></td>
    <td><p>India.</p></td>
  </tr>
  <tr>
    <td><strong>IO</strong></td>
    <td><p>British Indian Ocean Territory.</p></td>
  </tr>
  <tr>
    <td><strong>IQ</strong></td>
    <td><p>Iraq.</p></td>
  </tr>
  <tr>
    <td><strong>IR</strong></td>
    <td><p>Iran.</p></td>
  </tr>
  <tr>
    <td><strong>IS</strong></td>
    <td><p>Iceland.</p></td>
  </tr>
  <tr>
    <td><strong>IT</strong></td>
    <td><p>Italy.</p></td>
  </tr>
  <tr>
    <td><strong>JE</strong></td>
    <td><p>Jersey.</p></td>
  </tr>
  <tr>
    <td><strong>JM</strong></td>
    <td><p>Jamaica.</p></td>
  </tr>
  <tr>
    <td><strong>JO</strong></td>
    <td><p>Jordan.</p></td>
  </tr>
  <tr>
    <td><strong>JP</strong></td>
    <td><p>Japan.</p></td>
  </tr>
  <tr>
    <td><strong>KE</strong></td>
    <td><p>Kenya.</p></td>
  </tr>
  <tr>
    <td><strong>KG</strong></td>
    <td><p>Kyrgyzstan.</p></td>
  </tr>
  <tr>
    <td><strong>KH</strong></td>
    <td><p>Cambodia.</p></td>
  </tr>
  <tr>
    <td><strong>KI</strong></td>
    <td><p>Kiribati.</p></td>
  </tr>
  <tr>
    <td><strong>KM</strong></td>
    <td><p>Comoros.</p></td>
  </tr>
  <tr>
    <td><strong>KN</strong></td>
    <td><p>St. Kitts &amp; Nevis.</p></td>
  </tr>
  <tr>
    <td><strong>KP</strong></td>
    <td><p>North Korea.</p></td>
  </tr>
  <tr>
    <td><strong>KR</strong></td>
    <td><p>South Korea.</p></td>
  </tr>
  <tr>
    <td><strong>KW</strong></td>
    <td><p>Kuwait.</p></td>
  </tr>
  <tr>
    <td><strong>KY</strong></td>
    <td><p>Cayman Islands.</p></td>
  </tr>
  <tr>
    <td><strong>KZ</strong></td>
    <td><p>Kazakhstan.</p></td>
  </tr>
  <tr>
    <td><strong>LA</strong></td>
    <td><p>Laos.</p></td>
  </tr>
  <tr>
    <td><strong>LB</strong></td>
    <td><p>Lebanon.</p></td>
  </tr>
  <tr>
    <td><strong>LC</strong></td>
    <td><p>St. Lucia.</p></td>
  </tr>
  <tr>
    <td><strong>LI</strong></td>
    <td><p>Liechtenstein.</p></td>
  </tr>
  <tr>
    <td><strong>LK</strong></td>
    <td><p>Sri Lanka.</p></td>
  </tr>
  <tr>
    <td><strong>LR</strong></td>
    <td><p>Liberia.</p></td>
  </tr>
  <tr>
    <td><strong>LS</strong></td>
    <td><p>Lesotho.</p></td>
  </tr>
  <tr>
    <td><strong>LT</strong></td>
    <td><p>Lithuania.</p></td>
  </tr>
  <tr>
    <td><strong>LU</strong></td>
    <td><p>Luxembourg.</p></td>
  </tr>
  <tr>
    <td><strong>LV</strong></td>
    <td><p>Latvia.</p></td>
  </tr>
  <tr>
    <td><strong>LY</strong></td>
    <td><p>Libya.</p></td>
  </tr>
  <tr>
    <td><strong>MA</strong></td>
    <td><p>Morocco.</p></td>
  </tr>
  <tr>
    <td><strong>MC</strong></td>
    <td><p>Monaco.</p></td>
  </tr>
  <tr>
    <td><strong>MD</strong></td>
    <td><p>Moldova.</p></td>
  </tr>
  <tr>
    <td><strong>ME</strong></td>
    <td><p>Montenegro.</p></td>
  </tr>
  <tr>
    <td><strong>MF</strong></td>
    <td><p>St. Martin.</p></td>
  </tr>
  <tr>
    <td><strong>MG</strong></td>
    <td><p>Madagascar.</p></td>
  </tr>
  <tr>
    <td><strong>MK</strong></td>
    <td><p>North Macedonia.</p></td>
  </tr>
  <tr>
    <td><strong>ML</strong></td>
    <td><p>Mali.</p></td>
  </tr>
  <tr>
    <td><strong>MM</strong></td>
    <td><p>Myanmar (Burma).</p></td>
  </tr>
  <tr>
    <td><strong>MN</strong></td>
    <td><p>Mongolia.</p></td>
  </tr>
  <tr>
    <td><strong>MO</strong></td>
    <td><p>Macao SAR.</p></td>
  </tr>
  <tr>
    <td><strong>MQ</strong></td>
    <td><p>Martinique.</p></td>
  </tr>
  <tr>
    <td><strong>MR</strong></td>
    <td><p>Mauritania.</p></td>
  </tr>
  <tr>
    <td><strong>MS</strong></td>
    <td><p>Montserrat.</p></td>
  </tr>
  <tr>
    <td><strong>MT</strong></td>
    <td><p>Malta.</p></td>
  </tr>
  <tr>
    <td><strong>MU</strong></td>
    <td><p>Mauritius.</p></td>
  </tr>
  <tr>
    <td><strong>MV</strong></td>
    <td><p>Maldives.</p></td>
  </tr>
  <tr>
    <td><strong>MW</strong></td>
    <td><p>Malawi.</p></td>
  </tr>
  <tr>
    <td><strong>MX</strong></td>
    <td><p>Mexico.</p></td>
  </tr>
  <tr>
    <td><strong>MY</strong></td>
    <td><p>Malaysia.</p></td>
  </tr>
  <tr>
    <td><strong>MZ</strong></td>
    <td><p>Mozambique.</p></td>
  </tr>
  <tr>
    <td><strong>NA</strong></td>
    <td><p>Namibia.</p></td>
  </tr>
  <tr>
    <td><strong>NC</strong></td>
    <td><p>New Caledonia.</p></td>
  </tr>
  <tr>
    <td><strong>NE</strong></td>
    <td><p>Niger.</p></td>
  </tr>
  <tr>
    <td><strong>NF</strong></td>
    <td><p>Norfolk Island.</p></td>
  </tr>
  <tr>
    <td><strong>NG</strong></td>
    <td><p>Nigeria.</p></td>
  </tr>
  <tr>
    <td><strong>NI</strong></td>
    <td><p>Nicaragua.</p></td>
  </tr>
  <tr>
    <td><strong>NL</strong></td>
    <td><p>Netherlands.</p></td>
  </tr>
  <tr>
    <td><strong>NO</strong></td>
    <td><p>Norway.</p></td>
  </tr>
  <tr>
    <td><strong>NP</strong></td>
    <td><p>Nepal.</p></td>
  </tr>
  <tr>
    <td><strong>NR</strong></td>
    <td><p>Nauru.</p></td>
  </tr>
  <tr>
    <td><strong>NU</strong></td>
    <td><p>Niue.</p></td>
  </tr>
  <tr>
    <td><strong>NZ</strong></td>
    <td><p>New Zealand.</p></td>
  </tr>
  <tr>
    <td><strong>OM</strong></td>
    <td><p>Oman.</p></td>
  </tr>
  <tr>
    <td><strong>PA</strong></td>
    <td><p>Panama.</p></td>
  </tr>
  <tr>
    <td><strong>PE</strong></td>
    <td><p>Peru.</p></td>
  </tr>
  <tr>
    <td><strong>PF</strong></td>
    <td><p>French Polynesia.</p></td>
  </tr>
  <tr>
    <td><strong>PG</strong></td>
    <td><p>Papua New Guinea.</p></td>
  </tr>
  <tr>
    <td><strong>PH</strong></td>
    <td><p>Philippines.</p></td>
  </tr>
  <tr>
    <td><strong>PK</strong></td>
    <td><p>Pakistan.</p></td>
  </tr>
  <tr>
    <td><strong>PL</strong></td>
    <td><p>Poland.</p></td>
  </tr>
  <tr>
    <td><strong>PM</strong></td>
    <td><p>St. Pierre &amp; Miquelon.</p></td>
  </tr>
  <tr>
    <td><strong>PN</strong></td>
    <td><p>Pitcairn Islands.</p></td>
  </tr>
  <tr>
    <td><strong>PS</strong></td>
    <td><p>Palestinian Territories.</p></td>
  </tr>
  <tr>
    <td><strong>PT</strong></td>
    <td><p>Portugal.</p></td>
  </tr>
  <tr>
    <td><strong>PY</strong></td>
    <td><p>Paraguay.</p></td>
  </tr>
  <tr>
    <td><strong>QA</strong></td>
    <td><p>Qatar.</p></td>
  </tr>
  <tr>
    <td><strong>RE</strong></td>
    <td><p>Réunion.</p></td>
  </tr>
  <tr>
    <td><strong>RO</strong></td>
    <td><p>Romania.</p></td>
  </tr>
  <tr>
    <td><strong>RS</strong></td>
    <td><p>Serbia.</p></td>
  </tr>
  <tr>
    <td><strong>RU</strong></td>
    <td><p>Russia.</p></td>
  </tr>
  <tr>
    <td><strong>RW</strong></td>
    <td><p>Rwanda.</p></td>
  </tr>
  <tr>
    <td><strong>SA</strong></td>
    <td><p>Saudi Arabia.</p></td>
  </tr>
  <tr>
    <td><strong>SB</strong></td>
    <td><p>Solomon Islands.</p></td>
  </tr>
  <tr>
    <td><strong>SC</strong></td>
    <td><p>Seychelles.</p></td>
  </tr>
  <tr>
    <td><strong>SD</strong></td>
    <td><p>Sudan.</p></td>
  </tr>
  <tr>
    <td><strong>SE</strong></td>
    <td><p>Sweden.</p></td>
  </tr>
  <tr>
    <td><strong>SG</strong></td>
    <td><p>Singapore.</p></td>
  </tr>
  <tr>
    <td><strong>SH</strong></td>
    <td><p>St. Helena.</p></td>
  </tr>
  <tr>
    <td><strong>SI</strong></td>
    <td><p>Slovenia.</p></td>
  </tr>
  <tr>
    <td><strong>SJ</strong></td>
    <td><p>Svalbard &amp; Jan Mayen.</p></td>
  </tr>
  <tr>
    <td><strong>SK</strong></td>
    <td><p>Slovakia.</p></td>
  </tr>
  <tr>
    <td><strong>SL</strong></td>
    <td><p>Sierra Leone.</p></td>
  </tr>
  <tr>
    <td><strong>SM</strong></td>
    <td><p>San Marino.</p></td>
  </tr>
  <tr>
    <td><strong>SN</strong></td>
    <td><p>Senegal.</p></td>
  </tr>
  <tr>
    <td><strong>SO</strong></td>
    <td><p>Somalia.</p></td>
  </tr>
  <tr>
    <td><strong>SR</strong></td>
    <td><p>Suriname.</p></td>
  </tr>
  <tr>
    <td><strong>SS</strong></td>
    <td><p>South Sudan.</p></td>
  </tr>
  <tr>
    <td><strong>ST</strong></td>
    <td><p>São Tomé &amp; Príncipe.</p></td>
  </tr>
  <tr>
    <td><strong>SV</strong></td>
    <td><p>El Salvador.</p></td>
  </tr>
  <tr>
    <td><strong>SX</strong></td>
    <td><p>Sint Maarten.</p></td>
  </tr>
  <tr>
    <td><strong>SY</strong></td>
    <td><p>Syria.</p></td>
  </tr>
  <tr>
    <td><strong>SZ</strong></td>
    <td><p>Eswatini.</p></td>
  </tr>
  <tr>
    <td><strong>TA</strong></td>
    <td><p>Tristan da Cunha.</p></td>
  </tr>
  <tr>
    <td><strong>TC</strong></td>
    <td><p>Turks &amp; Caicos Islands.</p></td>
  </tr>
  <tr>
    <td><strong>TD</strong></td>
    <td><p>Chad.</p></td>
  </tr>
  <tr>
    <td><strong>TF</strong></td>
    <td><p>French Southern Territories.</p></td>
  </tr>
  <tr>
    <td><strong>TG</strong></td>
    <td><p>Togo.</p></td>
  </tr>
  <tr>
    <td><strong>TH</strong></td>
    <td><p>Thailand.</p></td>
  </tr>
  <tr>
    <td><strong>TJ</strong></td>
    <td><p>Tajikistan.</p></td>
  </tr>
  <tr>
    <td><strong>TK</strong></td>
    <td><p>Tokelau.</p></td>
  </tr>
  <tr>
    <td><strong>TL</strong></td>
    <td><p>Timor-Leste.</p></td>
  </tr>
  <tr>
    <td><strong>TM</strong></td>
    <td><p>Turkmenistan.</p></td>
  </tr>
  <tr>
    <td><strong>TN</strong></td>
    <td><p>Tunisia.</p></td>
  </tr>
  <tr>
    <td><strong>TO</strong></td>
    <td><p>Tonga.</p></td>
  </tr>
  <tr>
    <td><strong>TR</strong></td>
    <td><p>Turkey.</p></td>
  </tr>
  <tr>
    <td><strong>TT</strong></td>
    <td><p>Trinidad &amp; Tobago.</p></td>
  </tr>
  <tr>
    <td><strong>TV</strong></td>
    <td><p>Tuvalu.</p></td>
  </tr>
  <tr>
    <td><strong>TW</strong></td>
    <td><p>Taiwan.</p></td>
  </tr>
  <tr>
    <td><strong>TZ</strong></td>
    <td><p>Tanzania.</p></td>
  </tr>
  <tr>
    <td><strong>UA</strong></td>
    <td><p>Ukraine.</p></td>
  </tr>
  <tr>
    <td><strong>UG</strong></td>
    <td><p>Uganda.</p></td>
  </tr>
  <tr>
    <td><strong>UM</strong></td>
    <td><p>U.S. Outlying Islands.</p></td>
  </tr>
  <tr>
    <td><strong>US</strong></td>
    <td><p>United States.</p></td>
  </tr>
  <tr>
    <td><strong>UY</strong></td>
    <td><p>Uruguay.</p></td>
  </tr>
  <tr>
    <td><strong>UZ</strong></td>
    <td><p>Uzbekistan.</p></td>
  </tr>
  <tr>
    <td><strong>VA</strong></td>
    <td><p>Vatican City.</p></td>
  </tr>
  <tr>
    <td><strong>VC</strong></td>
    <td><p>St. Vincent &amp; Grenadines.</p></td>
  </tr>
  <tr>
    <td><strong>VE</strong></td>
    <td><p>Venezuela.</p></td>
  </tr>
  <tr>
    <td><strong>VG</strong></td>
    <td><p>British Virgin Islands.</p></td>
  </tr>
  <tr>
    <td><strong>VN</strong></td>
    <td><p>Vietnam.</p></td>
  </tr>
  <tr>
    <td><strong>VU</strong></td>
    <td><p>Vanuatu.</p></td>
  </tr>
  <tr>
    <td><strong>WF</strong></td>
    <td><p>Wallis &amp; Futuna.</p></td>
  </tr>
  <tr>
    <td><strong>WS</strong></td>
    <td><p>Samoa.</p></td>
  </tr>
  <tr>
    <td><strong>XK</strong></td>
    <td><p>Kosovo.</p></td>
  </tr>
  <tr>
    <td><strong>YE</strong></td>
    <td><p>Yemen.</p></td>
  </tr>
  <tr>
    <td><strong>YT</strong></td>
    <td><p>Mayotte.</p></td>
  </tr>
  <tr>
    <td><strong>ZA</strong></td>
    <td><p>South Africa.</p></td>
  </tr>
  <tr>
    <td><strong>ZM</strong></td>
    <td><p>Zambia.</p></td>
  </tr>
  <tr>
    <td><strong>ZW</strong></td>
    <td><p>Zimbabwe.</p></td>
  </tr>
  <tr>
    <td><strong>ZZ</strong></td>
    <td><p>Unknown Region.</p></td>
  </tr>
</table>

---

### CropRegion

<p>The part of the image that should remain after cropping.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BOTTOM</strong></td>
    <td><p>Keep the bottom of the image.</p></td>
  </tr>
  <tr>
    <td><strong>CENTER</strong></td>
    <td><p>Keep the center of the image.</p></td>
  </tr>
  <tr>
    <td><strong>LEFT</strong></td>
    <td><p>Keep the left of the image.</p></td>
  </tr>
  <tr>
    <td><strong>RIGHT</strong></td>
    <td><p>Keep the right of the image.</p></td>
  </tr>
  <tr>
    <td><strong>TOP</strong></td>
    <td><p>Keep the top of the image.</p></td>
  </tr>
</table>

---

### CurrencyCode

<p>Currency codes.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AED</strong></td>
    <td><p>United Arab Emirates Dirham (AED).</p></td>
  </tr>
  <tr>
    <td><strong>AFN</strong></td>
    <td><p>Afghan Afghani (AFN).</p></td>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>Albanian Lek (ALL).</p></td>
  </tr>
  <tr>
    <td><strong>AMD</strong></td>
    <td><p>Armenian Dram (AMD).</p></td>
  </tr>
  <tr>
    <td><strong>ANG</strong></td>
    <td><p>Netherlands Antillean Guilder.</p></td>
  </tr>
  <tr>
    <td><strong>AOA</strong></td>
    <td><p>Angolan Kwanza (AOA).</p></td>
  </tr>
  <tr>
    <td><strong>ARS</strong></td>
    <td><p>Argentine Pesos (ARS).</p></td>
  </tr>
  <tr>
    <td><strong>AUD</strong></td>
    <td><p>Australian Dollars (AUD).</p></td>
  </tr>
  <tr>
    <td><strong>AWG</strong></td>
    <td><p>Aruban Florin (AWG).</p></td>
  </tr>
  <tr>
    <td><strong>AZN</strong></td>
    <td><p>Azerbaijani Manat (AZN).</p></td>
  </tr>
  <tr>
    <td><strong>BAM</strong></td>
    <td><p>Bosnia and Herzegovina Convertible Mark (BAM).</p></td>
  </tr>
  <tr>
    <td><strong>BBD</strong></td>
    <td><p>Barbadian Dollar (BBD).</p></td>
  </tr>
  <tr>
    <td><strong>BDT</strong></td>
    <td><p>Bangladesh Taka (BDT).</p></td>
  </tr>
  <tr>
    <td><strong>BGN</strong></td>
    <td><p>Bulgarian Lev (BGN).</p></td>
  </tr>
  <tr>
    <td><strong>BHD</strong></td>
    <td><p>Bahraini Dinar (BHD).</p></td>
  </tr>
  <tr>
    <td><strong>BIF</strong></td>
    <td><p>Burundian Franc (BIF).</p></td>
  </tr>
  <tr>
    <td><strong>BMD</strong></td>
    <td><p>Bermudian Dollar (BMD).</p></td>
  </tr>
  <tr>
    <td><strong>BND</strong></td>
    <td><p>Brunei Dollar (BND).</p></td>
  </tr>
  <tr>
    <td><strong>BOB</strong></td>
    <td><p>Bolivian Boliviano (BOB).</p></td>
  </tr>
  <tr>
    <td><strong>BRL</strong></td>
    <td><p>Brazilian Real (BRL).</p></td>
  </tr>
  <tr>
    <td><strong>BSD</strong></td>
    <td><p>Bahamian Dollar (BSD).</p></td>
  </tr>
  <tr>
    <td><strong>BTN</strong></td>
    <td><p>Bhutanese Ngultrum (BTN).</p></td>
  </tr>
  <tr>
    <td><strong>BWP</strong></td>
    <td><p>Botswana Pula (BWP).</p></td>
  </tr>
  <tr>
    <td><strong>BYN</strong></td>
    <td><p>Belarusian Ruble (BYN).</p></td>
  </tr>
  <tr>
    <td><strong>BYR</strong></td>
    <td><p>Belarusian Ruble (BYR).</p></td>
  </tr>
  <tr>
    <td><strong>BZD</strong></td>
    <td><p>Belize Dollar (BZD).</p></td>
  </tr>
  <tr>
    <td><strong>CAD</strong></td>
    <td><p>Canadian Dollars (CAD).</p></td>
  </tr>
  <tr>
    <td><strong>CDF</strong></td>
    <td><p>Congolese franc (CDF).</p></td>
  </tr>
  <tr>
    <td><strong>CHF</strong></td>
    <td><p>Swiss Francs (CHF).</p></td>
  </tr>
  <tr>
    <td><strong>CLP</strong></td>
    <td><p>Chilean Peso (CLP).</p></td>
  </tr>
  <tr>
    <td><strong>CNY</strong></td>
    <td><p>Chinese Yuan Renminbi (CNY).</p></td>
  </tr>
  <tr>
    <td><strong>COP</strong></td>
    <td><p>Colombian Peso (COP).</p></td>
  </tr>
  <tr>
    <td><strong>CRC</strong></td>
    <td><p>Costa Rican Colones (CRC).</p></td>
  </tr>
  <tr>
    <td><strong>CVE</strong></td>
    <td><p>Cape Verdean escudo (CVE).</p></td>
  </tr>
  <tr>
    <td><strong>CZK</strong></td>
    <td><p>Czech Koruny (CZK).</p></td>
  </tr>
  <tr>
    <td><strong>DJF</strong></td>
    <td><p>Djiboutian Franc (DJF).</p></td>
  </tr>
  <tr>
    <td><strong>DKK</strong></td>
    <td><p>Danish Kroner (DKK).</p></td>
  </tr>
  <tr>
    <td><strong>DOP</strong></td>
    <td><p>Dominican Peso (DOP).</p></td>
  </tr>
  <tr>
    <td><strong>DZD</strong></td>
    <td><p>Algerian Dinar (DZD).</p></td>
  </tr>
  <tr>
    <td><strong>EGP</strong></td>
    <td><p>Egyptian Pound (EGP).</p></td>
  </tr>
  <tr>
    <td><strong>ERN</strong></td>
    <td><p>Eritrean Nakfa (ERN).</p></td>
  </tr>
  <tr>
    <td><strong>ETB</strong></td>
    <td><p>Ethiopian Birr (ETB).</p></td>
  </tr>
  <tr>
    <td><strong>EUR</strong></td>
    <td><p>Euro (EUR).</p></td>
  </tr>
  <tr>
    <td><strong>FJD</strong></td>
    <td><p>Fijian Dollars (FJD).</p></td>
  </tr>
  <tr>
    <td><strong>FKP</strong></td>
    <td><p>Falkland Islands Pounds (FKP).</p></td>
  </tr>
  <tr>
    <td><strong>GBP</strong></td>
    <td><p>United Kingdom Pounds (GBP).</p></td>
  </tr>
  <tr>
    <td><strong>GEL</strong></td>
    <td><p>Georgian Lari (GEL).</p></td>
  </tr>
  <tr>
    <td><strong>GHS</strong></td>
    <td><p>Ghanaian Cedi (GHS).</p></td>
  </tr>
  <tr>
    <td><strong>GIP</strong></td>
    <td><p>Gibraltar Pounds (GIP).</p></td>
  </tr>
  <tr>
    <td><strong>GMD</strong></td>
    <td><p>Gambian Dalasi (GMD).</p></td>
  </tr>
  <tr>
    <td><strong>GNF</strong></td>
    <td><p>Guinean Franc (GNF).</p></td>
  </tr>
  <tr>
    <td><strong>GTQ</strong></td>
    <td><p>Guatemalan Quetzal (GTQ).</p></td>
  </tr>
  <tr>
    <td><strong>GYD</strong></td>
    <td><p>Guyanese Dollar (GYD).</p></td>
  </tr>
  <tr>
    <td><strong>HKD</strong></td>
    <td><p>Hong Kong Dollars (HKD).</p></td>
  </tr>
  <tr>
    <td><strong>HNL</strong></td>
    <td><p>Honduran Lempira (HNL).</p></td>
  </tr>
  <tr>
    <td><strong>HRK</strong></td>
    <td><p>Croatian Kuna (HRK).</p></td>
  </tr>
  <tr>
    <td><strong>HTG</strong></td>
    <td><p>Haitian Gourde (HTG).</p></td>
  </tr>
  <tr>
    <td><strong>HUF</strong></td>
    <td><p>Hungarian Forint (HUF).</p></td>
  </tr>
  <tr>
    <td><strong>IDR</strong></td>
    <td><p>Indonesian Rupiah (IDR).</p></td>
  </tr>
  <tr>
    <td><strong>ILS</strong></td>
    <td><p>Israeli New Shekel (NIS).</p></td>
  </tr>
  <tr>
    <td><strong>INR</strong></td>
    <td><p>Indian Rupees (INR).</p></td>
  </tr>
  <tr>
    <td><strong>IQD</strong></td>
    <td><p>Iraqi Dinar (IQD).</p></td>
  </tr>
  <tr>
    <td><strong>IRR</strong></td>
    <td><p>Iranian Rial (IRR).</p></td>
  </tr>
  <tr>
    <td><strong>ISK</strong></td>
    <td><p>Icelandic Kronur (ISK).</p></td>
  </tr>
  <tr>
    <td><strong>JEP</strong></td>
    <td><p>Jersey Pound.</p></td>
  </tr>
  <tr>
    <td><strong>JMD</strong></td>
    <td><p>Jamaican Dollars (JMD).</p></td>
  </tr>
  <tr>
    <td><strong>JOD</strong></td>
    <td><p>Jordanian Dinar (JOD).</p></td>
  </tr>
  <tr>
    <td><strong>JPY</strong></td>
    <td><p>Japanese Yen (JPY).</p></td>
  </tr>
  <tr>
    <td><strong>KES</strong></td>
    <td><p>Kenyan Shilling (KES).</p></td>
  </tr>
  <tr>
    <td><strong>KGS</strong></td>
    <td><p>Kyrgyzstani Som (KGS).</p></td>
  </tr>
  <tr>
    <td><strong>KHR</strong></td>
    <td><p>Cambodian Riel.</p></td>
  </tr>
  <tr>
    <td><strong>KID</strong></td>
    <td><p>Kiribati Dollar (KID).</p></td>
  </tr>
  <tr>
    <td><strong>KMF</strong></td>
    <td><p>Comorian Franc (KMF).</p></td>
  </tr>
  <tr>
    <td><strong>KRW</strong></td>
    <td><p>South Korean Won (KRW).</p></td>
  </tr>
  <tr>
    <td><strong>KWD</strong></td>
    <td><p>Kuwaiti Dinar (KWD).</p></td>
  </tr>
  <tr>
    <td><strong>KYD</strong></td>
    <td><p>Cayman Dollars (KYD).</p></td>
  </tr>
  <tr>
    <td><strong>KZT</strong></td>
    <td><p>Kazakhstani Tenge (KZT).</p></td>
  </tr>
  <tr>
    <td><strong>LAK</strong></td>
    <td><p>Laotian Kip (LAK).</p></td>
  </tr>
  <tr>
    <td><strong>LBP</strong></td>
    <td><p>Lebanese Pounds (LBP).</p></td>
  </tr>
  <tr>
    <td><strong>LKR</strong></td>
    <td><p>Sri Lankan Rupees (LKR).</p></td>
  </tr>
  <tr>
    <td><strong>LRD</strong></td>
    <td><p>Liberian Dollar (LRD).</p></td>
  </tr>
  <tr>
    <td><strong>LSL</strong></td>
    <td><p>Lesotho Loti (LSL).</p></td>
  </tr>
  <tr>
    <td><strong>LTL</strong></td>
    <td><p>Lithuanian Litai (LTL).</p></td>
  </tr>
  <tr>
    <td><strong>LVL</strong></td>
    <td><p>Latvian Lati (LVL).</p></td>
  </tr>
  <tr>
    <td><strong>LYD</strong></td>
    <td><p>Libyan Dinar (LYD).</p></td>
  </tr>
  <tr>
    <td><strong>MAD</strong></td>
    <td><p>Moroccan Dirham.</p></td>
  </tr>
  <tr>
    <td><strong>MDL</strong></td>
    <td><p>Moldovan Leu (MDL).</p></td>
  </tr>
  <tr>
    <td><strong>MGA</strong></td>
    <td><p>Malagasy Ariary (MGA).</p></td>
  </tr>
  <tr>
    <td><strong>MKD</strong></td>
    <td><p>Macedonia Denar (MKD).</p></td>
  </tr>
  <tr>
    <td><strong>MMK</strong></td>
    <td><p>Burmese Kyat (MMK).</p></td>
  </tr>
  <tr>
    <td><strong>MNT</strong></td>
    <td><p>Mongolian Tugrik.</p></td>
  </tr>
  <tr>
    <td><strong>MOP</strong></td>
    <td><p>Macanese Pataca (MOP).</p></td>
  </tr>
  <tr>
    <td><strong>MRU</strong></td>
    <td><p>Mauritanian Ouguiya (MRU).</p></td>
  </tr>
  <tr>
    <td><strong>MUR</strong></td>
    <td><p>Mauritian Rupee (MUR).</p></td>
  </tr>
  <tr>
    <td><strong>MVR</strong></td>
    <td><p>Maldivian Rufiyaa (MVR).</p></td>
  </tr>
  <tr>
    <td><strong>MWK</strong></td>
    <td><p>Malawian Kwacha (MWK).</p></td>
  </tr>
  <tr>
    <td><strong>MXN</strong></td>
    <td><p>Mexican Pesos (MXN).</p></td>
  </tr>
  <tr>
    <td><strong>MYR</strong></td>
    <td><p>Malaysian Ringgits (MYR).</p></td>
  </tr>
  <tr>
    <td><strong>MZN</strong></td>
    <td><p>Mozambican Metical.</p></td>
  </tr>
  <tr>
    <td><strong>NAD</strong></td>
    <td><p>Namibian Dollar.</p></td>
  </tr>
  <tr>
    <td><strong>NGN</strong></td>
    <td><p>Nigerian Naira (NGN).</p></td>
  </tr>
  <tr>
    <td><strong>NIO</strong></td>
    <td><p>Nicaraguan Córdoba (NIO).</p></td>
  </tr>
  <tr>
    <td><strong>NOK</strong></td>
    <td><p>Norwegian Kroner (NOK).</p></td>
  </tr>
  <tr>
    <td><strong>NPR</strong></td>
    <td><p>Nepalese Rupee (NPR).</p></td>
  </tr>
  <tr>
    <td><strong>NZD</strong></td>
    <td><p>New Zealand Dollars (NZD).</p></td>
  </tr>
  <tr>
    <td><strong>OMR</strong></td>
    <td><p>Omani Rial (OMR).</p></td>
  </tr>
  <tr>
    <td><strong>PAB</strong></td>
    <td><p>Panamian Balboa (PAB).</p></td>
  </tr>
  <tr>
    <td><strong>PEN</strong></td>
    <td><p>Peruvian Nuevo Sol (PEN).</p></td>
  </tr>
  <tr>
    <td><strong>PGK</strong></td>
    <td><p>Papua New Guinean Kina (PGK).</p></td>
  </tr>
  <tr>
    <td><strong>PHP</strong></td>
    <td><p>Philippine Peso (PHP).</p></td>
  </tr>
  <tr>
    <td><strong>PKR</strong></td>
    <td><p>Pakistani Rupee (PKR).</p></td>
  </tr>
  <tr>
    <td><strong>PLN</strong></td>
    <td><p>Polish Zlotych (PLN).</p></td>
  </tr>
  <tr>
    <td><strong>PYG</strong></td>
    <td><p>Paraguayan Guarani (PYG).</p></td>
  </tr>
  <tr>
    <td><strong>QAR</strong></td>
    <td><p>Qatari Rial (QAR).</p></td>
  </tr>
  <tr>
    <td><strong>RON</strong></td>
    <td><p>Romanian Lei (RON).</p></td>
  </tr>
  <tr>
    <td><strong>RSD</strong></td>
    <td><p>Serbian dinar (RSD).</p></td>
  </tr>
  <tr>
    <td><strong>RUB</strong></td>
    <td><p>Russian Rubles (RUB).</p></td>
  </tr>
  <tr>
    <td><strong>RWF</strong></td>
    <td><p>Rwandan Franc (RWF).</p></td>
  </tr>
  <tr>
    <td><strong>SAR</strong></td>
    <td><p>Saudi Riyal (SAR).</p></td>
  </tr>
  <tr>
    <td><strong>SBD</strong></td>
    <td><p>Solomon Islands Dollar (SBD).</p></td>
  </tr>
  <tr>
    <td><strong>SCR</strong></td>
    <td><p>Seychellois Rupee (SCR).</p></td>
  </tr>
  <tr>
    <td><strong>SDG</strong></td>
    <td><p>Sudanese Pound (SDG).</p></td>
  </tr>
  <tr>
    <td><strong>SEK</strong></td>
    <td><p>Swedish Kronor (SEK).</p></td>
  </tr>
  <tr>
    <td><strong>SGD</strong></td>
    <td><p>Singapore Dollars (SGD).</p></td>
  </tr>
  <tr>
    <td><strong>SHP</strong></td>
    <td><p>Saint Helena Pounds (SHP).</p></td>
  </tr>
  <tr>
    <td><strong>SLL</strong></td>
    <td><p>Sierra Leonean Leone (SLL).</p></td>
  </tr>
  <tr>
    <td><strong>SOS</strong></td>
    <td><p>Somali Shilling (SOS).</p></td>
  </tr>
  <tr>
    <td><strong>SRD</strong></td>
    <td><p>Surinamese Dollar (SRD).</p></td>
  </tr>
  <tr>
    <td><strong>SSP</strong></td>
    <td><p>South Sudanese Pound (SSP).</p></td>
  </tr>
  <tr>
    <td><strong>STD</strong></td>
    <td><p>Sao Tome And Principe Dobra (STD).</p></td>
  </tr>
  <tr>
    <td><strong>SYP</strong></td>
    <td><p>Syrian Pound (SYP).</p></td>
  </tr>
  <tr>
    <td><strong>SZL</strong></td>
    <td><p>Swazi Lilangeni (SZL).</p></td>
  </tr>
  <tr>
    <td><strong>THB</strong></td>
    <td><p>Thai baht (THB).</p></td>
  </tr>
  <tr>
    <td><strong>TJS</strong></td>
    <td><p>Tajikistani Somoni (TJS).</p></td>
  </tr>
  <tr>
    <td><strong>TMT</strong></td>
    <td><p>Turkmenistani Manat (TMT).</p></td>
  </tr>
  <tr>
    <td><strong>TND</strong></td>
    <td><p>Tunisian Dinar (TND).</p></td>
  </tr>
  <tr>
    <td><strong>TOP</strong></td>
    <td><p>Tongan Pa&rsquo;anga (TOP).</p></td>
  </tr>
  <tr>
    <td><strong>TRY</strong></td>
    <td><p>Turkish Lira (TRY).</p></td>
  </tr>
  <tr>
    <td><strong>TTD</strong></td>
    <td><p>Trinidad and Tobago Dollars (TTD).</p></td>
  </tr>
  <tr>
    <td><strong>TWD</strong></td>
    <td><p>Taiwan Dollars (TWD).</p></td>
  </tr>
  <tr>
    <td><strong>TZS</strong></td>
    <td><p>Tanzanian Shilling (TZS).</p></td>
  </tr>
  <tr>
    <td><strong>UAH</strong></td>
    <td><p>Ukrainian Hryvnia (UAH).</p></td>
  </tr>
  <tr>
    <td><strong>UGX</strong></td>
    <td><p>Ugandan Shilling (UGX).</p></td>
  </tr>
  <tr>
    <td><strong>USD</strong></td>
    <td><p>United States Dollars (USD).</p></td>
  </tr>
  <tr>
    <td><strong>UYU</strong></td>
    <td><p>Uruguayan Pesos (UYU).</p></td>
  </tr>
  <tr>
    <td><strong>UZS</strong></td>
    <td><p>Uzbekistan som (UZS).</p></td>
  </tr>
  <tr>
    <td><strong>VEF</strong></td>
    <td><p>Venezuelan Bolivares (VEF).</p></td>
  </tr>
  <tr>
    <td><strong>VES</strong></td>
    <td><p>Venezuelan Bolivares (VES).</p></td>
  </tr>
  <tr>
    <td><strong>VND</strong></td>
    <td><p>Vietnamese đồng (VND).</p></td>
  </tr>
  <tr>
    <td><strong>VUV</strong></td>
    <td><p>Vanuatu Vatu (VUV).</p></td>
  </tr>
  <tr>
    <td><strong>WST</strong></td>
    <td><p>Samoan Tala (WST).</p></td>
  </tr>
  <tr>
    <td><strong>XAF</strong></td>
    <td><p>Central African CFA Franc (XAF).</p></td>
  </tr>
  <tr>
    <td><strong>XCD</strong></td>
    <td><p>East Caribbean Dollar (XCD).</p></td>
  </tr>
  <tr>
    <td><strong>XOF</strong></td>
    <td><p>West African CFA franc (XOF).</p></td>
  </tr>
  <tr>
    <td><strong>XPF</strong></td>
    <td><p>CFP Franc (XPF).</p></td>
  </tr>
  <tr>
    <td><strong>YER</strong></td>
    <td><p>Yemeni Rial (YER).</p></td>
  </tr>
  <tr>
    <td><strong>ZAR</strong></td>
    <td><p>South African Rand (ZAR).</p></td>
  </tr>
  <tr>
    <td><strong>ZMW</strong></td>
    <td><p>Zambian Kwacha (ZMW).</p></td>
  </tr>
</table>

---

### CustomerErrorCode

<p>Possible error codes that could be returned by CustomerUserError.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALREADY_ENABLED</strong></td>
    <td><p>Customer already enabled.</p></td>
  </tr>
  <tr>
    <td><strong>BAD_DOMAIN</strong></td>
    <td><p>Input email contains an invalid domain name.</p></td>
  </tr>
  <tr>
    <td><strong>BLANK</strong></td>
    <td><p>Input value is blank.</p></td>
  </tr>
  <tr>
    <td><strong>CONTAINS_HTML_TAGS</strong></td>
    <td><p>Input contains HTML tags.</p></td>
  </tr>
  <tr>
    <td><strong>CONTAINS_URL</strong></td>
    <td><p>Input contains URL.</p></td>
  </tr>
  <tr>
    <td><strong>CUSTOMER_DISABLED</strong></td>
    <td><p>Customer is disabled.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID</strong></td>
    <td><p>Input value is invalid.</p></td>
  </tr>
  <tr>
    <td><strong>INVALID_MULTIPASS_REQUEST</strong></td>
    <td><p>Multipass token is not valid.</p></td>
  </tr>
  <tr>
    <td><strong>NOT_FOUND</strong></td>
    <td><p>Address does not exist.</p></td>
  </tr>
  <tr>
    <td><strong>PASSWORD_STARTS_OR_ENDS_WITH_WHITESPACE</strong></td>
    <td><p>Input password starts or ends with whitespace.</p></td>
  </tr>
  <tr>
    <td><strong>TAKEN</strong></td>
    <td><p>Input value is already taken.</p></td>
  </tr>
  <tr>
    <td><strong>TOKEN_INVALID</strong></td>
    <td><p>Invalid activation token.</p></td>
  </tr>
  <tr>
    <td><strong>TOO_LONG</strong></td>
    <td><p>Input value is too long.</p></td>
  </tr>
  <tr>
    <td><strong>TOO_SHORT</strong></td>
    <td><p>Input value is too short.</p></td>
  </tr>
  <tr>
    <td><strong>UNIDENTIFIED_CUSTOMER</strong></td>
    <td><p>Unidentified customer.</p></td>
  </tr>
</table>

---

### DigitalWallet

<p>Digital wallet, such as Apple Pay, which can be used for accelerated checkouts.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ANDROID_PAY</strong></td>
    <td><p>Android Pay.</p></td>
  </tr>
  <tr>
    <td><strong>APPLE_PAY</strong></td>
    <td><p>Apple Pay.</p></td>
  </tr>
  <tr>
    <td><strong>GOOGLE_PAY</strong></td>
    <td><p>Google Pay.</p></td>
  </tr>
  <tr>
    <td><strong>SHOPIFY_PAY</strong></td>
    <td><p>Shopify Pay.</p></td>
  </tr>
</table>

---

### DiscountApplicationAllocationMethod

<p>The method by which the discount&rsquo;s value is allocated onto its entitled lines.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ACROSS</strong></td>
    <td><p>The value is spread across all entitled lines.</p></td>
  </tr>
  <tr>
    <td><strong>EACH</strong></td>
    <td><p>The value is applied onto every entitled line.</p></td>
  </tr>
  <tr>
    <td><strong>ONE</strong></td>
    <td><p>The value is specifically applied onto a particular line.</p></td>
  </tr>
</table>

---

### DiscountApplicationTargetSelection

<p>Which lines on the order that the discount is allocated over, of the type
defined by the Discount Application&rsquo;s target_type.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>The discount is allocated onto all the lines.</p></td>
  </tr>
  <tr>
    <td><strong>ENTITLED</strong></td>
    <td><p>The discount is allocated onto only the lines it is entitled for.</p></td>
  </tr>
  <tr>
    <td><strong>EXPLICIT</strong></td>
    <td><p>The discount is allocated onto explicitly chosen lines.</p></td>
  </tr>
</table>

---

### DiscountApplicationTargetType

<p>The type of line (i.e. line item or shipping line) on an order that the discount is applicable towards.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>LINE_ITEM</strong></td>
    <td><p>The discount applies onto line items.</p></td>
  </tr>
  <tr>
    <td><strong>SHIPPING_LINE</strong></td>
    <td><p>The discount applies onto shipping lines.</p></td>
  </tr>
</table>

---

### ImageContentType

<p>List of supported image content types.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>JPG</strong></td>
    <td><p>A JPG image.</p></td>
  </tr>
  <tr>
    <td><strong>PNG</strong></td>
    <td><p>A PNG image.</p></td>
  </tr>
  <tr>
    <td><strong>WEBP</strong></td>
    <td><p>A WEBP image.</p></td>
  </tr>
</table>

---

### MediaContentType

<p>The possible content types for a media object.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>EXTERNAL_VIDEO</strong></td>
    <td><p>An externally hosted video.</p></td>
  </tr>
  <tr>
    <td><strong>IMAGE</strong></td>
    <td><p>A Shopify hosted image.</p></td>
  </tr>
  <tr>
    <td><strong>MODEL_3D</strong></td>
    <td><p>A 3d model.</p></td>
  </tr>
  <tr>
    <td><strong>VIDEO</strong></td>
    <td><p>A Shopify hosted video.</p></td>
  </tr>
</table>

---

### MediaHost

<p>Host for a Media Resource.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>VIMEO</strong></td>
    <td><p>Host for Vimeo embedded videos.</p></td>
  </tr>
  <tr>
    <td><strong>YOUTUBE</strong></td>
    <td><p>Host for YouTube embedded videos.</p></td>
  </tr>
</table>

---

### MetafieldValueType

<p>Metafield value types.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTEGER</strong></td>
    <td><p>An integer metafield.</p></td>
  </tr>
  <tr>
    <td><strong>JSON_STRING</strong></td>
    <td><p>A json string metafield.</p></td>
  </tr>
  <tr>
    <td><strong>STRING</strong></td>
    <td><p>A string metafield.</p></td>
  </tr>
</table>

---

### OrderCancelReason

<p>Represents the reason for the order&rsquo;s cancellation.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CUSTOMER</strong></td>
    <td><p>The customer wanted to cancel the order.</p></td>
  </tr>
  <tr>
    <td><strong>DECLINED</strong></td>
    <td><p>Payment was declined.</p></td>
  </tr>
  <tr>
    <td><strong>FRAUD</strong></td>
    <td><p>The order was fraudulent.</p></td>
  </tr>
  <tr>
    <td><strong>INVENTORY</strong></td>
    <td><p>There was insufficient inventory.</p></td>
  </tr>
  <tr>
    <td><strong>OTHER</strong></td>
    <td><p>The order was canceled for an unlisted reason.</p></td>
  </tr>
</table>

---

### OrderFinancialStatus

<p>Represents the order&rsquo;s current financial status.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AUTHORIZED</strong></td>
    <td><p>Displayed as <strong>Authorized</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>PAID</strong></td>
    <td><p>Displayed as <strong>Paid</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>PARTIALLY_PAID</strong></td>
    <td><p>Displayed as <strong>Partially paid</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>PARTIALLY_REFUNDED</strong></td>
    <td><p>Displayed as <strong>Partially refunded</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>Displayed as <strong>Pending</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>REFUNDED</strong></td>
    <td><p>Displayed as <strong>Refunded</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>VOIDED</strong></td>
    <td><p>Displayed as <strong>Voided</strong>.</p></td>
  </tr>
</table>

---

### OrderFulfillmentStatus

<p>Represents the order&rsquo;s current fulfillment status.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>FULFILLED</strong></td>
    <td><p>Displayed as <strong>Fulfilled</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>IN_PROGRESS</strong></td>
    <td><p>Displayed as <strong>In progress</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>OPEN</strong></td>
    <td><p>Displayed as <strong>Open</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>PARTIALLY_FULFILLED</strong></td>
    <td><p>Displayed as <strong>Partially fulfilled</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING_FULFILLMENT</strong></td>
    <td><p>Displayed as <strong>Pending fulfillment</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>RESTOCKED</strong></td>
    <td><p>Displayed as <strong>Restocked</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>SCHEDULED</strong></td>
    <td><p>Displayed as <strong>Scheduled</strong>.</p></td>
  </tr>
  <tr>
    <td><strong>UNFULFILLED</strong></td>
    <td><p>Displayed as <strong>Unfulfilled</strong>.</p></td>
  </tr>
</table>

---

### OrderSortKeys

<p>The set of valid sort keys for the Order query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>PROCESSED_AT</strong></td>
    <td><p>Sort by the <code>processed_at</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>TOTAL_PRICE</strong></td>
    <td><p>Sort by the <code>total_price</code> value.</p></td>
  </tr>
</table>

---

### PageSortKeys

<p>The set of valid sort keys for the Page query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>TITLE</strong></td>
    <td><p>Sort by the <code>title</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Sort by the <code>updated_at</code> value.</p></td>
  </tr>
</table>

---

### PaymentTokenType

<p>The valid values for the types of payment token.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>APPLE_PAY</strong></td>
    <td><p>Apple Pay token type.</p></td>
  </tr>
  <tr>
    <td><strong>GOOGLE_PAY</strong></td>
    <td><p>Google Pay token type.</p></td>
  </tr>
  <tr>
    <td><strong>SHOPIFY_PAY</strong></td>
    <td><p>Shopify Pay token type.</p></td>
  </tr>
  <tr>
    <td><strong>VAULT</strong></td>
    <td><p>Vault payment token type.</p></td>
  </tr>
</table>

---

### ProductCollectionSortKeys

<p>The set of valid sort keys for the ProductCollection query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BEST_SELLING</strong></td>
    <td><p>Sort by the <code>best-selling</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>COLLECTION_DEFAULT</strong></td>
    <td><p>Sort by the <code>collection-default</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>CREATED</strong></td>
    <td><p>Sort by the <code>created</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>MANUAL</strong></td>
    <td><p>Sort by the <code>manual</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>PRICE</strong></td>
    <td><p>Sort by the <code>price</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>TITLE</strong></td>
    <td><p>Sort by the <code>title</code> value.</p></td>
  </tr>
</table>

---

### ProductImageSortKeys

<p>The set of valid sort keys for the ProductImage query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Sort by the <code>created_at</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>POSITION</strong></td>
    <td><p>Sort by the <code>position</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
</table>

---

### ProductMediaSortKeys

<p>The set of valid sort keys for the ProductMedia query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>POSITION</strong></td>
    <td><p>Sort by the <code>position</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
</table>

---

### ProductSortKeys

<p>The set of valid sort keys for the Product query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BEST_SELLING</strong></td>
    <td><p>Sort by the <code>best_selling</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Sort by the <code>created_at</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>PRICE</strong></td>
    <td><p>Sort by the <code>price</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>PRODUCT_TYPE</strong></td>
    <td><p>Sort by the <code>product_type</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>TITLE</strong></td>
    <td><p>Sort by the <code>title</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Sort by the <code>updated_at</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>VENDOR</strong></td>
    <td><p>Sort by the <code>vendor</code> value.</p></td>
  </tr>
</table>

---

### ProductVariantSortKeys

<p>The set of valid sort keys for the ProductVariant query.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ID</strong></td>
    <td><p>Sort by the <code>id</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>POSITION</strong></td>
    <td><p>Sort by the <code>position</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>RELEVANCE</strong></td>
    <td><p>During a search (i.e. when the <code>query</code> parameter has been specified on the connection) this sorts the
results by relevance to the search term(s). When no search query is specified, this sort key is not
deterministic and should not be used.</p></td>
  </tr>
  <tr>
    <td><strong>SKU</strong></td>
    <td><p>Sort by the <code>sku</code> value.</p></td>
  </tr>
  <tr>
    <td><strong>TITLE</strong></td>
    <td><p>Sort by the <code>title</code> value.</p></td>
  </tr>
</table>

---

### TransactionKind

<p>The different kinds of order transactions.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AUTHORIZATION</strong></td>
    <td><p>An amount reserved against the cardholder&rsquo;s funding source.
Money does not change hands until the authorization is captured.</p></td>
  </tr>
  <tr>
    <td><strong>CAPTURE</strong></td>
    <td><p>A transfer of the money that was reserved during the authorization stage.</p></td>
  </tr>
  <tr>
    <td><strong>CHANGE</strong></td>
    <td><p>Money returned to the customer when they have paid too much.</p></td>
  </tr>
  <tr>
    <td><strong>EMV_AUTHORIZATION</strong></td>
    <td><p>An authorization for a payment taken with an EMV credit card reader.</p></td>
  </tr>
  <tr>
    <td><strong>SALE</strong></td>
    <td><p>An authorization and capture performed together in a single step.</p></td>
  </tr>
</table>

---

### TransactionStatus

<p>Transaction statuses describe the status of a transaction.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ERROR</strong></td>
    <td><p>There was an error while processing the transaction.</p></td>
  </tr>
  <tr>
    <td><strong>FAILURE</strong></td>
    <td><p>The transaction failed.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>The transaction is pending.</p></td>
  </tr>
  <tr>
    <td><strong>SUCCESS</strong></td>
    <td><p>The transaction succeeded.</p></td>
  </tr>
</table>

---

### UnitPriceMeasurementMeasuredType

<p>The accepted types of unit of measurement.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AREA</strong></td>
    <td><p>Unit of measurements representing areas.</p></td>
  </tr>
  <tr>
    <td><strong>LENGTH</strong></td>
    <td><p>Unit of measurements representing lengths.</p></td>
  </tr>
  <tr>
    <td><strong>VOLUME</strong></td>
    <td><p>Unit of measurements representing volumes.</p></td>
  </tr>
  <tr>
    <td><strong>WEIGHT</strong></td>
    <td><p>Unit of measurements representing weights.</p></td>
  </tr>
</table>

---

### UnitPriceMeasurementMeasuredUnit

<p>The valid units of measurement for a unit price measurement.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CL</strong></td>
    <td><p>100 centiliters equals 1 liter.</p></td>
  </tr>
  <tr>
    <td><strong>CM</strong></td>
    <td><p>100 centimeters equals 1 meter.</p></td>
  </tr>
  <tr>
    <td><strong>G</strong></td>
    <td><p>Metric system unit of weight.</p></td>
  </tr>
  <tr>
    <td><strong>KG</strong></td>
    <td><p>1 kilogram equals 1000 grams.</p></td>
  </tr>
  <tr>
    <td><strong>L</strong></td>
    <td><p>Metric system unit of volume.</p></td>
  </tr>
  <tr>
    <td><strong>M</strong></td>
    <td><p>Metric system unit of length.</p></td>
  </tr>
  <tr>
    <td><strong>M2</strong></td>
    <td><p>Metric system unit of area.</p></td>
  </tr>
  <tr>
    <td><strong>M3</strong></td>
    <td><p>1 cubic meter equals 1000 liters.</p></td>
  </tr>
  <tr>
    <td><strong>MG</strong></td>
    <td><p>1000 milligrams equals 1 gram.</p></td>
  </tr>
  <tr>
    <td><strong>ML</strong></td>
    <td><p>1000 milliliters equals 1 liter.</p></td>
  </tr>
  <tr>
    <td><strong>MM</strong></td>
    <td><p>1000 millimeters equals 1 meter.</p></td>
  </tr>
</table>

---

### WeightUnit

<p>Units of measurement for weight.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>GRAMS</strong></td>
    <td><p>Metric system unit of mass.</p></td>
  </tr>
  <tr>
    <td><strong>KILOGRAMS</strong></td>
    <td><p>1 kilogram equals 1000 grams.</p></td>
  </tr>
  <tr>
    <td><strong>OUNCES</strong></td>
    <td><p>Imperial system unit of mass.</p></td>
  </tr>
  <tr>
    <td><strong>POUNDS</strong></td>
    <td><p>1 pound equals 16 ounces.</p></td>
  </tr>
</table>

---