# Objects

### About objects

[Objects](https://graphql.github.io/graphql-spec/June2018/#sec-Objects) in GraphQL represent the resources you can access. An object can contain a list of fields, which are specifically typed.

### ApiVersion

<p>A version of the API.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>displayName</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The human-readable name of the version.</p></td>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The unique identifier of an ApiVersion. All supported API versions have a date-based (YYYY-MM) or <code>unstable</code> handle.</p></td>
  </tr>
  <tr>
    <td><strong>supported</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether the version is supported by Shopify.</p></td>
  </tr>
</table>

---

### AppliedGiftCard

<p>Details about the gift card used on the checkout.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>amountUsed</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The amount that was taken from the gift card by applying it.</p></td>
  </tr>
  <tr>
    <td><strong>amountUsedV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The amount that was taken from the gift card by applying it.</p></td>
  </tr>
  <tr>
    <td><strong>balance</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The amount left on the gift card.</p></td>
  </tr>
  <tr>
    <td><strong>balanceV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The amount left on the gift card.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>lastCharacters</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The last characters of the gift card.</p></td>
  </tr>
  <tr>
    <td><strong>presentmentAmountUsed</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The amount that was applied to the checkout in its currency.</p></td>
  </tr>
</table>

---

### Article

<p>An article in an online store blog.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>author</strong> (<a href="objects.md#articleauthor">ArticleAuthor!</a>)</td> 
    <td><p>The article&rsquo;s author.</p></td>
  </tr>
  <tr>
    <td><strong>authorV2</strong> (<a href="objects.md#articleauthor">ArticleAuthor</a>)</td> 
    <td><p>The article&rsquo;s author.</p></td>
  </tr>
  <tr>
    <td><strong>blog</strong> (<a href="objects.md#blog">Blog!</a>)</td> 
    <td><p>The blog that the article belongs to.</p></td>
  </tr>
  <tr>
    <td><strong>comments</strong> (<a href="objects.md#commentconnection">CommentConnection!</a>)</td> 
    <td>
      <p><p>List of comments posted on the article.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>content</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td>
      <p><p>Stripped content of the article, single line with HTML tags removed.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>truncateAt (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncates string after the given length.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>contentHtml</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td><p>The content of the article, complete with HTML formatting.</p></td>
  </tr>
  <tr>
    <td><strong>excerpt</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td>
      <p><p>Stripped excerpt of the article, single line with HTML tags removed.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>truncateAt (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncates string after the given length.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>excerptHtml</strong> (<a href="scalars.md#html">HTML</a>)</td> 
    <td><p>The excerpt of the article, complete with HTML formatting.</p></td>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A human-friendly unique string for the Article automatically generated from its title.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>image</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td>
      <p><p>The image associated with the article.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>maxWidth (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image width in pixels between 1 and 2048. This argument is deprecated: Use <code>maxWidth</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>maxHeight (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image height in pixels between 1 and 2048. This argument is deprecated: Use
<code>maxHeight</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>crop (<a href="enums.md#cropregion">CropRegion</a>)</p>
            <p><p>Crops the image according to the specified region. This argument is
deprecated: Use <code>crop</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>scale (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image size multiplier for high-resolution retina displays. Must be between 1
and 3. This argument is deprecated: Use <code>scale</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>publishedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the article was published.</p></td>
  </tr>
  <tr>
    <td><strong>seo</strong> (<a href="objects.md#seo">SEO</a>)</td> 
    <td><p>The article’s SEO information.</p></td>
  </tr>
  <tr>
    <td><strong>tags</strong> (<a href="scalars.md#string">[String!]!</a>)</td> 
    <td><p>A categorization that a article can be tagged with.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The article’s name.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The url pointing to the article accessible from the web.</p></td>
  </tr>
</table>

---

### ArticleAuthor

<p>The author of an article.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>bio</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The author&rsquo;s bio.</p></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The author’s email.</p></td>
  </tr>
  <tr>
    <td><strong>firstName</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The author&rsquo;s first name.</p></td>
  </tr>
  <tr>
    <td><strong>lastName</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The author&rsquo;s last name.</p></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The author&rsquo;s full name.</p></td>
  </tr>
</table>

---

### ArticleConnection

<p>An auto-generated type for paginating through multiple Articles.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#articleedge">[ArticleEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### ArticleEdge

<p>An auto-generated type which holds one Article and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#article">Article!</a>)</td> 
    <td><p>The item at the end of ArticleEdge.</p></td>
  </tr>
</table>

---

### Attribute

<p>Represents a generic custom attribute.</p>  

#### Fields

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
    <td><strong>value</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>Value of the attribute.</p></td>
  </tr>
</table>

---

### AutomaticDiscountApplication

<p>Automatic discount applications capture the intentions of a discount that was automatically applied.</p> 

#### Implements


- [DiscountApplication](interfaces.md#discountapplication) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allocationMethod</strong> (<a href="enums.md#discountapplicationallocationmethod">DiscountApplicationAllocationMethod!</a>)</td> 
    <td><p>The method by which the discount&rsquo;s value is allocated to its entitled items.</p></td>
  </tr>
  <tr>
    <td><strong>targetSelection</strong> (<a href="enums.md#discountapplicationtargetselection">DiscountApplicationTargetSelection!</a>)</td> 
    <td><p>Which lines of targetType that the discount is allocated over.</p></td>
  </tr>
  <tr>
    <td><strong>targetType</strong> (<a href="enums.md#discountapplicationtargettype">DiscountApplicationTargetType!</a>)</td> 
    <td><p>The type of line that the discount is applicable towards.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The title of the application.</p></td>
  </tr>
  <tr>
    <td><strong>value</strong> (<a href="unions.md#pricingvalue">PricingValue!</a>)</td> 
    <td><p>The value of the discount application.</p></td>
  </tr>
</table>

---

### AvailableShippingRates

<p>A collection of available shipping rates for a checkout.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ready</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether or not the shipping rates are ready.
The <code>shippingRates</code> field is <code>null</code> when this value is <code>false</code>.
This field should be polled until its value becomes <code>true</code>.</p></td>
  </tr>
  <tr>
    <td><strong>shippingRates</strong> (<a href="objects.md#shippingrate">[ShippingRate!]</a>)</td> 
    <td><p>The fetched shipping rates. <code>null</code> until the <code>ready</code> field is <code>true</code>.</p></td>
  </tr>
</table>

---

### Blog

<p>An online store blog.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>articleByHandle</strong> (<a href="objects.md#article">Article</a>)</td> 
    <td>
      <p><p>Find an article by its handle.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>handle (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The handle of the article.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>articles</strong> (<a href="objects.md#articleconnection">ArticleConnection!</a>)</td> 
    <td>
      <p><p>List of the blog&rsquo;s articles.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#articlesortkeys">ArticleSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>author</code>
 - <code>blog_title</code>
 - <code>created_at</code>
 - <code>tag</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>authors</strong> (<a href="objects.md#articleauthor">[ArticleAuthor!]!</a>)</td> 
    <td><p>The authors who have contributed to the blog.</p></td>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A human-friendly unique string for the Blog automatically generated from its title.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>seo</strong> (<a href="objects.md#seo">SEO</a>)</td> 
    <td><p>The blog&rsquo;s SEO information.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The blogs’s title.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The url pointing to the blog accessible from the web.</p></td>
  </tr>
</table>

---

### BlogConnection

<p>An auto-generated type for paginating through multiple Blogs.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#blogedge">[BlogEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### BlogEdge

<p>An auto-generated type which holds one Blog and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#blog">Blog!</a>)</td> 
    <td><p>The item at the end of BlogEdge.</p></td>
  </tr>
</table>

---

### Checkout

<p>A container for all the information required to checkout items and pay.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>appliedGiftCards</strong> (<a href="objects.md#appliedgiftcard">[AppliedGiftCard!]!</a>)</td> 
    <td><p>The gift cards used on the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>availableShippingRates</strong> (<a href="objects.md#availableshippingrates">AvailableShippingRates</a>)</td> 
    <td><p>The available shipping rates for this Checkout.
Should only be used when checkout <code>requiresShipping</code> is <code>true</code> and
the shipping address is valid.</p></td>
  </tr>
  <tr>
    <td><strong>completedAt</strong> (<a href="scalars.md#datetime">DateTime</a>)</td> 
    <td><p>The date and time when the checkout was completed.</p></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the checkout was created.</p></td>
  </tr>
  <tr>
    <td><strong>currencyCode</strong> (<a href="enums.md#currencycode">CurrencyCode!</a>)</td> 
    <td><p>The currency code for the Checkout.</p></td>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="objects.md#attribute">[Attribute!]!</a>)</td> 
    <td><p>A list of extra information that is added to the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The customer associated with the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>discountApplications</strong> (<a href="objects.md#discountapplicationconnection">DiscountApplicationConnection!</a>)</td> 
    <td>
      <p><p>Discounts that have been applied on the checkout.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The email attached to this checkout.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>lineItems</strong> (<a href="objects.md#checkoutlineitemconnection">CheckoutLineItemConnection!</a>)</td> 
    <td>
      <p><p>A list of line item objects, each one containing information about an item in the checkout.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>lineItemsSubtotalPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The sum of all the prices of all the items in the checkout. Duties, taxes, shipping and discounts excluded.</p></td>
  </tr>
  <tr>
    <td><strong>note</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The note associated with the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>order</strong> (<a href="objects.md#order">Order</a>)</td> 
    <td><p>The resulting order from a paid checkout.</p></td>
  </tr>
  <tr>
    <td><strong>orderStatusUrl</strong> (<a href="scalars.md#url">URL</a>)</td> 
    <td><p>The Order Status Page for this Checkout, null when checkout is not completed.</p></td>
  </tr>
  <tr>
    <td><strong>paymentDue</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The amount left to be paid. This is equal to the cost of the line items, taxes
and shipping minus discounts and gift cards.</p></td>
  </tr>
  <tr>
    <td><strong>paymentDueV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The amount left to be paid. This is equal to the cost of the line items,
duties, taxes and shipping minus discounts and gift cards.</p></td>
  </tr>
  <tr>
    <td><strong>ready</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether or not the Checkout is ready and can be completed. Checkouts may
have asynchronous operations that can take time to finish. If you want
to complete a checkout or ensure all the fields are populated and up to
date, polling is required until the value is true.</p></td>
  </tr>
  <tr>
    <td><strong>requiresShipping</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>States whether or not the fulfillment requires shipping.</p></td>
  </tr>
  <tr>
    <td><strong>shippingAddress</strong> (<a href="objects.md#mailingaddress">MailingAddress</a>)</td> 
    <td><p>The shipping address to where the line items will be shipped.</p></td>
  </tr>
  <tr>
    <td><strong>shippingDiscountAllocations</strong> (<a href="objects.md#discountallocation">[DiscountAllocation!]!</a>)</td> 
    <td><p>The discounts that have been allocated onto the shipping line by discount applications.</p></td>
  </tr>
  <tr>
    <td><strong>shippingLine</strong> (<a href="objects.md#shippingrate">ShippingRate</a>)</td> 
    <td><p>Once a shipping rate is selected by the customer it is transitioned to a <code>shipping_line</code> object.</p></td>
  </tr>
  <tr>
    <td><strong>subtotalPrice</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>Price of the checkout before shipping and taxes.</p></td>
  </tr>
  <tr>
    <td><strong>subtotalPriceV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>Price of the checkout before duties, shipping and taxes.</p></td>
  </tr>
  <tr>
    <td><strong>taxExempt</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Specifies if the Checkout is tax exempt.</p></td>
  </tr>
  <tr>
    <td><strong>taxesIncluded</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Specifies if taxes are included in the line item and shipping line prices.</p></td>
  </tr>
  <tr>
    <td><strong>totalDuties</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>The sum of all the duties applied to the line items in the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>totalPrice</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The sum of all the prices of all the items in the checkout, taxes and discounts included.</p></td>
  </tr>
  <tr>
    <td><strong>totalPriceV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The sum of all the prices of all the items in the checkout, duties, taxes and discounts included.</p></td>
  </tr>
  <tr>
    <td><strong>totalTax</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The sum of all the taxes applied to the line items and shipping lines in the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>totalTaxV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The sum of all the taxes applied to the line items and shipping lines in the checkout.</p></td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the checkout was last updated.</p></td>
  </tr>
  <tr>
    <td><strong>webUrl</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The url pointing to the checkout accessible from the web.</p></td>
  </tr>
</table>

---

### CheckoutAttributesUpdatePayload

<p>Return type for <code>checkoutAttributesUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutAttributesUpdateV2Payload

<p>Return type for <code>checkoutAttributesUpdateV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCompleteFreePayload

<p>Return type for <code>checkoutCompleteFree</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCompleteWithCreditCardPayload

<p>Return type for <code>checkoutCompleteWithCreditCard</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The checkout on which the payment was applied.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>payment</strong> (<a href="objects.md#payment">Payment</a>)</td> 
    <td><p>A representation of the attempted payment.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCompleteWithCreditCardV2Payload

<p>Return type for <code>checkoutCompleteWithCreditCardV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The checkout on which the payment was applied.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>payment</strong> (<a href="objects.md#payment">Payment</a>)</td> 
    <td><p>A representation of the attempted payment.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCompleteWithTokenizedPaymentPayload

<p>Return type for <code>checkoutCompleteWithTokenizedPayment</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The checkout on which the payment was applied.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>payment</strong> (<a href="objects.md#payment">Payment</a>)</td> 
    <td><p>A representation of the attempted payment.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCompleteWithTokenizedPaymentV2Payload

<p>Return type for <code>checkoutCompleteWithTokenizedPaymentV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The checkout on which the payment was applied.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>payment</strong> (<a href="objects.md#payment">Payment</a>)</td> 
    <td><p>A representation of the attempted payment.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCompleteWithTokenizedPaymentV3Payload

<p>Return type for <code>checkoutCompleteWithTokenizedPaymentV3</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The checkout on which the payment was applied.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>payment</strong> (<a href="objects.md#payment">Payment</a>)</td> 
    <td><p>A representation of the attempted payment.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCreatePayload

<p>Return type for <code>checkoutCreate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The new checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCustomerAssociatePayload

<p>Return type for <code>checkoutCustomerAssociate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The associated customer object.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCustomerAssociateV2Payload

<p>Return type for <code>checkoutCustomerAssociateV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The associated customer object.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCustomerDisassociatePayload

<p>Return type for <code>checkoutCustomerDisassociate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutCustomerDisassociateV2Payload

<p>Return type for <code>checkoutCustomerDisassociateV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutDiscountCodeApplyPayload

<p>Return type for <code>checkoutDiscountCodeApply</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutDiscountCodeApplyV2Payload

<p>Return type for <code>checkoutDiscountCodeApplyV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutDiscountCodeRemovePayload

<p>Return type for <code>checkoutDiscountCodeRemove</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutEmailUpdatePayload

<p>Return type for <code>checkoutEmailUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The checkout object with the updated email.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutEmailUpdateV2Payload

<p>Return type for <code>checkoutEmailUpdateV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The checkout object with the updated email.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutGiftCardApplyPayload

<p>Return type for <code>checkoutGiftCardApply</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutGiftCardRemovePayload

<p>Return type for <code>checkoutGiftCardRemove</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutGiftCardRemoveV2Payload

<p>Return type for <code>checkoutGiftCardRemoveV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutGiftCardsAppendPayload

<p>Return type for <code>checkoutGiftCardsAppend</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutLineItem

<p>A single line item in the checkout, grouped by variant and attributes.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="objects.md#attribute">[Attribute!]!</a>)</td> 
    <td><p>Extra information in the form of an array of Key-Value pairs about the line item.</p></td>
  </tr>
  <tr>
    <td><strong>discountAllocations</strong> (<a href="objects.md#discountallocation">[DiscountAllocation!]!</a>)</td> 
    <td><p>The discounts that have been allocated onto the checkout line item by discount applications.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>quantity</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The quantity of the line item.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Title of the line item. Defaults to the product&rsquo;s title.</p></td>
  </tr>
  <tr>
    <td><strong>unitPrice</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>Unit price of the line item.</p></td>
  </tr>
  <tr>
    <td><strong>variant</strong> (<a href="objects.md#productvariant">ProductVariant</a>)</td> 
    <td><p>Product variant of the line item.</p></td>
  </tr>
</table>

---

### CheckoutLineItemConnection

<p>An auto-generated type for paginating through multiple CheckoutLineItems.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#checkoutlineitemedge">[CheckoutLineItemEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### CheckoutLineItemEdge

<p>An auto-generated type which holds one CheckoutLineItem and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#checkoutlineitem">CheckoutLineItem!</a>)</td> 
    <td><p>The item at the end of CheckoutLineItemEdge.</p></td>
  </tr>
</table>

---

### CheckoutLineItemsAddPayload

<p>Return type for <code>checkoutLineItemsAdd</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutLineItemsRemovePayload

<p>Return type for <code>checkoutLineItemsRemove</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutLineItemsReplacePayload

<p>Return type for <code>checkoutLineItemsReplace</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutLineItemsUpdatePayload

<p>Return type for <code>checkoutLineItemsUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutShippingAddressUpdatePayload

<p>Return type for <code>checkoutShippingAddressUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutShippingAddressUpdateV2Payload

<p>Return type for <code>checkoutShippingAddressUpdateV2</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutShippingLineUpdatePayload

<p>Return type for <code>checkoutShippingLineUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The updated checkout object.</p></td>
  </tr>
  <tr>
    <td><strong>checkoutUserErrors</strong> (<a href="objects.md#checkoutusererror">[CheckoutUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CheckoutUserError

<p>Represents an error that happens during execution of a checkout mutation.</p> 

#### Implements


- [DisplayableError](interfaces.md#displayableerror) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>code</strong> (<a href="enums.md#checkouterrorcode">CheckoutErrorCode</a>)</td> 
    <td><p>Error code to uniquely identify the error.</p></td>
  </tr>
  <tr>
    <td><strong>field</strong> (<a href="scalars.md#string">[String!]</a>)</td> 
    <td><p>Path to the input field which caused the error.</p></td>
  </tr>
  <tr>
    <td><strong>message</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The error message.</p></td>
  </tr>
</table>

---

### Collection

<p>A collection represents a grouping of products that a shop owner can create to
organize them or make their shops easier to browse.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td>
      <p><p>Stripped description of the collection, single line with HTML tags removed.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>truncateAt (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncates string after the given length.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>descriptionHtml</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td><p>The description of the collection, complete with HTML formatting.</p></td>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A human-friendly unique string for the collection automatically generated from its title.
Limit of 255 characters.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>image</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td>
      <p><p>Image associated with the collection.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>maxWidth (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image width in pixels between 1 and 2048. This argument is deprecated: Use <code>maxWidth</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>maxHeight (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image height in pixels between 1 and 2048. This argument is deprecated: Use
<code>maxHeight</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>crop (<a href="enums.md#cropregion">CropRegion</a>)</p>
            <p><p>Crops the image according to the specified region. This argument is
deprecated: Use <code>crop</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>scale (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image size multiplier for high-resolution retina displays. Must be between 1
and 3. This argument is deprecated: Use <code>scale</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>products</strong> (<a href="objects.md#productconnection">ProductConnection!</a>)</td> 
    <td>
      <p><p>List of products in the collection.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#productcollectionsortkeys">ProductCollectionSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The collection’s name. Limit of 255 characters.</p></td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the collection was last modified.</p></td>
  </tr>
</table>

---

### CollectionConnection

<p>An auto-generated type for paginating through multiple Collections.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#collectionedge">[CollectionEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### CollectionEdge

<p>An auto-generated type which holds one Collection and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#collection">Collection!</a>)</td> 
    <td><p>The item at the end of CollectionEdge.</p></td>
  </tr>
</table>

---

### Comment

<p>A comment on an article.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>author</strong> (<a href="objects.md#commentauthor">CommentAuthor!</a>)</td> 
    <td><p>The comment’s author.</p></td>
  </tr>
  <tr>
    <td><strong>content</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td>
      <p><p>Stripped content of the comment, single line with HTML tags removed.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>truncateAt (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncates string after the given length.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>contentHtml</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td><p>The content of the comment, complete with HTML formatting.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
</table>

---

### CommentAuthor

<p>The author of a comment.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The author&rsquo;s email.</p></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The author’s name.</p></td>
  </tr>
</table>

---

### CommentConnection

<p>An auto-generated type for paginating through multiple Comments.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#commentedge">[CommentEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### CommentEdge

<p>An auto-generated type which holds one Comment and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#comment">Comment!</a>)</td> 
    <td><p>The item at the end of CommentEdge.</p></td>
  </tr>
</table>

---

### CreditCard

<p>Credit card information used for a payment.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>brand</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The brand of the credit card.</p></td>
  </tr>
  <tr>
    <td><strong>expiryMonth</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The expiry month of the credit card.</p></td>
  </tr>
  <tr>
    <td><strong>expiryYear</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The expiry year of the credit card.</p></td>
  </tr>
  <tr>
    <td><strong>firstDigits</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The credit card&rsquo;s BIN number.</p></td>
  </tr>
  <tr>
    <td><strong>firstName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The first name of the card holder.</p></td>
  </tr>
  <tr>
    <td><strong>lastDigits</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The last 4 digits of the credit card.</p></td>
  </tr>
  <tr>
    <td><strong>lastName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The last name of the card holder.</p></td>
  </tr>
  <tr>
    <td><strong>maskedNumber</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The masked credit card number with only the last 4 digits displayed.</p></td>
  </tr>
</table>

---

### Customer

<p>A customer represents a customer account with the shop. Customer accounts store
contact information for the customer, saving logged-in customers the trouble of
having to provide it at every checkout.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>acceptsMarketing</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates whether the customer has consented to be sent marketing material via email.</p></td>
  </tr>
  <tr>
    <td><strong>addresses</strong> (<a href="objects.md#mailingaddressconnection">MailingAddressConnection!</a>)</td> 
    <td>
      <p><p>A list of addresses for the customer.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the customer was created.</p></td>
  </tr>
  <tr>
    <td><strong>defaultAddress</strong> (<a href="objects.md#mailingaddress">MailingAddress</a>)</td> 
    <td><p>The customer’s default address.</p></td>
  </tr>
  <tr>
    <td><strong>displayName</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The customer’s name, email or phone number.</p></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The customer’s email address.</p></td>
  </tr>
  <tr>
    <td><strong>firstName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The customer’s first name.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>A unique identifier for the customer.</p></td>
  </tr>
  <tr>
    <td><strong>lastIncompleteCheckout</strong> (<a href="objects.md#checkout">Checkout</a>)</td> 
    <td><p>The customer&rsquo;s most recently updated, incomplete checkout.</p></td>
  </tr>
  <tr>
    <td><strong>lastName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The customer’s last name.</p></td>
  </tr>
  <tr>
    <td><strong>orders</strong> (<a href="objects.md#orderconnection">OrderConnection!</a>)</td> 
    <td>
      <p><p>The orders associated with the customer.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#ordersortkeys">OrderSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>processed_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>phone</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The customer’s phone number.</p></td>
  </tr>
  <tr>
    <td><strong>tags</strong> (<a href="scalars.md#string">[String!]!</a>)</td> 
    <td><p>A comma separated list of tags that have been added to the customer.
Additional access scope required: unauthenticated_read_customer_tags.</p></td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the customer information was updated.</p></td>
  </tr>
</table>

---

### CustomerAccessToken

<p>A CustomerAccessToken represents the unique token required to make modifications to the customer object.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>accessToken</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The customer’s access token.</p></td>
  </tr>
  <tr>
    <td><strong>expiresAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the customer access token expires.</p></td>
  </tr>
</table>

---

### CustomerAccessTokenCreatePayload

<p>Return type for <code>customerAccessTokenCreate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>The newly created customer access token object.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerAccessTokenCreateWithMultipassPayload

<p>Return type for <code>customerAccessTokenCreateWithMultipass</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>An access token object associated with the customer.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerAccessTokenDeletePayload

<p>Return type for <code>customerAccessTokenDelete</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>deletedAccessToken</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The destroyed access token.</p></td>
  </tr>
  <tr>
    <td><strong>deletedCustomerAccessTokenId</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>ID of the destroyed customer access token.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerAccessTokenRenewPayload

<p>Return type for <code>customerAccessTokenRenew</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>The renewed customer access token object.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerActivateByUrlPayload

<p>Return type for <code>customerActivateByUrl</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The customer that was activated.</p></td>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>A new customer access token for the customer.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerActivatePayload

<p>Return type for <code>customerActivate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The customer object.</p></td>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>A newly created customer access token object for the customer.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerAddressCreatePayload

<p>Return type for <code>customerAddressCreate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customerAddress</strong> (<a href="objects.md#mailingaddress">MailingAddress</a>)</td> 
    <td><p>The new customer address object.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerAddressDeletePayload

<p>Return type for <code>customerAddressDelete</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>deletedCustomerAddressId</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>ID of the deleted customer address.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerAddressUpdatePayload

<p>Return type for <code>customerAddressUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customerAddress</strong> (<a href="objects.md#mailingaddress">MailingAddress</a>)</td> 
    <td><p>The customer’s updated mailing address.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerCreatePayload

<p>Return type for <code>customerCreate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The created customer object.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerDefaultAddressUpdatePayload

<p>Return type for <code>customerDefaultAddressUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The updated customer object.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerRecoverPayload

<p>Return type for <code>customerRecover</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerResetByUrlPayload

<p>Return type for <code>customerResetByUrl</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The customer object which was reset.</p></td>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>A newly created customer access token object for the customer.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerResetPayload

<p>Return type for <code>customerReset</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The customer object which was reset.</p></td>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>A newly created customer access token object for the customer.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerUpdatePayload

<p>Return type for <code>customerUpdate</code> mutation.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td><p>The updated customer object.</p></td>
  </tr>
  <tr>
    <td><strong>customerAccessToken</strong> (<a href="objects.md#customeraccesstoken">CustomerAccessToken</a>)</td> 
    <td><p>The newly created customer access token. If the customer&rsquo;s password is updated, all previous access tokens
(including the one used to perform this mutation) become invalid, and a new token is generated.</p></td>
  </tr>
  <tr>
    <td><strong>customerUserErrors</strong> (<a href="objects.md#customerusererror">[CustomerUserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
  <tr>
    <td><strong>userErrors</strong> (<a href="objects.md#usererror">[UserError!]!</a>)</td> 
    <td><p>List of errors that occurred executing the mutation.</p></td>
  </tr>
</table>

---

### CustomerUserError

<p>Represents an error that happens during execution of a customer mutation.</p> 

#### Implements


- [DisplayableError](interfaces.md#displayableerror) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>code</strong> (<a href="enums.md#customererrorcode">CustomerErrorCode</a>)</td> 
    <td><p>Error code to uniquely identify the error.</p></td>
  </tr>
  <tr>
    <td><strong>field</strong> (<a href="scalars.md#string">[String!]</a>)</td> 
    <td><p>Path to the input field which caused the error.</p></td>
  </tr>
  <tr>
    <td><strong>message</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The error message.</p></td>
  </tr>
</table>

---

### DiscountAllocation

<p>An amount discounting the line that has been allocated by a discount.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allocatedAmount</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>Amount of discount allocated.</p></td>
  </tr>
  <tr>
    <td><strong>discountApplication</strong> (<a href="interfaces.md#discountapplication">DiscountApplication!</a>)</td> 
    <td><p>The discount this allocated amount originated from.</p></td>
  </tr>
</table>

---

### DiscountApplicationConnection

<p>An auto-generated type for paginating through multiple DiscountApplications.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#discountapplicationedge">[DiscountApplicationEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### DiscountApplicationEdge

<p>An auto-generated type which holds one DiscountApplication and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="interfaces.md#discountapplication">DiscountApplication!</a>)</td> 
    <td><p>The item at the end of DiscountApplicationEdge.</p></td>
  </tr>
</table>

---

### DiscountCodeApplication

<p>Discount code applications capture the intentions of a discount code at
the time that it is applied.</p> 

#### Implements


- [DiscountApplication](interfaces.md#discountapplication) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allocationMethod</strong> (<a href="enums.md#discountapplicationallocationmethod">DiscountApplicationAllocationMethod!</a>)</td> 
    <td><p>The method by which the discount&rsquo;s value is allocated to its entitled items.</p></td>
  </tr>
  <tr>
    <td><strong>applicable</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Specifies whether the discount code was applied successfully.</p></td>
  </tr>
  <tr>
    <td><strong>code</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The string identifying the discount code that was used at the time of application.</p></td>
  </tr>
  <tr>
    <td><strong>targetSelection</strong> (<a href="enums.md#discountapplicationtargetselection">DiscountApplicationTargetSelection!</a>)</td> 
    <td><p>Which lines of targetType that the discount is allocated over.</p></td>
  </tr>
  <tr>
    <td><strong>targetType</strong> (<a href="enums.md#discountapplicationtargettype">DiscountApplicationTargetType!</a>)</td> 
    <td><p>The type of line that the discount is applicable towards.</p></td>
  </tr>
  <tr>
    <td><strong>value</strong> (<a href="unions.md#pricingvalue">PricingValue!</a>)</td> 
    <td><p>The value of the discount application.</p></td>
  </tr>
</table>

---

### Domain

<p>Represents a web address.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>host</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The host name of the domain (eg: <code>example.com</code>).</p></td>
  </tr>
  <tr>
    <td><strong>sslEnabled</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether SSL is enabled or not.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The URL of the domain (eg: <code>https://example.com</code>).</p></td>
  </tr>
</table>

---

### ExternalVideo

<p>Represents a video hosted outside of Shopify.</p> 

#### Implements


- [Media](interfaces.md#media)
- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>alt</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A word or phrase to share the nature or contents of a media.</p></td>
  </tr>
  <tr>
    <td><strong>embeddedUrl</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The URL.</p></td>
  </tr>
  <tr>
    <td><strong>host</strong> (<a href="enums.md#mediahost">MediaHost!</a>)</td> 
    <td><p>The host of the external video.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>mediaContentType</strong> (<a href="enums.md#mediacontenttype">MediaContentType!</a>)</td> 
    <td><p>The media content type.</p></td>
  </tr>
  <tr>
    <td><strong>previewImage</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td><p>The preview image for the media.</p></td>
  </tr>
</table>

---

### Fulfillment

<p>Represents a single fulfillment in an order.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>fulfillmentLineItems</strong> (<a href="objects.md#fulfillmentlineitemconnection">FulfillmentLineItemConnection!</a>)</td> 
    <td>
      <p><p>List of the fulfillment&rsquo;s line items.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>trackingCompany</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of the tracking company.</p></td>
  </tr>
  <tr>
    <td><strong>trackingInfo</strong> (<a href="objects.md#fulfillmenttrackinginfo">[FulfillmentTrackingInfo!]!</a>)</td> 
    <td>
      <p><p>Tracking information associated with the fulfillment,
such as the tracking number and tracking URL.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncate the array result to this size.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### FulfillmentLineItem

<p>Represents a single line item in a fulfillment. There is at most one fulfillment line item for each order line item.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>lineItem</strong> (<a href="objects.md#orderlineitem">OrderLineItem!</a>)</td> 
    <td><p>The associated order&rsquo;s line item.</p></td>
  </tr>
  <tr>
    <td><strong>quantity</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The amount fulfilled in this fulfillment.</p></td>
  </tr>
</table>

---

### FulfillmentLineItemConnection

<p>An auto-generated type for paginating through multiple FulfillmentLineItems.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#fulfillmentlineitemedge">[FulfillmentLineItemEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### FulfillmentLineItemEdge

<p>An auto-generated type which holds one FulfillmentLineItem and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#fulfillmentlineitem">FulfillmentLineItem!</a>)</td> 
    <td><p>The item at the end of FulfillmentLineItemEdge.</p></td>
  </tr>
</table>

---

### FulfillmentTrackingInfo

<p>Tracking information associated with the fulfillment.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>number</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The tracking number of the fulfillment.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#url">URL</a>)</td> 
    <td><p>The URL to track the fulfillment.</p></td>
  </tr>
</table>

---

### Image

<p>Represents an image resource.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>altText</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A word or phrase to share the nature or contents of an image.</p></td>
  </tr>
  <tr>
    <td><strong>height</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The original height of the image in pixels. Returns <code>null</code> if the image is not hosted by Shopify.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID</a>)</td> 
    <td><p>A unique identifier for the image.</p></td>
  </tr>
  <tr>
    <td><strong>originalSrc</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The location of the original image as a URL.</p>

<p>If there are any existing transformations in the original source URL, they will remain and not be stripped.</p></td>
  </tr>
  <tr>
    <td><strong>src</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The location of the image as a URL.</p></td>
  </tr>
  <tr>
    <td><strong>transformedSrc</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td>
      <p><p>The location of the transformed image as a URL.</p>

<p>All transformation arguments are considered &ldquo;best-effort&rdquo;. If they can be applied to an image, they will be.
Otherwise any transformations which an image type does not support will be ignored.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>maxWidth (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image width in pixels between 1 and 5760.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>maxHeight (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image height in pixels between 1 and 5760.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>crop (<a href="enums.md#cropregion">CropRegion</a>)</p>
            <p><p>Crops the image according to the specified region.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>scale (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image size multiplier for high-resolution retina displays. Must be between 1 and 3.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>preferredContentType (<a href="enums.md#imagecontenttype">ImageContentType</a>)</p>
            <p><p>Best effort conversion of image into content type (SVG -&gt; PNG, Anything -&gt; JGP, Anything -&gt; WEBP are supported).</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>width</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The original width of the image in pixels. Returns <code>null</code> if the image is not hosted by Shopify.</p></td>
  </tr>
</table>

---

### ImageConnection

<p>An auto-generated type for paginating through multiple Images.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#imageedge">[ImageEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### ImageEdge

<p>An auto-generated type which holds one Image and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#image">Image!</a>)</td> 
    <td><p>The item at the end of ImageEdge.</p></td>
  </tr>
</table>

---

### MailingAddress

<p>Represents a mailing address for customers and shipping.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

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
    <td><strong>countryCode</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The two-letter code for the country of the address.</p>

<p>For example, US.</p></td>
  </tr>
  <tr>
    <td><strong>countryCodeV2</strong> (<a href="enums.md#countrycode">CountryCode</a>)</td> 
    <td><p>The two-letter code for the country of the address.</p>

<p>For example, US.</p></td>
  </tr>
  <tr>
    <td><strong>firstName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The first name of the customer.</p></td>
  </tr>
  <tr>
    <td><strong>formatted</strong> (<a href="scalars.md#string">[String!]!</a>)</td> 
    <td>
      <p><p>A formatted version of the address, customized by the provided arguments.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>withName (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Whether to include the customer&rsquo;s name in the formatted address.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>withCompany (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Whether to include the customer&rsquo;s company in the formatted address.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>formattedArea</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A comma-separated list of the values for city, province, and country.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>lastName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The last name of the customer.</p></td>
  </tr>
  <tr>
    <td><strong>latitude</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The latitude coordinate of the customer address.</p></td>
  </tr>
  <tr>
    <td><strong>longitude</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The longitude coordinate of the customer address.</p></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The full name of the customer, based on firstName and lastName.</p></td>
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
    <td><strong>provinceCode</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The two-letter code for the region.</p>

<p>For example, ON.</p></td>
  </tr>
  <tr>
    <td><strong>zip</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The zip or postal code of the address.</p></td>
  </tr>
</table>

---

### MailingAddressConnection

<p>An auto-generated type for paginating through multiple MailingAddresses.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#mailingaddressedge">[MailingAddressEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### MailingAddressEdge

<p>An auto-generated type which holds one MailingAddress and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#mailingaddress">MailingAddress!</a>)</td> 
    <td><p>The item at the end of MailingAddressEdge.</p></td>
  </tr>
</table>

---

### ManualDiscountApplication

<p>Manual discount applications capture the intentions of a discount that was manually created.</p> 

#### Implements


- [DiscountApplication](interfaces.md#discountapplication) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allocationMethod</strong> (<a href="enums.md#discountapplicationallocationmethod">DiscountApplicationAllocationMethod!</a>)</td> 
    <td><p>The method by which the discount&rsquo;s value is allocated to its entitled items.</p></td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The description of the application.</p></td>
  </tr>
  <tr>
    <td><strong>targetSelection</strong> (<a href="enums.md#discountapplicationtargetselection">DiscountApplicationTargetSelection!</a>)</td> 
    <td><p>Which lines of targetType that the discount is allocated over.</p></td>
  </tr>
  <tr>
    <td><strong>targetType</strong> (<a href="enums.md#discountapplicationtargettype">DiscountApplicationTargetType!</a>)</td> 
    <td><p>The type of line that the discount is applicable towards.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The title of the application.</p></td>
  </tr>
  <tr>
    <td><strong>value</strong> (<a href="unions.md#pricingvalue">PricingValue!</a>)</td> 
    <td><p>The value of the discount application.</p></td>
  </tr>
</table>

---

### MediaConnection

<p>An auto-generated type for paginating through multiple Media.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#mediaedge">[MediaEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### MediaEdge

<p>An auto-generated type which holds one Media and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="interfaces.md#media">Media!</a>)</td> 
    <td><p>The item at the end of MediaEdge.</p></td>
  </tr>
</table>

---

### MediaImage

<p>Represents a Shopify hosted image.</p> 

#### Implements


- [Media](interfaces.md#media)
- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>alt</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A word or phrase to share the nature or contents of a media.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>image</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td><p>The image for the media.</p></td>
  </tr>
  <tr>
    <td><strong>mediaContentType</strong> (<a href="enums.md#mediacontenttype">MediaContentType!</a>)</td> 
    <td><p>The media content type.</p></td>
  </tr>
  <tr>
    <td><strong>previewImage</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td><p>The preview image for the media.</p></td>
  </tr>
</table>

---

### Metafield

<p>Metafields represent custom metadata attached to a resource. Metafields can be sorted into namespaces and are
comprised of keys, values, and value types.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the storefront metafield was created.</p></td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The description of a metafield.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>key</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The key name for a metafield.</p></td>
  </tr>
  <tr>
    <td><strong>namespace</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The namespace for a metafield.</p></td>
  </tr>
  <tr>
    <td><strong>parentResource</strong> (<a href="unions.md#metafieldparentresource">MetafieldParentResource!</a>)</td> 
    <td><p>The parent object that the metafield belongs to.</p></td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the storefront metafield was updated.</p></td>
  </tr>
  <tr>
    <td><strong>value</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The value of a metafield.</p></td>
  </tr>
  <tr>
    <td><strong>valueType</strong> (<a href="enums.md#metafieldvaluetype">MetafieldValueType!</a>)</td> 
    <td><p>Represents the metafield value type.</p></td>
  </tr>
</table>

---

### MetafieldConnection

<p>An auto-generated type for paginating through multiple Metafields.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#metafieldedge">[MetafieldEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### MetafieldEdge

<p>An auto-generated type which holds one Metafield and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#metafield">Metafield!</a>)</td> 
    <td><p>The item at the end of MetafieldEdge.</p></td>
  </tr>
</table>

---

### Model3d

<p>Represents a Shopify hosted 3D model.</p> 

#### Implements


- [Media](interfaces.md#media)
- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>alt</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A word or phrase to share the nature or contents of a media.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>mediaContentType</strong> (<a href="enums.md#mediacontenttype">MediaContentType!</a>)</td> 
    <td><p>The media content type.</p></td>
  </tr>
  <tr>
    <td><strong>previewImage</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td><p>The preview image for the media.</p></td>
  </tr>
  <tr>
    <td><strong>sources</strong> (<a href="objects.md#model3dsource">[Model3dSource!]!</a>)</td> 
    <td><p>The sources for a 3d model.</p></td>
  </tr>
</table>

---

### Model3dSource

<p>Represents a source for a Shopify hosted 3d model.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>filesize</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The filesize of the 3d model.</p></td>
  </tr>
  <tr>
    <td><strong>format</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The format of the 3d model.</p></td>
  </tr>
  <tr>
    <td><strong>mimeType</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The MIME type of the 3d model.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The URL of the 3d model.</p></td>
  </tr>
</table>

---

### MoneyV2

<p>A monetary value with currency.</p>

<p>To format currencies, combine this type&rsquo;s amount and currencyCode fields with your client&rsquo;s locale.</p>

<p>For example, in JavaScript you could use Intl.NumberFormat:</p>

<pre><code class="language-js">new Intl.NumberFormat(locale, {
  style: 'currency',
  currency: currencyCode
}).format(amount);
</code></pre>

<p>Other formatting libraries include:</p>

<ul>
<li>iOS - <a href="https://developer.apple.com/documentation/foundation/numberformatter">NumberFormatter</a></li>
<li>Android - <a href="https://developer.android.com/reference/java/text/NumberFormat.html">NumberFormat</a></li>
<li>PHP - <a href="http://php.net/manual/en/class.numberformatter.php">NumberFormatter</a></li>
</ul>

<p>For a more general solution, the [Unicode CLDR number formatting database] is available with many implementations
(such as <a href="https://github.com/twitter/twitter-cldr-rb">TwitterCldr</a>).</p>  

#### Fields

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

### MoneyV2Connection

<p>An auto-generated type for paginating through multiple MoneyV2s.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#moneyv2edge">[MoneyV2Edge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### MoneyV2Edge

<p>An auto-generated type which holds one MoneyV2 and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The item at the end of MoneyV2Edge.</p></td>
  </tr>
</table>

---

### Order

<p>An order is a customer’s completed request to purchase one or more products from
a shop. An order is created when a customer completes the checkout process,
during which time they provides an email address, billing address and payment information.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cancelReason</strong> (<a href="enums.md#ordercancelreason">OrderCancelReason</a>)</td> 
    <td><p>The reason for the order&rsquo;s cancellation. Returns <code>null</code> if the order wasn&rsquo;t canceled.</p></td>
  </tr>
  <tr>
    <td><strong>canceledAt</strong> (<a href="scalars.md#datetime">DateTime</a>)</td> 
    <td><p>The date and time when the order was canceled. Returns null if the order wasn&rsquo;t canceled.</p></td>
  </tr>
  <tr>
    <td><strong>currencyCode</strong> (<a href="enums.md#currencycode">CurrencyCode!</a>)</td> 
    <td><p>The code of the currency used for the payment.</p></td>
  </tr>
  <tr>
    <td><strong>currentSubtotalPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The subtotal of line items and their discounts, excluding line items that have
been removed. Does not contain order-level discounts, duties, shipping costs,
or shipping discounts. Taxes are not included unless the order is a
taxes-included order.</p></td>
  </tr>
  <tr>
    <td><strong>currentTotalDuties</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>The total cost of duties for the order, including refunds.</p></td>
  </tr>
  <tr>
    <td><strong>currentTotalPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The total amount of the order, including duties, taxes and discounts, minus amounts for line items that have been removed.</p></td>
  </tr>
  <tr>
    <td><strong>currentTotalTax</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The total of all taxes applied to the order, excluding taxes for returned line items.</p></td>
  </tr>
  <tr>
    <td><strong>customerLocale</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The locale code in which this specific order happened.</p></td>
  </tr>
  <tr>
    <td><strong>customerUrl</strong> (<a href="scalars.md#url">URL</a>)</td> 
    <td><p>The unique URL that the customer can use to access the order.</p></td>
  </tr>
  <tr>
    <td><strong>discountApplications</strong> (<a href="objects.md#discountapplicationconnection">DiscountApplicationConnection!</a>)</td> 
    <td>
      <p><p>Discounts that have been applied on the order.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>edited</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether the order has had any edits applied or not.</p></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The customer&rsquo;s email address.</p></td>
  </tr>
  <tr>
    <td><strong>financialStatus</strong> (<a href="enums.md#orderfinancialstatus">OrderFinancialStatus</a>)</td> 
    <td><p>The financial status of the order.</p></td>
  </tr>
  <tr>
    <td><strong>fulfillmentStatus</strong> (<a href="enums.md#orderfulfillmentstatus">OrderFulfillmentStatus!</a>)</td> 
    <td><p>The fulfillment status for the order.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>lineItems</strong> (<a href="objects.md#orderlineitemconnection">OrderLineItemConnection!</a>)</td> 
    <td>
      <p><p>List of the order’s line items.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Unique identifier for the order that appears on the order.
For example, <em>#1000</em> or _Store1001.</p></td>
  </tr>
  <tr>
    <td><strong>orderNumber</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>A unique numeric identifier for the order for use by shop owner and customer.</p></td>
  </tr>
  <tr>
    <td><strong>originalTotalDuties</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>The total cost of duties charged at checkout.</p></td>
  </tr>
  <tr>
    <td><strong>originalTotalPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The total price of the order before any applied edits.</p></td>
  </tr>
  <tr>
    <td><strong>phone</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The customer&rsquo;s phone number for receiving SMS notifications.</p></td>
  </tr>
  <tr>
    <td><strong>processedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the order was imported.
This value can be set to dates in the past when importing from other systems.
If no value is provided, it will be auto-generated based on current date and time.</p></td>
  </tr>
  <tr>
    <td><strong>shippingAddress</strong> (<a href="objects.md#mailingaddress">MailingAddress</a>)</td> 
    <td><p>The address to where the order will be shipped.</p></td>
  </tr>
  <tr>
    <td><strong>shippingDiscountAllocations</strong> (<a href="objects.md#discountallocation">[DiscountAllocation!]!</a>)</td> 
    <td><p>The discounts that have been allocated onto the shipping line by discount applications.</p></td>
  </tr>
  <tr>
    <td><strong>statusUrl</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The unique URL for the order&rsquo;s status page.</p></td>
  </tr>
  <tr>
    <td><strong>subtotalPrice</strong> (<a href="scalars.md#money">Money</a>)</td> 
    <td><p>Price of the order before shipping and taxes.</p></td>
  </tr>
  <tr>
    <td><strong>subtotalPriceV2</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>Price of the order before duties, shipping and taxes.</p></td>
  </tr>
  <tr>
    <td><strong>successfulFulfillments</strong> (<a href="objects.md#fulfillment">[Fulfillment!]</a>)</td> 
    <td>
      <p><p>List of the order’s successful fulfillments.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncate the array result to this size.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>totalPrice</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The sum of all the prices of all the items in the order, taxes and discounts included (must be positive).</p></td>
  </tr>
  <tr>
    <td><strong>totalPriceV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The sum of all the prices of all the items in the order, duties, taxes and discounts included (must be positive).</p></td>
  </tr>
  <tr>
    <td><strong>totalRefunded</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The total amount that has been refunded.</p></td>
  </tr>
  <tr>
    <td><strong>totalRefundedV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The total amount that has been refunded.</p></td>
  </tr>
  <tr>
    <td><strong>totalShippingPrice</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The total cost of shipping.</p></td>
  </tr>
  <tr>
    <td><strong>totalShippingPriceV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The total cost of shipping.</p></td>
  </tr>
  <tr>
    <td><strong>totalTax</strong> (<a href="scalars.md#money">Money</a>)</td> 
    <td><p>The total cost of taxes.</p></td>
  </tr>
  <tr>
    <td><strong>totalTaxV2</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>The total cost of taxes.</p></td>
  </tr>
</table>

---

### OrderConnection

<p>An auto-generated type for paginating through multiple Orders.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#orderedge">[OrderEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### OrderEdge

<p>An auto-generated type which holds one Order and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#order">Order!</a>)</td> 
    <td><p>The item at the end of OrderEdge.</p></td>
  </tr>
</table>

---

### OrderLineItem

<p>Represents a single line in an order. There is one line item for each distinct product variant.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>currentQuantity</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The number of entries associated to the line item minus the items that have been removed.</p></td>
  </tr>
  <tr>
    <td><strong>customAttributes</strong> (<a href="objects.md#attribute">[Attribute!]!</a>)</td> 
    <td><p>List of custom attributes associated to the line item.</p></td>
  </tr>
  <tr>
    <td><strong>discountAllocations</strong> (<a href="objects.md#discountallocation">[DiscountAllocation!]!</a>)</td> 
    <td><p>The discounts that have been allocated onto the order line item by discount applications.</p></td>
  </tr>
  <tr>
    <td><strong>discountedTotalPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The total price of the line item, including discounts, and displayed in the presentment currency.</p></td>
  </tr>
  <tr>
    <td><strong>originalTotalPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The total price of the line item, not including any discounts. The total price
is calculated using the original unit price multiplied by the quantity, and it
is displayed in the presentment currency.</p></td>
  </tr>
  <tr>
    <td><strong>quantity</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The number of products variants associated to the line item.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The title of the product combined with title of the variant.</p></td>
  </tr>
  <tr>
    <td><strong>variant</strong> (<a href="objects.md#productvariant">ProductVariant</a>)</td> 
    <td><p>The product variant object associated to the line item.</p></td>
  </tr>
</table>

---

### OrderLineItemConnection

<p>An auto-generated type for paginating through multiple OrderLineItems.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#orderlineitemedge">[OrderLineItemEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### OrderLineItemEdge

<p>An auto-generated type which holds one OrderLineItem and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#orderlineitem">OrderLineItem!</a>)</td> 
    <td><p>The item at the end of OrderLineItemEdge.</p></td>
  </tr>
</table>

---

### Page

<p>Shopify merchants can create pages to hold static HTML content. Each Page object
represents a custom page on the online store.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>body</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td><p>The description of the page, complete with HTML formatting.</p></td>
  </tr>
  <tr>
    <td><strong>bodySummary</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Summary of the page body.</p></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The timestamp of the page creation.</p></td>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A human-friendly unique string for the page automatically generated from its title.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>seo</strong> (<a href="objects.md#seo">SEO</a>)</td> 
    <td><p>The page&rsquo;s SEO information.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The title of the page.</p></td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The timestamp of the latest page update.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The url pointing to the page accessible from the web.</p></td>
  </tr>
</table>

---

### PageConnection

<p>An auto-generated type for paginating through multiple Pages.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#pageedge">[PageEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### PageEdge

<p>An auto-generated type which holds one Page and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#page">Page!</a>)</td> 
    <td><p>The item at the end of PageEdge.</p></td>
  </tr>
</table>

---

### PageInfo

<p>Information about pagination in a connection.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>hasNextPage</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if there are more pages to fetch.</p></td>
  </tr>
  <tr>
    <td><strong>hasPreviousPage</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if there are any pages prior to the current page.</p></td>
  </tr>
</table>

---

### Payment

<p>A payment applied to a checkout.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

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
    <td><strong>amountV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The amount of the payment.</p></td>
  </tr>
  <tr>
    <td><strong>billingAddress</strong> (<a href="objects.md#mailingaddress">MailingAddress</a>)</td> 
    <td><p>The billing address for the payment.</p></td>
  </tr>
  <tr>
    <td><strong>checkout</strong> (<a href="objects.md#checkout">Checkout!</a>)</td> 
    <td><p>The checkout to which the payment belongs.</p></td>
  </tr>
  <tr>
    <td><strong>creditCard</strong> (<a href="objects.md#creditcard">CreditCard</a>)</td> 
    <td><p>The credit card used for the payment in the case of direct payments.</p></td>
  </tr>
  <tr>
    <td><strong>errorMessage</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A message describing a processing error during asynchronous processing.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>idempotencyKey</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A client-side generated token to identify a payment and perform idempotent operations.
For more information, refer to
<a href="https://shopify.dev/concepts/about-apis/idempotent-requests">Idempotent requests</a>.</p></td>
  </tr>
  <tr>
    <td><strong>nextActionUrl</strong> (<a href="scalars.md#url">URL</a>)</td> 
    <td><p>The URL where the customer needs to be redirected so they can complete the 3D Secure payment flow.</p></td>
  </tr>
  <tr>
    <td><strong>ready</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether or not the payment is still processing asynchronously.</p></td>
  </tr>
  <tr>
    <td><strong>test</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>A flag to indicate if the payment is to be done in test mode for gateways that support it.</p></td>
  </tr>
  <tr>
    <td><strong>transaction</strong> (<a href="objects.md#transaction">Transaction</a>)</td> 
    <td><p>The actual transaction recorded by Shopify after having processed the payment with the gateway.</p></td>
  </tr>
</table>

---

### PaymentSettings

<p>Settings related to payments.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>acceptedCardBrands</strong> (<a href="enums.md#cardbrand">[CardBrand!]!</a>)</td> 
    <td><p>List of the card brands which the shop accepts.</p></td>
  </tr>
  <tr>
    <td><strong>cardVaultUrl</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>The url pointing to the endpoint to vault credit cards.</p></td>
  </tr>
  <tr>
    <td><strong>countryCode</strong> (<a href="enums.md#countrycode">CountryCode!</a>)</td> 
    <td><p>The country where the shop is located.</p></td>
  </tr>
  <tr>
    <td><strong>currencyCode</strong> (<a href="enums.md#currencycode">CurrencyCode!</a>)</td> 
    <td><p>The three-letter code for the shop&rsquo;s primary currency.</p></td>
  </tr>
  <tr>
    <td><strong>enabledPresentmentCurrencies</strong> (<a href="enums.md#currencycode">[CurrencyCode!]!</a>)</td> 
    <td><p>A list of enabled currencies (ISO 4217 format) that the shop accepts.
Merchants can enable currencies from their Shopify Payments settings in the Shopify admin.</p></td>
  </tr>
  <tr>
    <td><strong>shopifyPaymentsAccountId</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The shop’s Shopify Payments account id.</p></td>
  </tr>
  <tr>
    <td><strong>supportedDigitalWallets</strong> (<a href="enums.md#digitalwallet">[DigitalWallet!]!</a>)</td> 
    <td><p>List of the digital wallets which the shop supports.</p></td>
  </tr>
</table>

---

### PricingPercentageValue

<p>The value of the percentage pricing object.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>percentage</strong> (<a href="scalars.md#float">Float!</a>)</td> 
    <td><p>The percentage value of the object.</p></td>
  </tr>
</table>

---

### Product

<p>A product represents an individual item for sale in a Shopify store. Products are often physical, but they don&rsquo;t have to be.
For example, a digital download (such as a movie, music or ebook file) also
qualifies as a product, as do services (such as equipment rental, work for hire,
customization of another product or an extended warranty).</p> 

#### Implements


- [HasMetafields](interfaces.md#hasmetafields)
- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>availableForSale</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if at least one product variant is available for sale.</p></td>
  </tr>
  <tr>
    <td><strong>collections</strong> (<a href="objects.md#collectionconnection">CollectionConnection!</a>)</td> 
    <td>
      <p><p>List of collections a product belongs to.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>compareAtPriceRange</strong> (<a href="objects.md#productpricerange">ProductPriceRange!</a>)</td> 
    <td><p>The compare at price of the product across all variants.</p></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the product was created.</p></td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td>
      <p><p>Stripped description of the product, single line with HTML tags removed.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>truncateAt (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncates string after the given length.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>descriptionHtml</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td><p>The description of the product, complete with HTML formatting.</p></td>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A human-friendly unique string for the Product automatically generated from its title.
They are used by the Liquid templating language to refer to objects.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>images</strong> (<a href="objects.md#imageconnection">ImageConnection!</a>)</td> 
    <td>
      <p><p>List of images associated with the product.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#productimagesortkeys">ProductImageSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>maxWidth (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image width in pixels between 1 and 2048. This argument is deprecated: Use <code>maxWidth</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>maxHeight (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image height in pixels between 1 and 2048. This argument is deprecated: Use
<code>maxHeight</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>crop (<a href="enums.md#cropregion">CropRegion</a>)</p>
            <p><p>Crops the image according to the specified region. This argument is
deprecated: Use <code>crop</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>scale (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image size multiplier for high-resolution retina displays. Must be between 1
and 3. This argument is deprecated: Use <code>scale</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>media</strong> (<a href="objects.md#mediaconnection">MediaConnection!</a>)</td> 
    <td>
      <p><p>The media associated with the product.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#productmediasortkeys">ProductMediaSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>metafield</strong> (<a href="objects.md#metafield">Metafield</a>)</td> 
    <td>
      <p><p>The metafield associated with the resource.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>namespace (<a href="scalars.md#string">String!</a>)</p>
            <p><p>Container for a set of metafields (maximum of 20 characters).</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>key (<a href="scalars.md#string">String!</a>)</p>
            <p><p>Identifier for the metafield (maximum of 30 characters).</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>metafields</strong> (<a href="objects.md#metafieldconnection">MetafieldConnection!</a>)</td> 
    <td>
      <p><p>A paginated list of metafields associated with the resource.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>namespace (<a href="scalars.md#string">String</a>)</p>
            <p><p>Container for a set of metafields (maximum of 20 characters).</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>onlineStoreUrl</strong> (<a href="scalars.md#url">URL</a>)</td> 
    <td><p>The online store URL for the product.
A value of <code>null</code> indicates that the product is not published to the Online Store sales channel.</p></td>
  </tr>
  <tr>
    <td><strong>options</strong> (<a href="objects.md#productoption">[ProductOption!]!</a>)</td> 
    <td>
      <p><p>List of product options.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Truncate the array result to this size.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>presentmentPriceRanges</strong> (<a href="objects.md#productpricerangeconnection">ProductPriceRangeConnection!</a>)</td> 
    <td>
      <p><p>List of price ranges in the presentment currencies for this shop.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>presentmentCurrencies (<a href="enums.md#currencycode">[CurrencyCode!]</a>)</p>
            <p><p>Specifies the presentment currencies to return a price range in.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>priceRange</strong> (<a href="objects.md#productpricerange">ProductPriceRange!</a>)</td> 
    <td><p>The price range.</p></td>
  </tr>
  <tr>
    <td><strong>productType</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A categorization that a product can be tagged with, commonly used for filtering and searching.</p></td>
  </tr>
  <tr>
    <td><strong>publishedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the product was published to the channel.</p></td>
  </tr>
  <tr>
    <td><strong>seo</strong> (<a href="objects.md#seo">SEO!</a>)</td> 
    <td><p>The product&rsquo;s SEO information.</p></td>
  </tr>
  <tr>
    <td><strong>tags</strong> (<a href="scalars.md#string">[String!]!</a>)</td> 
    <td><p>A comma separated list of tags that have been added to the product.
Additional access scope required for private apps: unauthenticated_read_product_tags.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The product’s title.</p></td>
  </tr>
  <tr>
    <td><strong>totalInventory</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The total quantity of inventory in stock for this Product.</p></td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>The date and time when the product was last modified.
A product&rsquo;s <code>updatedAt</code> value can change for different reasons. For example, if an order
is placed for a product that has inventory tracking set up, then the inventory adjustment
is counted as an update.</p></td>
  </tr>
  <tr>
    <td><strong>variantBySelectedOptions</strong> (<a href="objects.md#productvariant">ProductVariant</a>)</td> 
    <td>
      <p><p>Find a product’s variant based on its selected options.
This is useful for converting a user’s selection of product options into a single matching variant.
If there is not a variant for the selected options, <code>null</code> will be returned.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>selectedOptions (<a href="input_objects.md#selectedoptioninput">[SelectedOptionInput!]!</a>)</p>
            <p><p>The input fields used for a selected option.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>variants</strong> (<a href="objects.md#productvariantconnection">ProductVariantConnection!</a>)</td> 
    <td>
      <p><p>List of the product’s variants.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#productvariantsortkeys">ProductVariantSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>vendor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The product’s vendor name.</p></td>
  </tr>
</table>

---

### ProductConnection

<p>An auto-generated type for paginating through multiple Products.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#productedge">[ProductEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### ProductEdge

<p>An auto-generated type which holds one Product and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#product">Product!</a>)</td> 
    <td><p>The item at the end of ProductEdge.</p></td>
  </tr>
</table>

---

### ProductOption

<p>Product property names like &ldquo;Size&rdquo;, &ldquo;Color&rdquo;, and &ldquo;Material&rdquo; that the customers can select.
Variants are selected based on permutations of these options.
255 characters limit each.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The product option’s name.</p></td>
  </tr>
  <tr>
    <td><strong>values</strong> (<a href="scalars.md#string">[String!]!</a>)</td> 
    <td><p>The corresponding value to the product option name.</p></td>
  </tr>
</table>

---

### ProductPriceRange

<p>The price range of the product.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>maxVariantPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The highest variant&rsquo;s price.</p></td>
  </tr>
  <tr>
    <td><strong>minVariantPrice</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The lowest variant&rsquo;s price.</p></td>
  </tr>
</table>

---

### ProductPriceRangeConnection

<p>An auto-generated type for paginating through multiple ProductPriceRanges.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#productpricerangeedge">[ProductPriceRangeEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### ProductPriceRangeEdge

<p>An auto-generated type which holds one ProductPriceRange and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#productpricerange">ProductPriceRange!</a>)</td> 
    <td><p>The item at the end of ProductPriceRangeEdge.</p></td>
  </tr>
</table>

---

### ProductVariant

<p>A product variant represents a different version of a product, such as differing sizes or differing colors.</p> 

#### Implements


- [HasMetafields](interfaces.md#hasmetafields)
- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>available</strong> (<a href="scalars.md#boolean">Boolean</a>)</td> 
    <td><p>Indicates if the product variant is in stock.</p></td>
  </tr>
  <tr>
    <td><strong>availableForSale</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if the product variant is available for sale.</p></td>
  </tr>
  <tr>
    <td><strong>compareAtPrice</strong> (<a href="scalars.md#money">Money</a>)</td> 
    <td><p>The compare at price of the variant. This can be used to mark a variant as on
sale, when <code>compareAtPrice</code> is higher than <code>price</code>.</p></td>
  </tr>
  <tr>
    <td><strong>compareAtPriceV2</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>The compare at price of the variant. This can be used to mark a variant as on
sale, when <code>compareAtPriceV2</code> is higher than <code>priceV2</code>.</p></td>
  </tr>
  <tr>
    <td><strong>currentlyNotInStock</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether a product is out of stock but still available for purchase (used for backorders).</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>image</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td>
      <p><p>Image associated with the product variant. This field falls back to the product image if no image is available.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>maxWidth (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image width in pixels between 1 and 2048. This argument is deprecated: Use <code>maxWidth</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>maxHeight (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image height in pixels between 1 and 2048. This argument is deprecated: Use
<code>maxHeight</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>crop (<a href="enums.md#cropregion">CropRegion</a>)</p>
            <p><p>Crops the image according to the specified region. This argument is
deprecated: Use <code>crop</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>scale (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Image size multiplier for high-resolution retina displays. Must be between 1
and 3. This argument is deprecated: Use <code>scale</code> on <code>Image.transformedSrc</code> instead.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>metafield</strong> (<a href="objects.md#metafield">Metafield</a>)</td> 
    <td>
      <p><p>The metafield associated with the resource.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>namespace (<a href="scalars.md#string">String!</a>)</p>
            <p><p>Container for a set of metafields (maximum of 20 characters).</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>key (<a href="scalars.md#string">String!</a>)</p>
            <p><p>Identifier for the metafield (maximum of 30 characters).</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>metafields</strong> (<a href="objects.md#metafieldconnection">MetafieldConnection!</a>)</td> 
    <td>
      <p><p>A paginated list of metafields associated with the resource.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>namespace (<a href="scalars.md#string">String</a>)</p>
            <p><p>Container for a set of metafields (maximum of 20 characters).</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>presentmentPrices</strong> (<a href="objects.md#productvariantpricepairconnection">ProductVariantPricePairConnection!</a>)</td> 
    <td>
      <p><p>List of prices and compare-at prices in the presentment currencies for this shop.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>presentmentCurrencies (<a href="enums.md#currencycode">[CurrencyCode!]</a>)</p>
            <p><p>The presentment currencies prices should return in.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>presentmentUnitPrices</strong> (<a href="objects.md#moneyv2connection">MoneyV2Connection!</a>)</td> 
    <td>
      <p><p>List of unit prices in the presentment currencies for this shop.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>presentmentCurrencies (<a href="enums.md#currencycode">[CurrencyCode!]</a>)</p>
            <p><p>Specify the currencies in which to return presentment unit prices.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>price</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The product variant’s price.</p></td>
  </tr>
  <tr>
    <td><strong>priceV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The product variant’s price.</p></td>
  </tr>
  <tr>
    <td><strong>product</strong> (<a href="objects.md#product">Product!</a>)</td> 
    <td><p>The product object that the product variant belongs to.</p></td>
  </tr>
  <tr>
    <td><strong>quantityAvailable</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>The total sellable quantity of the variant for online sales channels.</p></td>
  </tr>
  <tr>
    <td><strong>requiresShipping</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether a customer needs to provide a shipping address when placing an order for the product variant.</p></td>
  </tr>
  <tr>
    <td><strong>selectedOptions</strong> (<a href="objects.md#selectedoption">[SelectedOption!]!</a>)</td> 
    <td><p>List of product options applied to the variant.</p></td>
  </tr>
  <tr>
    <td><strong>sku</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The SKU (stock keeping unit) associated with the variant.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The product variant’s title.</p></td>
  </tr>
  <tr>
    <td><strong>unitPrice</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>The unit price value for the variant based on the variant&rsquo;s measurement.</p></td>
  </tr>
  <tr>
    <td><strong>unitPriceMeasurement</strong> (<a href="objects.md#unitpricemeasurement">UnitPriceMeasurement</a>)</td> 
    <td><p>The unit price measurement for the variant.</p></td>
  </tr>
  <tr>
    <td><strong>weight</strong> (<a href="scalars.md#float">Float</a>)</td> 
    <td><p>The weight of the product variant in the unit system specified with <code>weight_unit</code>.</p></td>
  </tr>
  <tr>
    <td><strong>weightUnit</strong> (<a href="enums.md#weightunit">WeightUnit!</a>)</td> 
    <td><p>Unit of measurement for weight.</p></td>
  </tr>
</table>

---

### ProductVariantConnection

<p>An auto-generated type for paginating through multiple ProductVariants.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#productvariantedge">[ProductVariantEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### ProductVariantEdge

<p>An auto-generated type which holds one ProductVariant and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#productvariant">ProductVariant!</a>)</td> 
    <td><p>The item at the end of ProductVariantEdge.</p></td>
  </tr>
</table>

---

### ProductVariantPricePair

<p>The compare-at price and price of a variant sharing a currency.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>compareAtPrice</strong> (<a href="objects.md#moneyv2">MoneyV2</a>)</td> 
    <td><p>The compare-at price of the variant with associated currency.</p></td>
  </tr>
  <tr>
    <td><strong>price</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The price of the variant with associated currency.</p></td>
  </tr>
</table>

---

### ProductVariantPricePairConnection

<p>An auto-generated type for paginating through multiple ProductVariantPricePairs.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#productvariantpricepairedge">[ProductVariantPricePairEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### ProductVariantPricePairEdge

<p>An auto-generated type which holds one ProductVariantPricePair and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#productvariantpricepair">ProductVariantPricePair!</a>)</td> 
    <td><p>The item at the end of ProductVariantPricePairEdge.</p></td>
  </tr>
</table>

---

### QueryRoot

<p>The schema’s entry-point for queries. This acts as the public, top-level API from which all queries must start.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>articles</strong> (<a href="objects.md#articleconnection">ArticleConnection!</a>)</td> 
    <td>
      <p><p>List of the shop&rsquo;s articles.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#articlesortkeys">ArticleSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>author</code>
 - <code>blog_title</code>
 - <code>created_at</code>
 - <code>tag</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>blogByHandle</strong> (<a href="objects.md#blog">Blog</a>)</td> 
    <td>
      <p><p>Find a blog by its handle.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>handle (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The handle of the blog.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>blogs</strong> (<a href="objects.md#blogconnection">BlogConnection!</a>)</td> 
    <td>
      <p><p>List of the shop&rsquo;s blogs.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#blogsortkeys">BlogSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>created_at</code>
 - <code>handle</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>collectionByHandle</strong> (<a href="objects.md#collection">Collection</a>)</td> 
    <td>
      <p><p>Find a collection by its handle.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>handle (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The handle of the collection.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>collections</strong> (<a href="objects.md#collectionconnection">CollectionConnection!</a>)</td> 
    <td>
      <p><p>List of the shop’s collections.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#collectionsortkeys">CollectionSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>collection_type</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>customer</strong> (<a href="objects.md#customer">Customer</a>)</td> 
    <td>
      <p><p>Find a customer by its access token.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>customerAccessToken (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The customer access token.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="interfaces.md#node">Node</a>)</td> 
    <td>
      <p><p>Returns a specific node by ID.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>id (<a href="scalars.md#id">ID!</a>)</p>
            <p><p>The ID of the Node to return.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>nodes</strong> (<a href="interfaces.md#node">[Node]!</a>)</td> 
    <td>
      <p><p>Returns the list of nodes with the given IDs.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>ids (<a href="scalars.md#id">[ID!]!</a>)</p>
            <p><p>The IDs of the Nodes to return.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>pageByHandle</strong> (<a href="objects.md#page">Page</a>)</td> 
    <td>
      <p><p>Find a page by its handle.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>handle (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The handle of the page.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>pages</strong> (<a href="objects.md#pageconnection">PageConnection!</a>)</td> 
    <td>
      <p><p>List of the shop&rsquo;s pages.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#pagesortkeys">PageSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>created_at</code>
 - <code>handle</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>productByHandle</strong> (<a href="objects.md#product">Product</a>)</td> 
    <td>
      <p><p>Find a product by its handle.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>handle (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The handle of the product.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>productRecommendations</strong> (<a href="objects.md#product">[Product!]</a>)</td> 
    <td>
      <p><p>Find recommended products related to a given <code>product_id</code>.
To learn more about how recommendations are generated, see
<a href="https://help.shopify.com/themes/development/recommended-products"><em>Showing product recommendations on product pages</em></a>.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>productId (<a href="scalars.md#id">ID!</a>)</p>
            <p><p>The id of the product.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>productTags</strong> (<a href="objects.md#stringconnection">StringConnection!</a>)</td> 
    <td>
      <p><p>Tags added to products.
Additional access scope required: unauthenticated_read_product_tags.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int!</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>productTypes</strong> (<a href="objects.md#stringconnection">StringConnection!</a>)</td> 
    <td>
      <p><p>List of product types for the shop&rsquo;s products that are published to your app.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int!</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>products</strong> (<a href="objects.md#productconnection">ProductConnection!</a>)</td> 
    <td>
      <p><p>List of the shop’s products.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#productsortkeys">ProductSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>available_for_sale</code>
 - <code>created_at</code>
 - <code>product_type</code>
 - <code>tag</code>
 - <code>title</code>
 - <code>updated_at</code>
 - <code>variants.price</code>
 - <code>vendor</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>publicApiVersions</strong> (<a href="objects.md#apiversion">[ApiVersion!]!</a>)</td> 
    <td><p>The list of public Storefront API versions, including supported, release candidate and unstable versions.</p></td>
  </tr>
  <tr>
    <td><strong>shop</strong> (<a href="objects.md#shop">Shop!</a>)</td> 
    <td><p>The shop associated with the storefront access token.</p></td>
  </tr>
</table>

---

### SEO

<p>SEO information.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The meta description.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The SEO title.</p></td>
  </tr>
</table>

---

### ScriptDiscountApplication

<p>Script discount applications capture the intentions of a discount that
was created by a Shopify Script.</p> 

#### Implements


- [DiscountApplication](interfaces.md#discountapplication) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>allocationMethod</strong> (<a href="enums.md#discountapplicationallocationmethod">DiscountApplicationAllocationMethod!</a>)</td> 
    <td><p>The method by which the discount&rsquo;s value is allocated to its entitled items.</p></td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The description of the application as defined by the Script.</p></td>
  </tr>
  <tr>
    <td><strong>targetSelection</strong> (<a href="enums.md#discountapplicationtargetselection">DiscountApplicationTargetSelection!</a>)</td> 
    <td><p>Which lines of targetType that the discount is allocated over.</p></td>
  </tr>
  <tr>
    <td><strong>targetType</strong> (<a href="enums.md#discountapplicationtargettype">DiscountApplicationTargetType!</a>)</td> 
    <td><p>The type of line that the discount is applicable towards.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The title of the application as defined by the Script.</p></td>
  </tr>
  <tr>
    <td><strong>value</strong> (<a href="unions.md#pricingvalue">PricingValue!</a>)</td> 
    <td><p>The value of the discount application.</p></td>
  </tr>
</table>

---

### SelectedOption

<p>Properties used by customers to select a product variant.
Products can have multiple options, like different sizes or colors.</p>  

#### Fields

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

### ShippingRate

<p>A shipping rate to be applied to a checkout.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Human-readable unique identifier for this shipping rate.</p></td>
  </tr>
  <tr>
    <td><strong>price</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>Price of this shipping rate.</p></td>
  </tr>
  <tr>
    <td><strong>priceV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>Price of this shipping rate.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Title of this shipping rate.</p></td>
  </tr>
</table>

---

### Shop

<p>Shop represents a collection of the general settings and information about the shop.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>articles</strong> (<a href="objects.md#articleconnection">ArticleConnection!</a>)</td> 
    <td>
      <p><p>List of the shop&rsquo; articles.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#articlesortkeys">ArticleSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>author</code>
 - <code>blog_title</code>
 - <code>created_at</code>
 - <code>tag</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>blogs</strong> (<a href="objects.md#blogconnection">BlogConnection!</a>)</td> 
    <td>
      <p><p>List of the shop&rsquo; blogs.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#blogsortkeys">BlogSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>created_at</code>
 - <code>handle</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>collectionByHandle</strong> (<a href="objects.md#collection">Collection</a>)</td> 
    <td>
      <p><p>Find a collection by its handle.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>handle (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The handle of the collection.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>collections</strong> (<a href="objects.md#collectionconnection">CollectionConnection!</a>)</td> 
    <td>
      <p><p>List of the shop’s collections.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#collectionsortkeys">CollectionSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>collection_type</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>currencyCode</strong> (<a href="enums.md#currencycode">CurrencyCode!</a>)</td> 
    <td><p>The three-letter code for the currency that the shop accepts.</p></td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A description of the shop.</p></td>
  </tr>
  <tr>
    <td><strong>moneyFormat</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A string representing the way currency is formatted when the currency isn’t specified.</p></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The shop’s name.</p></td>
  </tr>
  <tr>
    <td><strong>paymentSettings</strong> (<a href="objects.md#paymentsettings">PaymentSettings!</a>)</td> 
    <td><p>Settings related to payments.</p></td>
  </tr>
  <tr>
    <td><strong>primaryDomain</strong> (<a href="objects.md#domain">Domain!</a>)</td> 
    <td><p>The shop’s primary domain.</p></td>
  </tr>
  <tr>
    <td><strong>privacyPolicy</strong> (<a href="objects.md#shoppolicy">ShopPolicy</a>)</td> 
    <td><p>The shop’s privacy policy.</p></td>
  </tr>
  <tr>
    <td><strong>productByHandle</strong> (<a href="objects.md#product">Product</a>)</td> 
    <td>
      <p><p>Find a product by its handle.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>handle (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The handle of the product.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>productTags</strong> (<a href="objects.md#stringconnection">StringConnection!</a>)</td> 
    <td>
      <p><p>A list of tags that have been added to products.
Additional access scope required: unauthenticated_read_product_tags.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int!</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>productTypes</strong> (<a href="objects.md#stringconnection">StringConnection!</a>)</td> 
    <td>
      <p><p>List of the shop’s product types.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int!</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>products</strong> (<a href="objects.md#productconnection">ProductConnection!</a>)</td> 
    <td>
      <p><p>List of the shop’s products.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the first <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns up to the last <code>n</code> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>reverse (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Reverse the order of the underlying list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>sortKey (<a href="enums.md#productsortkeys">ProductSortKeys</a>)</p>
            <p><p>Sort the underlying list by the given key.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>query (<a href="scalars.md#string">String</a>)</p>
            <p><p>Supported filter parameters:
 - <code>available_for_sale</code>
 - <code>created_at</code>
 - <code>product_type</code>
 - <code>tag</code>
 - <code>title</code>
 - <code>updated_at</code>
 - <code>variants.price</code>
 - <code>vendor</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>refundPolicy</strong> (<a href="objects.md#shoppolicy">ShopPolicy</a>)</td> 
    <td><p>The shop’s refund policy.</p></td>
  </tr>
  <tr>
    <td><strong>shippingPolicy</strong> (<a href="objects.md#shoppolicy">ShopPolicy</a>)</td> 
    <td><p>The shop’s shipping policy.</p></td>
  </tr>
  <tr>
    <td><strong>shipsToCountries</strong> (<a href="enums.md#countrycode">[CountryCode!]!</a>)</td> 
    <td><p>Countries that the shop ships to.</p></td>
  </tr>
  <tr>
    <td><strong>shopifyPaymentsAccountId</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The shop’s Shopify Payments account id.</p></td>
  </tr>
  <tr>
    <td><strong>termsOfService</strong> (<a href="objects.md#shoppolicy">ShopPolicy</a>)</td> 
    <td><p>The shop’s terms of service.</p></td>
  </tr>
</table>

---

### ShopPolicy

<p>Policy that a merchant has configured for their store, such as their refund or privacy policy.</p> 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>body</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Policy text, maximum size of 64kb.</p></td>
  </tr>
  <tr>
    <td><strong>handle</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Policy’s handle.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Policy’s title.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#url">URL!</a>)</td> 
    <td><p>Public URL to the policy.</p></td>
  </tr>
</table>

---

### StringConnection

<p>An auto-generated type for paginating through multiple Strings.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#stringedge">[StringEdge!]!</a>)</td> 
    <td><p>A list of edges.</p></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td><p>Information to aid in pagination.</p></td>
  </tr>
</table>

---

### StringEdge

<p>An auto-generated type which holds one String and a cursor during pagination.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A cursor for use in pagination.</p></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The item at the end of StringEdge.</p></td>
  </tr>
</table>

---

### Transaction

<p>An object representing exchange of money for a product or service.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>amount</strong> (<a href="scalars.md#money">Money!</a>)</td> 
    <td><p>The amount of money that the transaction was for.</p></td>
  </tr>
  <tr>
    <td><strong>amountV2</strong> (<a href="objects.md#moneyv2">MoneyV2!</a>)</td> 
    <td><p>The amount of money that the transaction was for.</p></td>
  </tr>
  <tr>
    <td><strong>kind</strong> (<a href="enums.md#transactionkind">TransactionKind!</a>)</td> 
    <td><p>The kind of the transaction.</p></td>
  </tr>
  <tr>
    <td><strong>status</strong> (<a href="enums.md#transactionstatus">TransactionStatus!</a>)</td> 
    <td><p>The status of the transaction.</p></td>
  </tr>
  <tr>
    <td><strong>statusV2</strong> (<a href="enums.md#transactionstatus">TransactionStatus</a>)</td> 
    <td><p>The status of the transaction.</p></td>
  </tr>
  <tr>
    <td><strong>test</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether the transaction was done in test mode or not.</p></td>
  </tr>
</table>

---

### UnitPriceMeasurement

<p>The measurement used to calculate a unit price for a product variant (e.g. $9.99 / 100ml).</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>measuredType</strong> (<a href="enums.md#unitpricemeasurementmeasuredtype">UnitPriceMeasurementMeasuredType</a>)</td> 
    <td><p>The type of unit of measurement for the unit price measurement.</p></td>
  </tr>
  <tr>
    <td><strong>quantityUnit</strong> (<a href="enums.md#unitpricemeasurementmeasuredunit">UnitPriceMeasurementMeasuredUnit</a>)</td> 
    <td><p>The quantity unit for the unit price measurement.</p></td>
  </tr>
  <tr>
    <td><strong>quantityValue</strong> (<a href="scalars.md#float">Float!</a>)</td> 
    <td><p>The quantity value for the unit price measurement.</p></td>
  </tr>
  <tr>
    <td><strong>referenceUnit</strong> (<a href="enums.md#unitpricemeasurementmeasuredunit">UnitPriceMeasurementMeasuredUnit</a>)</td> 
    <td><p>The reference unit for the unit price measurement.</p></td>
  </tr>
  <tr>
    <td><strong>referenceValue</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The reference value for the unit price measurement.</p></td>
  </tr>
</table>

---

### UserError

<p>Represents an error in the input of a mutation.</p> 

#### Implements


- [DisplayableError](interfaces.md#displayableerror) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>field</strong> (<a href="scalars.md#string">[String!]</a>)</td> 
    <td><p>Path to the input field which caused the error.</p></td>
  </tr>
  <tr>
    <td><strong>message</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The error message.</p></td>
  </tr>
</table>

---

### Video

<p>Represents a Shopify hosted video.</p> 

#### Implements


- [Media](interfaces.md#media)
- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>alt</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>A word or phrase to share the nature or contents of a media.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>Globally unique identifier.</p></td>
  </tr>
  <tr>
    <td><strong>mediaContentType</strong> (<a href="enums.md#mediacontenttype">MediaContentType!</a>)</td> 
    <td><p>The media content type.</p></td>
  </tr>
  <tr>
    <td><strong>previewImage</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td><p>The preview image for the media.</p></td>
  </tr>
  <tr>
    <td><strong>sources</strong> (<a href="objects.md#videosource">[VideoSource!]!</a>)</td> 
    <td><p>The sources for a video.</p></td>
  </tr>
</table>

---

### VideoSource

<p>Represents a source for a Shopify hosted video.</p>  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>format</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The format of the video source.</p></td>
  </tr>
  <tr>
    <td><strong>height</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The height of the video.</p></td>
  </tr>
  <tr>
    <td><strong>mimeType</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The video MIME type.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The URL of the video.</p></td>
  </tr>
  <tr>
    <td><strong>width</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>The width of the video.</p></td>
  </tr>
</table>

---