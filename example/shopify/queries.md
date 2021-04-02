# Queries

### About queries

The schema’s entry-point for queries. This acts as the public, top-level API from which all queries must start.

### articles

#### Type: [ArticleConnection!](objects.md#articleconnection)

List of the shop's articles. 

#### Arguments

| Name | Description |
|------|-------------|
| first ([Int](scalars.md#int)) | <p>Returns up to the first <code>n</code> elements from the list.</p> |
| after ([String](scalars.md#string)) | <p>Returns the elements that come after the specified cursor.</p> |
| last ([Int](scalars.md#int)) | <p>Returns up to the last <code>n</code> elements from the list.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements that come before the specified cursor.</p> |
| reverse ([Boolean](scalars.md#boolean)) | <p>Reverse the order of the underlying list.</p> |
| sortKey ([ArticleSortKeys](enums.md#articlesortkeys)) | <p>Sort the underlying list by the given key.</p> |
| query ([String](scalars.md#string)) | <p>Supported filter parameters:
 - <code>author</code>
 - <code>blog_title</code>
 - <code>created_at</code>
 - <code>tag</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p> |

---

### blogByHandle

#### Type: [Blog](objects.md#blog)

Find a blog by its handle. 

#### Arguments

| Name | Description |
|------|-------------|
| handle ([String!](scalars.md#string)) | <p>The handle of the blog.</p> |

---

### blogs

#### Type: [BlogConnection!](objects.md#blogconnection)

List of the shop's blogs. 

#### Arguments

| Name | Description |
|------|-------------|
| first ([Int](scalars.md#int)) | <p>Returns up to the first <code>n</code> elements from the list.</p> |
| after ([String](scalars.md#string)) | <p>Returns the elements that come after the specified cursor.</p> |
| last ([Int](scalars.md#int)) | <p>Returns up to the last <code>n</code> elements from the list.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements that come before the specified cursor.</p> |
| reverse ([Boolean](scalars.md#boolean)) | <p>Reverse the order of the underlying list.</p> |
| sortKey ([BlogSortKeys](enums.md#blogsortkeys)) | <p>Sort the underlying list by the given key.</p> |
| query ([String](scalars.md#string)) | <p>Supported filter parameters:
 - <code>created_at</code>
 - <code>handle</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p> |

---

### collectionByHandle

#### Type: [Collection](objects.md#collection)

Find a collection by its handle. 

#### Arguments

| Name | Description |
|------|-------------|
| handle ([String!](scalars.md#string)) | <p>The handle of the collection.</p> |

---

### collections

#### Type: [CollectionConnection!](objects.md#collectionconnection)

List of the shop’s collections. 

#### Arguments

| Name | Description |
|------|-------------|
| first ([Int](scalars.md#int)) | <p>Returns up to the first <code>n</code> elements from the list.</p> |
| after ([String](scalars.md#string)) | <p>Returns the elements that come after the specified cursor.</p> |
| last ([Int](scalars.md#int)) | <p>Returns up to the last <code>n</code> elements from the list.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements that come before the specified cursor.</p> |
| reverse ([Boolean](scalars.md#boolean)) | <p>Reverse the order of the underlying list.</p> |
| sortKey ([CollectionSortKeys](enums.md#collectionsortkeys)) | <p>Sort the underlying list by the given key.</p> |
| query ([String](scalars.md#string)) | <p>Supported filter parameters:
 - <code>collection_type</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p> |

---

### customer

#### Type: [Customer](objects.md#customer)

Find a customer by its access token. 

#### Arguments

| Name | Description |
|------|-------------|
| customerAccessToken ([String!](scalars.md#string)) | <p>The customer access token.</p> |

---

### node

#### Type: [Node](interfaces.md#node)

Returns a specific node by ID. 

#### Arguments

| Name | Description |
|------|-------------|
| id ([ID!](scalars.md#id)) | <p>The ID of the Node to return.</p> |

---

### nodes

#### Type: [[Node]!](interfaces.md#node)

Returns the list of nodes with the given IDs. 

#### Arguments

| Name | Description |
|------|-------------|
| ids ([[ID!]!](scalars.md#id)) | <p>The IDs of the Nodes to return.</p> |

---

### pageByHandle

#### Type: [Page](objects.md#page)

Find a page by its handle. 

#### Arguments

| Name | Description |
|------|-------------|
| handle ([String!](scalars.md#string)) | <p>The handle of the page.</p> |

---

### pages

#### Type: [PageConnection!](objects.md#pageconnection)

List of the shop's pages. 

#### Arguments

| Name | Description |
|------|-------------|
| first ([Int](scalars.md#int)) | <p>Returns up to the first <code>n</code> elements from the list.</p> |
| after ([String](scalars.md#string)) | <p>Returns the elements that come after the specified cursor.</p> |
| last ([Int](scalars.md#int)) | <p>Returns up to the last <code>n</code> elements from the list.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements that come before the specified cursor.</p> |
| reverse ([Boolean](scalars.md#boolean)) | <p>Reverse the order of the underlying list.</p> |
| sortKey ([PageSortKeys](enums.md#pagesortkeys)) | <p>Sort the underlying list by the given key.</p> |
| query ([String](scalars.md#string)) | <p>Supported filter parameters:
 - <code>created_at</code>
 - <code>handle</code>
 - <code>title</code>
 - <code>updated_at</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p> |

---

### productByHandle

#### Type: [Product](objects.md#product)

Find a product by its handle. 

#### Arguments

| Name | Description |
|------|-------------|
| handle ([String!](scalars.md#string)) | <p>The handle of the product.</p> |

---

### productRecommendations

#### Type: [[Product!]](objects.md#product)

Find recommended products related to a given `product_id`.
To learn more about how recommendations are generated, see
[*Showing product recommendations on product pages*](https://help.shopify.com/themes/development/recommended-products). 

#### Arguments

| Name | Description |
|------|-------------|
| productId ([ID!](scalars.md#id)) | <p>The id of the product.</p> |

---

### productTags

#### Type: [StringConnection!](objects.md#stringconnection)

Tags added to products.
Additional access scope required: unauthenticated_read_product_tags. 

#### Arguments

| Name | Description |
|------|-------------|
| first ([Int!](scalars.md#int)) | <p>Returns up to the first <code>n</code> elements from the list.</p> |

---

### productTypes

#### Type: [StringConnection!](objects.md#stringconnection)

List of product types for the shop's products that are published to your app. 

#### Arguments

| Name | Description |
|------|-------------|
| first ([Int!](scalars.md#int)) | <p>Returns up to the first <code>n</code> elements from the list.</p> |

---

### products

#### Type: [ProductConnection!](objects.md#productconnection)

List of the shop’s products. 

#### Arguments

| Name | Description |
|------|-------------|
| first ([Int](scalars.md#int)) | <p>Returns up to the first <code>n</code> elements from the list.</p> |
| after ([String](scalars.md#string)) | <p>Returns the elements that come after the specified cursor.</p> |
| last ([Int](scalars.md#int)) | <p>Returns up to the last <code>n</code> elements from the list.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements that come before the specified cursor.</p> |
| reverse ([Boolean](scalars.md#boolean)) | <p>Reverse the order of the underlying list.</p> |
| sortKey ([ProductSortKeys](enums.md#productsortkeys)) | <p>Sort the underlying list by the given key.</p> |
| query ([String](scalars.md#string)) | <p>Supported filter parameters:
 - <code>available_for_sale</code>
 - <code>created_at</code>
 - <code>product_type</code>
 - <code>tag</code>
 - <code>title</code>
 - <code>updated_at</code>
 - <code>variants.price</code>
 - <code>vendor</code></p>

<p>See the detailed <a href="https://help.shopify.com/api/getting-started/search-syntax">search syntax</a>
for more information about using filters.</p> |

---

### publicApiVersions

#### Type: [[ApiVersion!]!](objects.md#apiversion)

The list of public Storefront API versions, including supported, release candidate and unstable versions. 

---

### shop

#### Type: [Shop!](objects.md#shop)

The shop associated with the storefront access token. 

---