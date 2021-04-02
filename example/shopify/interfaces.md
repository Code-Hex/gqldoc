# Interfaces

### About interfaces

[Interfaces](https://graphql.github.io/graphql-spec/June2018/#sec-Interfaces) serve as parent objects from which other objects can inherit.

### DiscountApplication

<p>Discount applications capture the intentions of a discount source at
the time of application.</p> 

#### Implemented by


- [AutomaticDiscountApplication](objects.md#automaticdiscountapplication)
- [DiscountCodeApplication](objects.md#discountcodeapplication)
- [ManualDiscountApplication](objects.md#manualdiscountapplication)
- [ScriptDiscountApplication](objects.md#scriptdiscountapplication) 

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
    <td><strong>value</strong> (<a href="unions.md#pricingvalue">PricingValue!</a>)</td> 
    <td><p>The value of the discount application.</p></td>
  </tr>
</table>

---

### DisplayableError

<p>Represents an error in the input of a mutation.</p> 

#### Implemented by


- [CheckoutUserError](objects.md#checkoutusererror)
- [CustomerUserError](objects.md#customerusererror)
- [UserError](objects.md#usererror) 

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

### HasMetafields

<p>Represents information about the metafields associated to the specified resource.</p> 

#### Implemented by


- [Product](objects.md#product)
- [ProductVariant](objects.md#productvariant) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
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
</table>

---

### Media

<p>Represents a media interface.</p> 

#### Implemented by


- [ExternalVideo](objects.md#externalvideo)
- [MediaImage](objects.md#mediaimage)
- [Model3d](objects.md#model3d)
- [Video](objects.md#video) 

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
    <td><strong>mediaContentType</strong> (<a href="enums.md#mediacontenttype">MediaContentType!</a>)</td> 
    <td><p>The media content type.</p></td>
  </tr>
  <tr>
    <td><strong>previewImage</strong> (<a href="objects.md#image">Image</a>)</td> 
    <td><p>The preview image for the media.</p></td>
  </tr>
</table>

---

### Node

<p>An object with an ID to support global identification.</p> 

#### Implemented by


- [AppliedGiftCard](objects.md#appliedgiftcard)
- [Article](objects.md#article)
- [Blog](objects.md#blog)
- [Checkout](objects.md#checkout)
- [CheckoutLineItem](objects.md#checkoutlineitem)
- [Collection](objects.md#collection)
- [Comment](objects.md#comment)
- [ExternalVideo](objects.md#externalvideo)
- [MailingAddress](objects.md#mailingaddress)
- [MediaImage](objects.md#mediaimage)
- [Metafield](objects.md#metafield)
- [Model3d](objects.md#model3d)
- [Order](objects.md#order)
- [Page](objects.md#page)
- [Payment](objects.md#payment)
- [Product](objects.md#product)
- [ProductOption](objects.md#productoption)
- [ProductVariant](objects.md#productvariant)
- [ShopPolicy](objects.md#shoppolicy)
- [Video](objects.md#video) 

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
</table>

---