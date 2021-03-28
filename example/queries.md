# Queries

### About queries

The query root of GitHub's GraphQL interface.

### codeOfConduct

#### Type: [CodeOfConduct](objects.md#codeofconduct)

Look up a code of conduct by its key 

#### Arguments

| Name | Description |
|------|-------------|
| key ([String!](scalars.md#string)) | The code of conduct's key |

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
| invitationToken ([String](scalars.md#string)) | The enterprise invitation token. |
| slug ([String!](scalars.md#string)) | The enterprise URL slug. |

---

### enterpriseAdministratorInvitation

#### Type: [EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)

Look up a pending enterprise administrator invitation by invitee, enterprise and role. 

#### Arguments

| Name | Description |
|------|-------------|
| enterpriseSlug ([String!](scalars.md#string)) | The slug of the enterprise the user was invited to join. |
| role ([EnterpriseAdministratorRole!](enums.md#enterpriseadministratorrole)) | The role for the business member invitation. |
| userLogin ([String!](scalars.md#string)) | The login of the user invited to join the business. |

---

### enterpriseAdministratorInvitationByToken

#### Type: [EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)

Look up a pending enterprise administrator invitation by invitation token. 

#### Arguments

| Name | Description |
|------|-------------|
| invitationToken ([String!](scalars.md#string)) | The invitation token sent with the invitation email. |

---

### license

#### Type: [License](objects.md#license)

Look up an open source license by its key 

#### Arguments

| Name | Description |
|------|-------------|
| key ([String!](scalars.md#string)) | The license's downcased SPDX ID |

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
| excludeEmpty ([Boolean](scalars.md#boolean)) | Exclude categories with no listings. |
| excludeSubcategories ([Boolean](scalars.md#boolean)) | Returns top level categories only, excluding any subcategories. |
| includeCategories ([[String!]](scalars.md#string)) | Return only the specified categories. |

---

### marketplaceCategory

#### Type: [MarketplaceCategory](objects.md#marketplacecategory)

Look up a Marketplace category by its slug. 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](scalars.md#string)) | The URL slug of the category. |
| useTopicAliases ([Boolean](scalars.md#boolean)) | Also check topic aliases for the category slug |

---

### marketplaceListing

#### Type: [MarketplaceListing](objects.md#marketplacelisting)

Look up a single Marketplace listing 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](scalars.md#string)) | Select the listing that matches this slug. It's the short name of the listing used in its URL. |

---

### marketplaceListings

#### Type: [MarketplaceListingConnection!](objects.md#marketplacelistingconnection)

Look up Marketplace listings 

#### Arguments

| Name | Description |
|------|-------------|
| adminId ([ID](scalars.md#id)) | Select listings that can be administered by the specified user. |
| after ([String](scalars.md#string)) | Returns the elements in the list that come after the specified cursor. |
| allStates ([Boolean](scalars.md#boolean)) | Select listings visible to the viewer even if they are not approved. If omitted or false, only approved listings will be returned. |
| before ([String](scalars.md#string)) | Returns the elements in the list that come before the specified cursor. |
| categorySlug ([String](scalars.md#string)) | Select only listings with the given category. |
| first ([Int](scalars.md#int)) | Returns the first _n_ elements from the list. |
| last ([Int](scalars.md#int)) | Returns the last _n_ elements from the list. |
| organizationId ([ID](scalars.md#id)) | Select listings for products owned by the specified organization. |
| primaryCategoryOnly ([Boolean](scalars.md#boolean)) | Select only listings where the primary category matches the given category slug. |
| slugs ([[String]](scalars.md#string)) | Select the listings with these slugs, if they are visible to the viewer. |
| useTopicAliases ([Boolean](scalars.md#boolean)) | Also check topic aliases for the category slug |
| viewerCanAdmin ([Boolean](scalars.md#boolean)) | Select listings to which user has admin access. If omitted, listings visible to the viewer are returned. |
| withFreeTrialsOnly ([Boolean](scalars.md#boolean)) | Select only listings that offer a free trial. |

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
| id ([ID!](scalars.md#id)) | ID of the object. |

---

### nodes

#### Type: [[Node]!](interfaces.md#node)

Lookup nodes by a list of IDs. 

#### Arguments

| Name | Description |
|------|-------------|
| ids ([[ID!]!](scalars.md#id)) | The list of node IDs. |

---

### organization

#### Type: [Organization](objects.md#organization)

Lookup a organization by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](scalars.md#string)) | The organization's login. |

---

### rateLimit

#### Type: [RateLimit](objects.md#ratelimit)

The client's rate limit information. 

#### Arguments

| Name | Description |
|------|-------------|
| dryRun ([Boolean](scalars.md#boolean)) | If true, calculate the cost for the query without evaluating it |

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
| name ([String!](scalars.md#string)) | The name of the repository |
| owner ([String!](scalars.md#string)) | The login field of a user or organization |

---

### repositoryOwner

#### Type: [RepositoryOwner](interfaces.md#repositoryowner)

Lookup a repository owner (ie. either a User or an Organization) by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](scalars.md#string)) | The username to lookup the owner by. |

---

### resource

#### Type: [UniformResourceLocatable](interfaces.md#uniformresourcelocatable)

Lookup resource by a URL. 

#### Arguments

| Name | Description |
|------|-------------|
| url ([URI!](scalars.md#uri)) | The URL. |

---

### search

#### Type: [SearchResultItemConnection!](objects.md#searchresultitemconnection)

Perform a search across resources. 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](scalars.md#string)) | Returns the elements in the list that come before the specified cursor. |
| first ([Int](scalars.md#int)) | Returns the first _n_ elements from the list. |
| last ([Int](scalars.md#int)) | Returns the last _n_ elements from the list. |
| query ([String!](scalars.md#string)) | The search string to look for. |
| type ([SearchType!](enums.md#searchtype)) | The types of search items to search within. |

---

### securityAdvisories

#### Type: [SecurityAdvisoryConnection!](objects.md#securityadvisoryconnection)

GitHub Security Advisories 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](scalars.md#string)) | Returns the elements in the list that come before the specified cursor. |
| first ([Int](scalars.md#int)) | Returns the first _n_ elements from the list. |
| identifier ([SecurityAdvisoryIdentifierFilter](input_objects.md#securityadvisoryidentifierfilter)) | Filter advisories by identifier, e.g. GHSA or CVE. |
| last ([Int](scalars.md#int)) | Returns the last _n_ elements from the list. |
| orderBy ([SecurityAdvisoryOrder](input_objects.md#securityadvisoryorder)) | Ordering options for the returned topics. |
| publishedSince ([DateTime](scalars.md#datetime)) | Filter advisories to those published since a time in the past. |
| updatedSince ([DateTime](scalars.md#datetime)) | Filter advisories to those updated since a time in the past. |

---

### securityAdvisory

#### Type: [SecurityAdvisory](objects.md#securityadvisory)

Fetch a Security Advisory by its GHSA ID 

#### Arguments

| Name | Description |
|------|-------------|
| ghsaId ([String!](scalars.md#string)) | GitHub Security Advisory ID. |

---

### securityVulnerabilities

#### Type: [SecurityVulnerabilityConnection!](objects.md#securityvulnerabilityconnection)

Software Vulnerabilities documented by GitHub Security Advisories 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](scalars.md#string)) | Returns the elements in the list that come before the specified cursor. |
| ecosystem ([SecurityAdvisoryEcosystem](enums.md#securityadvisoryecosystem)) | An ecosystem to filter vulnerabilities by. |
| first ([Int](scalars.md#int)) | Returns the first _n_ elements from the list. |
| last ([Int](scalars.md#int)) | Returns the last _n_ elements from the list. |
| orderBy ([SecurityVulnerabilityOrder](input_objects.md#securityvulnerabilityorder)) | Ordering options for the returned topics. |
| package ([String](scalars.md#string)) | A package name to filter vulnerabilities by. |
| severities ([[SecurityAdvisorySeverity!]](enums.md#securityadvisoryseverity)) | A list of severities to filter vulnerabilities by. |

---

### sponsorables

#### Type: [SponsorableItemConnection!](objects.md#sponsorableitemconnection)

Users and organizations who can be sponsored via GitHub Sponsors. 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](scalars.md#string)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](scalars.md#string)) | Returns the elements in the list that come before the specified cursor. |
| dependencyEcosystem ([SecurityAdvisoryEcosystem](enums.md#securityadvisoryecosystem)) | Optional filter for which dependencies should be checked for sponsorable owners. Only sponsorable owners of dependencies in this ecosystem will be included. Used when onlyDependencies = true. |
| first ([Int](scalars.md#int)) | Returns the first _n_ elements from the list. |
| last ([Int](scalars.md#int)) | Returns the last _n_ elements from the list. |
| onlyDependencies ([Boolean](scalars.md#boolean)) | Whether only sponsorables who own the viewer's dependencies will be returned. Must be authenticated to use. Can check an organization instead for their dependencies owned by sponsorables by passing orgLoginForDependencies. |
| orderBy ([SponsorableOrder](input_objects.md#sponsorableorder)) | Ordering options for users and organizations returned from the connection. |
| orgLoginForDependencies ([String](scalars.md#string)) | Optional organization username for whose dependencies should be checked. Used when onlyDependencies = true. Omit to check your own dependencies. If you are not an administrator of the organization, only dependencies from its public repositories will be considered. |

---

### sponsorsListing

#### Type: [SponsorsListing](objects.md#sponsorslisting)

Look up a single Sponsors Listing 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](scalars.md#string)) | Select the Sponsors listing which matches this slug |

---

### topic

#### Type: [Topic](objects.md#topic)

Look up a topic by name. 

#### Arguments

| Name | Description |
|------|-------------|
| name ([String!](scalars.md#string)) | The topic's name. |

---

### user

#### Type: [User](objects.md#user)

Lookup a user by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](scalars.md#string)) | The user's login. |

---

### viewer

#### Type: [User!](objects.md#user)

The currently authenticated user. 

---