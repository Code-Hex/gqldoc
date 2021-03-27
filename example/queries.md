# Queries

### About queries

The query root of GitHub's GraphQL interface.

### codeOfConduct

#### Type: [CodeOfConduct](http://example.com)

Look up a code of conduct by its key 

#### Arguments

| Name | Description |
|------|-------------|
| key ([String!](http://example.com)) | The code of conduct's key |

---

### codesOfConduct

#### Type: [[CodeOfConduct]](http://example.com)

Look up a code of conduct by its key 

---

### enterprise

#### Type: [Enterprise](http://example.com)

Look up an enterprise by URL slug. 

#### Arguments

| Name | Description |
|------|-------------|
| invitationToken ([String](http://example.com)) | The enterprise invitation token. |
| slug ([String!](http://example.com)) | The enterprise URL slug. |

---

### enterpriseAdministratorInvitation

#### Type: [EnterpriseAdministratorInvitation](http://example.com)

Look up a pending enterprise administrator invitation by invitee, enterprise and role. 

#### Arguments

| Name | Description |
|------|-------------|
| enterpriseSlug ([String!](http://example.com)) | The slug of the enterprise the user was invited to join. |
| role ([EnterpriseAdministratorRole!](http://example.com)) | The role for the business member invitation. |
| userLogin ([String!](http://example.com)) | The login of the user invited to join the business. |

---

### enterpriseAdministratorInvitationByToken

#### Type: [EnterpriseAdministratorInvitation](http://example.com)

Look up a pending enterprise administrator invitation by invitation token. 

#### Arguments

| Name | Description |
|------|-------------|
| invitationToken ([String!](http://example.com)) | The invitation token sent with the invitation email. |

---

### license

#### Type: [License](http://example.com)

Look up an open source license by its key 

#### Arguments

| Name | Description |
|------|-------------|
| key ([String!](http://example.com)) | The license's downcased SPDX ID |

---

### licenses

#### Type: [[License]!](http://example.com)

Return a list of known open source licenses 

---

### marketplaceCategories

#### Type: [[MarketplaceCategory!]!](http://example.com)

Get alphabetically sorted list of Marketplace categories 

#### Arguments

| Name | Description |
|------|-------------|
| excludeEmpty ([Boolean](http://example.com)) | Exclude categories with no listings. |
| excludeSubcategories ([Boolean](http://example.com)) | Returns top level categories only, excluding any subcategories. |
| includeCategories ([[String!]](http://example.com)) | Return only the specified categories. |

---

### marketplaceCategory

#### Type: [MarketplaceCategory](http://example.com)

Look up a Marketplace category by its slug. 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](http://example.com)) | The URL slug of the category. |
| useTopicAliases ([Boolean](http://example.com)) | Also check topic aliases for the category slug |

---

### marketplaceListing

#### Type: [MarketplaceListing](http://example.com)

Look up a single Marketplace listing 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](http://example.com)) | Select the listing that matches this slug. It's the short name of the listing used in its URL. |

---

### marketplaceListings

#### Type: [MarketplaceListingConnection!](http://example.com)

Look up Marketplace listings 

#### Arguments

| Name | Description |
|------|-------------|
| adminId ([ID](http://example.com)) | Select listings that can be administered by the specified user. |
| after ([String](http://example.com)) | Returns the elements in the list that come after the specified cursor. |
| allStates ([Boolean](http://example.com)) | Select listings visible to the viewer even if they are not approved. If omitted or false, only approved listings will be returned. |
| before ([String](http://example.com)) | Returns the elements in the list that come before the specified cursor. |
| categorySlug ([String](http://example.com)) | Select only listings with the given category. |
| first ([Int](http://example.com)) | Returns the first _n_ elements from the list. |
| last ([Int](http://example.com)) | Returns the last _n_ elements from the list. |
| organizationId ([ID](http://example.com)) | Select listings for products owned by the specified organization. |
| primaryCategoryOnly ([Boolean](http://example.com)) | Select only listings where the primary category matches the given category slug. |
| slugs ([[String]](http://example.com)) | Select the listings with these slugs, if they are visible to the viewer. |
| useTopicAliases ([Boolean](http://example.com)) | Also check topic aliases for the category slug |
| viewerCanAdmin ([Boolean](http://example.com)) | Select listings to which user has admin access. If omitted, listings visible to the viewer are returned. |
| withFreeTrialsOnly ([Boolean](http://example.com)) | Select only listings that offer a free trial. |

---

### meta

#### Type: [GitHubMetadata!](http://example.com)

Return information about the GitHub instance 

---

### node

#### Type: [Node](http://example.com)

Fetches an object given its ID. 

#### Arguments

| Name | Description |
|------|-------------|
| id ([ID!](http://example.com)) | ID of the object. |

---

### nodes

#### Type: [[Node]!](http://example.com)

Lookup nodes by a list of IDs. 

#### Arguments

| Name | Description |
|------|-------------|
| ids ([[ID!]!](http://example.com)) | The list of node IDs. |

---

### organization

#### Type: [Organization](http://example.com)

Lookup a organization by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](http://example.com)) | The organization's login. |

---

### rateLimit

#### Type: [RateLimit](http://example.com)

The client's rate limit information. 

#### Arguments

| Name | Description |
|------|-------------|
| dryRun ([Boolean](http://example.com)) | If true, calculate the cost for the query without evaluating it |

---

### relay

#### Type: [Query!](http://example.com)

Hack to workaround https://github.com/facebook/relay/issues/112 re-exposing the root query object 

---

### repository

#### Type: [Repository](http://example.com)

Lookup a given repository by the owner and repository name. 

#### Arguments

| Name | Description |
|------|-------------|
| name ([String!](http://example.com)) | The name of the repository |
| owner ([String!](http://example.com)) | The login field of a user or organization |

---

### repositoryOwner

#### Type: [RepositoryOwner](http://example.com)

Lookup a repository owner (ie. either a User or an Organization) by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](http://example.com)) | The username to lookup the owner by. |

---

### resource

#### Type: [UniformResourceLocatable](http://example.com)

Lookup resource by a URL. 

#### Arguments

| Name | Description |
|------|-------------|
| url ([URI!](http://example.com)) | The URL. |

---

### search

#### Type: [SearchResultItemConnection!](http://example.com)

Perform a search across resources. 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](http://example.com)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](http://example.com)) | Returns the elements in the list that come before the specified cursor. |
| first ([Int](http://example.com)) | Returns the first _n_ elements from the list. |
| last ([Int](http://example.com)) | Returns the last _n_ elements from the list. |
| query ([String!](http://example.com)) | The search string to look for. |
| type ([SearchType!](http://example.com)) | The types of search items to search within. |

---

### securityAdvisories

#### Type: [SecurityAdvisoryConnection!](http://example.com)

GitHub Security Advisories 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](http://example.com)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](http://example.com)) | Returns the elements in the list that come before the specified cursor. |
| first ([Int](http://example.com)) | Returns the first _n_ elements from the list. |
| identifier ([SecurityAdvisoryIdentifierFilter](http://example.com)) | Filter advisories by identifier, e.g. GHSA or CVE. |
| last ([Int](http://example.com)) | Returns the last _n_ elements from the list. |
| orderBy ([SecurityAdvisoryOrder](http://example.com)) | Ordering options for the returned topics. |
| publishedSince ([DateTime](http://example.com)) | Filter advisories to those published since a time in the past. |
| updatedSince ([DateTime](http://example.com)) | Filter advisories to those updated since a time in the past. |

---

### securityAdvisory

#### Type: [SecurityAdvisory](http://example.com)

Fetch a Security Advisory by its GHSA ID 

#### Arguments

| Name | Description |
|------|-------------|
| ghsaId ([String!](http://example.com)) | GitHub Security Advisory ID. |

---

### securityVulnerabilities

#### Type: [SecurityVulnerabilityConnection!](http://example.com)

Software Vulnerabilities documented by GitHub Security Advisories 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](http://example.com)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](http://example.com)) | Returns the elements in the list that come before the specified cursor. |
| ecosystem ([SecurityAdvisoryEcosystem](http://example.com)) | An ecosystem to filter vulnerabilities by. |
| first ([Int](http://example.com)) | Returns the first _n_ elements from the list. |
| last ([Int](http://example.com)) | Returns the last _n_ elements from the list. |
| orderBy ([SecurityVulnerabilityOrder](http://example.com)) | Ordering options for the returned topics. |
| package ([String](http://example.com)) | A package name to filter vulnerabilities by. |
| severities ([[SecurityAdvisorySeverity!]](http://example.com)) | A list of severities to filter vulnerabilities by. |

---

### sponsorables

#### Type: [SponsorableItemConnection!](http://example.com)

Users and organizations who can be sponsored via GitHub Sponsors. 

#### Arguments

| Name | Description |
|------|-------------|
| after ([String](http://example.com)) | Returns the elements in the list that come after the specified cursor. |
| before ([String](http://example.com)) | Returns the elements in the list that come before the specified cursor. |
| dependencyEcosystem ([SecurityAdvisoryEcosystem](http://example.com)) | Optional filter for which dependencies should be checked for sponsorable owners. Only sponsorable owners of dependencies in this ecosystem will be included. Used when onlyDependencies = true. |
| first ([Int](http://example.com)) | Returns the first _n_ elements from the list. |
| last ([Int](http://example.com)) | Returns the last _n_ elements from the list. |
| onlyDependencies ([Boolean](http://example.com)) | Whether only sponsorables who own the viewer's dependencies will be returned. Must be authenticated to use. Can check an organization instead for their dependencies owned by sponsorables by passing orgLoginForDependencies. |
| orderBy ([SponsorableOrder](http://example.com)) | Ordering options for users and organizations returned from the connection. |
| orgLoginForDependencies ([String](http://example.com)) | Optional organization username for whose dependencies should be checked. Used when onlyDependencies = true. Omit to check your own dependencies. If you are not an administrator of the organization, only dependencies from its public repositories will be considered. |

---

### sponsorsListing

#### Type: [SponsorsListing](http://example.com)

Look up a single Sponsors Listing 

#### Arguments

| Name | Description |
|------|-------------|
| slug ([String!](http://example.com)) | Select the Sponsors listing which matches this slug |

---

### topic

#### Type: [Topic](http://example.com)

Look up a topic by name. 

#### Arguments

| Name | Description |
|------|-------------|
| name ([String!](http://example.com)) | The topic's name. |

---

### user

#### Type: [User](http://example.com)

Lookup a user by login. 

#### Arguments

| Name | Description |
|------|-------------|
| login ([String!](http://example.com)) | The user's login. |

---

### viewer

#### Type: [User!](http://example.com)

The currently authenticated user. 

---