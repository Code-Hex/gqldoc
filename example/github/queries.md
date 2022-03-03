# Queries

### About queries

The query root of GitHub's GraphQL interface.

### codeOfConduct

#### Type: [CodeOfConduct](objects.md#codeofconduct)

Look up a code of conduct by its key 

#### Arguments

| Name | Description |
|------|-------------|
| key ([String!](scalars.md#string)) | <p>The code of conduct&rsquo;s key</p> |

---

### codesOfConduct

#### Type: [[CodeOfConduct]](objects.md#codeofconduct)

Look up a code of conduct by its key 

---

### enterprise

#### Type: [Enterprise](objects.md#enterprise)

Look up an enterprise by URL slug. 

#### Arguments

| Name | Description |
|------|-------------|
| invitationToken ([String](scalars.md#string)) | <p>The enterprise invitation token.</p> |
| slug ([String!](scalars.md#string)) | <p>The enterprise URL slug.</p> |

---

### enterpriseAdministratorInvitation

#### Type: [EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)

Look up a pending enterprise administrator invitation by invitee, enterprise and role. 

#### Arguments

| Name | Description |
|------|-------------|
| enterpriseSlug ([String!](scalars.md#string)) | <p>The slug of the enterprise the user was invited to join.</p> |
| role ([EnterpriseAdministratorRole!](enums.md#enterpriseadministratorrole)) | <p>The role for the business member invitation.</p> |
| userLogin ([String!](scalars.md#string)) | <p>The login of the user invited to join the business.</p> |

---

### enterpriseAdministratorInvitationByToken

#### Type: [EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)

Look up a pending enterprise administrator invitation by invitation token. 

#### Arguments

| Name | Description |
|------|-------------|
| invitationToken ([String!](scalars.md#string)) | <p>The invitation token sent with the invitation email.</p> |

---

### license

#### Type: [License](objects.md#license)

Look up an open source license by its key 

#### Arguments

| Name | Description |
|------|-------------|
| key ([String!](scalars.md#string)) | <p>The license&rsquo;s downcased SPDX ID</p> |

---

### licenses

#### Type: [[License]!](objects.md#license)

Return a list of known open source licenses 

---

### marketplaceCategories

#### Type: [[MarketplaceCategory!]!](objects.md#marketplacecategory)

Get alphabetically sorted list of Marketplace categories 

#### Arguments

| Name | Description |
|------|-------------|
| excludeEmpty ([Boolean](scalars.md#boolean)) | <p>Exclude categories with no listings.</p> |
| excludeSubcategories ([Boolean](scalars.md#boolean)) | <p>Returns top level categories only, excluding any subcategories.</p> |
| includeCategories ([[String!]](scalars.md#string)) | <p>Return only the specified categories.</p> |

---

### marketplaceCategory

#### Type: [MarketplaceCategory](objects.md#marketplacecategory)

Look up a Marketplace category by its slug. 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](scalars.md#string)) | <p>The URL slug of the category.</p> |
| useTopicAliases ([Boolean](scalars.md#boolean)) | <p>Also check topic aliases for the category slug</p> |

---

### marketplaceListing

#### Type: [MarketplaceListing](objects.md#marketplacelisting)

Look up a single Marketplace listing 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](scalars.md#string)) | <p>Select the listing that matches this slug. It&rsquo;s the short name of the listing used in its URL.</p> |

---

### marketplaceListings

#### Type: [MarketplaceListingConnection!](objects.md#marketplacelistingconnection)

Look up Marketplace listings 

#### Arguments

| Name | Description |
|------|-------------|
| adminId ([ID](scalars.md#id)) | <p>Select listings that can be administered by the specified user.</p> |
| after ([String](scalars.md#string)) | <p>Returns the elements in the list that come after the specified cursor.</p> |
| allStates ([Boolean](scalars.md#boolean)) | <p>Select listings visible to the viewer even if they are not approved. If omitted or
false, only approved listings will be returned.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements in the list that come before the specified cursor.</p> |
| categorySlug ([String](scalars.md#string)) | <p>Select only listings with the given category.</p> |
| first ([Int](scalars.md#int)) | <p>Returns the first <em>n</em> elements from the list.</p> |
| last ([Int](scalars.md#int)) | <p>Returns the last <em>n</em> elements from the list.</p> |
| organizationId ([ID](scalars.md#id)) | <p>Select listings for products owned by the specified organization.</p> |
| primaryCategoryOnly ([Boolean](scalars.md#boolean)) | <p>Select only listings where the primary category matches the given category slug.</p> |
| slugs ([[String]](scalars.md#string)) | <p>Select the listings with these slugs, if they are visible to the viewer.</p> |
| useTopicAliases ([Boolean](scalars.md#boolean)) | <p>Also check topic aliases for the category slug</p> |
| viewerCanAdmin ([Boolean](scalars.md#boolean)) | <p>Select listings to which user has admin access. If omitted, listings visible to the
viewer are returned.</p> |
| withFreeTrialsOnly ([Boolean](scalars.md#boolean)) | <p>Select only listings that offer a free trial.</p> |

---

### meta

#### Type: [GitHubMetadata!](objects.md#githubmetadata)

Return information about the GitHub instance 

---

### node

#### Type: [Node](interfaces.md#node)

Fetches an object given its ID. 

#### Arguments

| Name | Description |
|------|-------------|
| id ([ID!](scalars.md#id)) | <p>ID of the object.</p> |

---

### nodes

#### Type: [[Node]!](interfaces.md#node)

Lookup nodes by a list of IDs. 

#### Arguments

| Name | Description |
|------|-------------|
| ids ([[ID!]!](scalars.md#id)) | <p>The list of node IDs.</p> |

---

### organization

#### Type: [Organization](objects.md#organization)

Lookup a organization by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](scalars.md#string)) | <p>The organization&rsquo;s login.</p> |

---

### rateLimit

#### Type: [RateLimit](objects.md#ratelimit)

The client's rate limit information. 

#### Arguments

| Name | Description |
|------|-------------|
| dryRun ([Boolean](scalars.md#boolean)) | <p>If true, calculate the cost for the query without evaluating it</p> |

---

### relay

#### Type: [Query!](objects.md#query)

Hack to workaround https://github.com/facebook/relay/issues/112 re-exposing the root query object 

---

### repository

#### Type: [Repository](objects.md#repository)

Lookup a given repository by the owner and repository name. 

#### Arguments

| Name | Description |
|------|-------------|
| followRenames ([Boolean](scalars.md#boolean)) | <p>Follow repository renames. If disabled, a repository referenced by its old name will return an error.</p> |
| name ([String!](scalars.md#string)) | <p>The name of the repository</p> |
| owner ([String!](scalars.md#string)) | <p>The login field of a user or organization</p> |

---

### repositoryOwner

#### Type: [RepositoryOwner](interfaces.md#repositoryowner)

Lookup a repository owner (ie. either a User or an Organization) by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](scalars.md#string)) | <p>The username to lookup the owner by.</p> |

---

### resource

#### Type: [UniformResourceLocatable](interfaces.md#uniformresourcelocatable)

Lookup resource by a URL. 

#### Arguments

| Name | Description |
|------|-------------|
| url ([URI!](scalars.md#uri)) | <p>The URL.</p> |

---

### search

#### Type: [SearchResultItemConnection!](objects.md#searchresultitemconnection)

Perform a search across resources. 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | <p>Returns the elements in the list that come after the specified cursor.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements in the list that come before the specified cursor.</p> |
| first ([Int](scalars.md#int)) | <p>Returns the first <em>n</em> elements from the list.</p> |
| last ([Int](scalars.md#int)) | <p>Returns the last <em>n</em> elements from the list.</p> |
| query ([String!](scalars.md#string)) | <p>The search string to look for.</p> |
| type ([SearchType!](enums.md#searchtype)) | <p>The types of search items to search within.</p> |

---

### securityAdvisories

#### Type: [SecurityAdvisoryConnection!](objects.md#securityadvisoryconnection)

GitHub Security Advisories 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | <p>Returns the elements in the list that come after the specified cursor.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements in the list that come before the specified cursor.</p> |
| first ([Int](scalars.md#int)) | <p>Returns the first <em>n</em> elements from the list.</p> |
| identifier ([SecurityAdvisoryIdentifierFilter](input_objects.md#securityadvisoryidentifierfilter)) | <p>Filter advisories by identifier, e.g. GHSA or CVE.</p> |
| last ([Int](scalars.md#int)) | <p>Returns the last <em>n</em> elements from the list.</p> |
| orderBy ([SecurityAdvisoryOrder](input_objects.md#securityadvisoryorder)) | <p>Ordering options for the returned topics.</p> |
| publishedSince ([DateTime](scalars.md#datetime)) | <p>Filter advisories to those published since a time in the past.</p> |
| updatedSince ([DateTime](scalars.md#datetime)) | <p>Filter advisories to those updated since a time in the past.</p> |

---

### securityAdvisory

#### Type: [SecurityAdvisory](objects.md#securityadvisory)

Fetch a Security Advisory by its GHSA ID 

#### Arguments

| Name | Description |
|------|-------------|
| ghsaId ([String!](scalars.md#string)) | <p>GitHub Security Advisory ID.</p> |

---

### securityVulnerabilities

#### Type: [SecurityVulnerabilityConnection!](objects.md#securityvulnerabilityconnection)

Software Vulnerabilities documented by GitHub Security Advisories 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | <p>Returns the elements in the list that come after the specified cursor.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements in the list that come before the specified cursor.</p> |
| ecosystem ([SecurityAdvisoryEcosystem](enums.md#securityadvisoryecosystem)) | <p>An ecosystem to filter vulnerabilities by.</p> |
| first ([Int](scalars.md#int)) | <p>Returns the first <em>n</em> elements from the list.</p> |
| last ([Int](scalars.md#int)) | <p>Returns the last <em>n</em> elements from the list.</p> |
| orderBy ([SecurityVulnerabilityOrder](input_objects.md#securityvulnerabilityorder)) | <p>Ordering options for the returned topics.</p> |
| package ([String](scalars.md#string)) | <p>A package name to filter vulnerabilities by.</p> |
| severities ([[SecurityAdvisorySeverity!]](enums.md#securityadvisoryseverity)) | <p>A list of severities to filter vulnerabilities by.</p> |

---

### sponsorables

#### Type: [SponsorableItemConnection!](objects.md#sponsorableitemconnection)

Users and organizations who can be sponsored via GitHub Sponsors. 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | <p>Returns the elements in the list that come after the specified cursor.</p> |
| before ([String](scalars.md#string)) | <p>Returns the elements in the list that come before the specified cursor.</p> |
| dependencyEcosystem ([SecurityAdvisoryEcosystem](enums.md#securityadvisoryecosystem)) | <p>Optional filter for which dependencies should be checked for sponsorable
owners. Only sponsorable owners of dependencies in this ecosystem will be
included. Used when onlyDependencies = true.</p>

<p><strong>Upcoming Change on 2022-07-01 UTC</strong>
<strong>Description:</strong> <code>dependencyEcosystem</code> will be removed. Use the ecosystem argument instead.
<strong>Reason:</strong> The type is switching from SecurityAdvisoryEcosystem to DependencyGraphEcosystem.</p> |
| ecosystem ([DependencyGraphEcosystem](enums.md#dependencygraphecosystem)) | <p>Optional filter for which dependencies should be checked for sponsorable
owners. Only sponsorable owners of dependencies in this ecosystem will be
included. Used when onlyDependencies = true.</p> |
| first ([Int](scalars.md#int)) | <p>Returns the first <em>n</em> elements from the list.</p> |
| last ([Int](scalars.md#int)) | <p>Returns the last <em>n</em> elements from the list.</p> |
| onlyDependencies ([Boolean](scalars.md#boolean)) | <p>Whether only sponsorables who own the viewer&rsquo;s dependencies will be
returned. Must be authenticated to use. Can check an organization instead
for their dependencies owned by sponsorables by passing
orgLoginForDependencies.</p> |
| orderBy ([SponsorableOrder](input_objects.md#sponsorableorder)) | <p>Ordering options for users and organizations returned from the connection.</p> |
| orgLoginForDependencies ([String](scalars.md#string)) | <p>Optional organization username for whose dependencies should be checked.
Used when onlyDependencies = true. Omit to check your own dependencies. If
you are not an administrator of the organization, only dependencies from its
public repositories will be considered.</p> |

---

### topic

#### Type: [Topic](objects.md#topic)

Look up a topic by name. 

#### Arguments

| Name | Description |
|------|-------------|
| name ([String!](scalars.md#string)) | <p>The topic&rsquo;s name.</p> |

---

### user

#### Type: [User](objects.md#user)

Lookup a user by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](scalars.md#string)) | <p>The user&rsquo;s login.</p> |

---

### viewer

#### Type: [User!](objects.md#user)

The currently authenticated user. 

---