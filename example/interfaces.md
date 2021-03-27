# Interfaces

### About interfaces

[Interfaces](https://graphql.github.io/graphql-spec/June2018/#sec-Interfaces) serve as parent objects from which other objects can inherit.

### ProjectOwner

<p>Represents an owner of a Project.</p> 

#### Implemented by


- [Organization](http://example.com)
- [Repository](http://example.com)
- [User](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>project</strong> (<a href="http://example.com">Project</a>)</td> 
    <td>
      <p>Find project by number.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>number (<a href="http://example.com">Int!</a>)</p>
            <p><p>The project number to find.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>projects</strong> (<a href="http://example.com">ProjectConnection!</a>)</td> 
    <td>
      <p>A list of projects under the owner.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">ProjectOrder</a>)</p>
            <p><p>Ordering options for projects returned from the connection</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>search (<a href="http://example.com">String</a>)</p>
            <p><p>Query to search projects by, currently only searching by name.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>states (<a href="http://example.com">[ProjectState!]</a>)</p>
            <p><p>A list of states to filter the projects by.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>projectsResourcePath</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP path listing owners projects</td>
  </tr>
  <tr>
    <td><strong>projectsUrl</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP URL listing owners projects</td>
  </tr>
  <tr>
    <td><strong>viewerCanCreateProjects</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Can the current viewer create new projects on this owner.</td>
  </tr>
</table>

---

### ProfileOwner

<p>Represents any entity on GitHub that has a profile page.</p> 

#### Implemented by


- [Organization](http://example.com)
- [User](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>anyPinnableItems</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>
      <p>Determine if this repository owner has any items that can be pinned to their profile.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>type (<a href="http://example.com">PinnableItemType</a>)</p>
            <p><p>Filter to only a particular kind of pinnable item.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The public profile email.</td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>itemShowcase</strong> (<a href="http://example.com">ProfileItemShowcase!</a>)</td> 
    <td>Showcases a selection of repositories and gists that the profile owner has either curated or that have been selected automatically based on popularity.</td>
  </tr>
  <tr>
    <td><strong>location</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The public profile location.</td>
  </tr>
  <tr>
    <td><strong>login</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The username used to login.</td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The public profile name.</td>
  </tr>
  <tr>
    <td><strong>pinnableItems</strong> (<a href="http://example.com">PinnableItemConnection!</a>)</td> 
    <td>
      <p>A list of repositories and gists this profile owner can pin to their profile.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>types (<a href="http://example.com">[PinnableItemType!]</a>)</p>
            <p><p>Filter the types of pinnable items that are returned.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>pinnedItems</strong> (<a href="http://example.com">PinnableItemConnection!</a>)</td> 
    <td>
      <p>A list of repositories and gists this profile owner has pinned to their profile</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>types (<a href="http://example.com">[PinnableItemType!]</a>)</p>
            <p><p>Filter the types of pinned items that are returned.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>pinnedItemsRemaining</strong> (<a href="http://example.com">Int!</a>)</td> 
    <td>Returns how many more items this profile owner can pin to their profile.</td>
  </tr>
  <tr>
    <td><strong>viewerCanChangePinnedItems</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Can the viewer pin repositories and gists to the profile?</td>
  </tr>
  <tr>
    <td><strong>websiteUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The public profile website URL.</td>
  </tr>
</table>

---

### Contribution

<p>Represents a contribution a user made on GitHub, such as opening an issue.</p> 

#### Implemented by


- [CreatedCommitContribution](http://example.com)
- [CreatedIssueContribution](http://example.com)
- [CreatedPullRequestContribution](http://example.com)
- [CreatedPullRequestReviewContribution](http://example.com)
- [CreatedRepositoryContribution](http://example.com)
- [JoinedGitHubContribution](http://example.com)
- [RestrictedContribution](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>isRestricted</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Whether this contribution is associated with a record you do not have access to. For example, your own 'first issue' contribution may have been made on a repository you can no longer access.</td>
  </tr>
  <tr>
    <td><strong>occurredAt</strong> (<a href="http://example.com">DateTime!</a>)</td> 
    <td>When this contribution was made.</td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP path for this contribution.</td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP URL for this contribution.</td>
  </tr>
  <tr>
    <td><strong>user</strong> (<a href="http://example.com">User!</a>)</td> 
    <td>The user who made this contribution.</td>
  </tr>
</table>

---

### Lockable

<p>An object that can be locked.</p> 

#### Implemented by


- [Issue](http://example.com)
- [PullRequest](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>activeLockReason</strong> (<a href="http://example.com">LockReason</a>)</td> 
    <td>Reason that the conversation was locked.</td>
  </tr>
  <tr>
    <td><strong>locked</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>`true` if the object is locked</td>
  </tr>
</table>

---

### HovercardContext

<p>An individual line of a hovercard</p> 

#### Implemented by


- [GenericHovercardContext](http://example.com)
- [OrganizationTeamsHovercardContext](http://example.com)
- [OrganizationsHovercardContext](http://example.com)
- [ReviewStatusHovercardContext](http://example.com)
- [ViewerHovercardContext](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>message</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>A string describing this context</td>
  </tr>
  <tr>
    <td><strong>octicon</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>An octicon to accompany this context</td>
  </tr>
</table>

---

### MemberStatusable

<p>Entities that have members who can set status messages.</p> 

#### Implemented by


- [Organization](http://example.com)
- [Team](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>memberStatuses</strong> (<a href="http://example.com">UserStatusConnection!</a>)</td> 
    <td>
      <p>Get the status messages members of this entity have set that are either public or visible only to the organization.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">UserStatusOrder</a>)</p>
            <p><p>Ordering options for user statuses returned from the connection.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### UniformResourceLocatable

<p>Represents a type that can be retrieved by a URL.</p> 

#### Implemented by


- [Bot](http://example.com)
- [CheckRun](http://example.com)
- [ClosedEvent](http://example.com)
- [Commit](http://example.com)
- [ConvertToDraftEvent](http://example.com)
- [CrossReferencedEvent](http://example.com)
- [Gist](http://example.com)
- [Issue](http://example.com)
- [Mannequin](http://example.com)
- [MergedEvent](http://example.com)
- [Milestone](http://example.com)
- [Organization](http://example.com)
- [PullRequest](http://example.com)
- [PullRequestCommit](http://example.com)
- [ReadyForReviewEvent](http://example.com)
- [Release](http://example.com)
- [Repository](http://example.com)
- [RepositoryTopic](http://example.com)
- [ReviewDismissedEvent](http://example.com)
- [TeamDiscussion](http://example.com)
- [TeamDiscussionComment](http://example.com)
- [User](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTML path to this resource.</td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The URL to this resource.</td>
  </tr>
</table>

---

### Updatable

<p>Entities that can be updated.</p> 

#### Implemented by


- [CommitComment](http://example.com)
- [GistComment](http://example.com)
- [Issue](http://example.com)
- [IssueComment](http://example.com)
- [Project](http://example.com)
- [PullRequest](http://example.com)
- [PullRequestReview](http://example.com)
- [PullRequestReviewComment](http://example.com)
- [TeamDiscussion](http://example.com)
- [TeamDiscussionComment](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>viewerCanUpdate</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Check if the current viewer can update this object.</td>
  </tr>
</table>

---

### Starrable

<p>Things that can be starred.</p> 

#### Implemented by


- [Gist](http://example.com)
- [Repository](http://example.com)
- [Topic](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>stargazerCount</strong> (<a href="http://example.com">Int!</a>)</td> 
    <td>Returns a count of how many stargazers there are on this object</td>
  </tr>
  <tr>
    <td><strong>stargazers</strong> (<a href="http://example.com">StargazerConnection!</a>)</td> 
    <td>
      <p>A list of users who have starred this starrable.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">StarOrder</a>)</p>
            <p><p>Order for connection</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerHasStarred</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Returns a boolean indicating whether the viewing user has starred this starrable.</td>
  </tr>
</table>

---

### Assignable

<p>An object that can have users assigned to it.</p> 

#### Implemented by


- [Issue](http://example.com)
- [PullRequest](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>assignees</strong> (<a href="http://example.com">UserConnection!</a>)</td> 
    <td>
      <p>A list of Users assigned to this object.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### Closable

<p>An object that can be closed</p> 

#### Implemented by


- [Issue](http://example.com)
- [Milestone](http://example.com)
- [Project](http://example.com)
- [PullRequest](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>closed</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>`true` if the object is closed (definition of closed may depend on type)</td>
  </tr>
  <tr>
    <td><strong>closedAt</strong> (<a href="http://example.com">DateTime</a>)</td> 
    <td>Identifies the date and time when the object was closed.</td>
  </tr>
</table>

---

### TopicAuditEntryData

<p>Metadata for an audit entry with a topic.</p> 

#### Implemented by


- [RepoAddTopicAuditEntry](http://example.com)
- [RepoRemoveTopicAuditEntry](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>topic</strong> (<a href="http://example.com">Topic</a>)</td> 
    <td>The name of the topic added to the repository</td>
  </tr>
  <tr>
    <td><strong>topicName</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The name of the topic added to the repository</td>
  </tr>
</table>

---

### OauthApplicationAuditEntryData

<p>Metadata for an audit entry with action oauth_application.*</p> 

#### Implemented by


- [OauthApplicationCreateAuditEntry](http://example.com)
- [OrgOauthAppAccessApprovedAuditEntry](http://example.com)
- [OrgOauthAppAccessDeniedAuditEntry](http://example.com)
- [OrgOauthAppAccessRequestedAuditEntry](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>oauthApplicationName</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The name of the OAuth Application.</td>
  </tr>
  <tr>
    <td><strong>oauthApplicationResourcePath</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP path for the OAuth Application</td>
  </tr>
  <tr>
    <td><strong>oauthApplicationUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP URL for the OAuth Application</td>
  </tr>
</table>

---

### Sponsorable

<p>Entities that can be sponsored through GitHub Sponsors</p> 

#### Implemented by


- [Organization](http://example.com)
- [User](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>hasSponsorsListing</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>True if this user/organization has a GitHub Sponsors listing.</td>
  </tr>
  <tr>
    <td><strong>isSponsoredBy</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>
      <p>Check if the given account is sponsoring this user/organization.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>accountLogin (<a href="http://example.com">String!</a>)</p>
            <p><p>The target account&rsquo;s login.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>isSponsoringViewer</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>True if the viewer is sponsored by this user/organization.</td>
  </tr>
  <tr>
    <td><strong>sponsorsListing</strong> (<a href="http://example.com">SponsorsListing</a>)</td> 
    <td>The GitHub Sponsors listing for this user or organization.</td>
  </tr>
  <tr>
    <td><strong>sponsorshipForViewerAsSponsor</strong> (<a href="http://example.com">Sponsorship</a>)</td> 
    <td>The viewer's sponsorship of this entity.</td>
  </tr>
  <tr>
    <td><strong>sponsorshipsAsMaintainer</strong> (<a href="http://example.com">SponsorshipConnection!</a>)</td> 
    <td>
      <p>This object's sponsorships as the maintainer.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>includePrivate (<a href="http://example.com">Boolean</a>)</p>
            <p><p>Whether or not to include private sponsorships in the result set</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">SponsorshipOrder</a>)</p>
            <p><p>Ordering options for sponsorships returned from this connection. If left
blank, the sponsorships will be ordered based on relevancy to the viewer.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>sponsorshipsAsSponsor</strong> (<a href="http://example.com">SponsorshipConnection!</a>)</td> 
    <td>
      <p>This object's sponsorships as the sponsor.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">SponsorshipOrder</a>)</p>
            <p><p>Ordering options for sponsorships returned from this connection. If left
blank, the sponsorships will be ordered based on relevancy to the viewer.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerCanSponsor</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Whether or not the viewer is able to sponsor this user/organization.</td>
  </tr>
  <tr>
    <td><strong>viewerIsSponsoring</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>True if the viewer is sponsoring this user/organization.</td>
  </tr>
</table>

---

### AuditEntry

<p>An entry in the audit log.</p> 

#### Implemented by


- [MembersCanDeleteReposClearAuditEntry](http://example.com)
- [MembersCanDeleteReposDisableAuditEntry](http://example.com)
- [MembersCanDeleteReposEnableAuditEntry](http://example.com)
- [OauthApplicationCreateAuditEntry](http://example.com)
- [OrgAddBillingManagerAuditEntry](http://example.com)
- [OrgAddMemberAuditEntry](http://example.com)
- [OrgBlockUserAuditEntry](http://example.com)
- [OrgConfigDisableCollaboratorsOnlyAuditEntry](http://example.com)
- [OrgConfigEnableCollaboratorsOnlyAuditEntry](http://example.com)
- [OrgCreateAuditEntry](http://example.com)
- [OrgDisableOauthAppRestrictionsAuditEntry](http://example.com)
- [OrgDisableSamlAuditEntry](http://example.com)
- [OrgDisableTwoFactorRequirementAuditEntry](http://example.com)
- [OrgEnableOauthAppRestrictionsAuditEntry](http://example.com)
- [OrgEnableSamlAuditEntry](http://example.com)
- [OrgEnableTwoFactorRequirementAuditEntry](http://example.com)
- [OrgInviteMemberAuditEntry](http://example.com)
- [OrgInviteToBusinessAuditEntry](http://example.com)
- [OrgOauthAppAccessApprovedAuditEntry](http://example.com)
- [OrgOauthAppAccessDeniedAuditEntry](http://example.com)
- [OrgOauthAppAccessRequestedAuditEntry](http://example.com)
- [OrgRemoveBillingManagerAuditEntry](http://example.com)
- [OrgRemoveMemberAuditEntry](http://example.com)
- [OrgRemoveOutsideCollaboratorAuditEntry](http://example.com)
- [OrgRestoreMemberAuditEntry](http://example.com)
- [OrgUnblockUserAuditEntry](http://example.com)
- [OrgUpdateDefaultRepositoryPermissionAuditEntry](http://example.com)
- [OrgUpdateMemberAuditEntry](http://example.com)
- [OrgUpdateMemberRepositoryCreationPermissionAuditEntry](http://example.com)
- [OrgUpdateMemberRepositoryInvitationPermissionAuditEntry](http://example.com)
- [PrivateRepositoryForkingDisableAuditEntry](http://example.com)
- [PrivateRepositoryForkingEnableAuditEntry](http://example.com)
- [RepoAccessAuditEntry](http://example.com)
- [RepoAddMemberAuditEntry](http://example.com)
- [RepoAddTopicAuditEntry](http://example.com)
- [RepoArchivedAuditEntry](http://example.com)
- [RepoChangeMergeSettingAuditEntry](http://example.com)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigLockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoCreateAuditEntry](http://example.com)
- [RepoDestroyAuditEntry](http://example.com)
- [RepoRemoveMemberAuditEntry](http://example.com)
- [RepoRemoveTopicAuditEntry](http://example.com)
- [RepositoryVisibilityChangeDisableAuditEntry](http://example.com)
- [RepositoryVisibilityChangeEnableAuditEntry](http://example.com)
- [TeamAddMemberAuditEntry](http://example.com)
- [TeamAddRepositoryAuditEntry](http://example.com)
- [TeamChangeParentTeamAuditEntry](http://example.com)
- [TeamRemoveMemberAuditEntry](http://example.com)
- [TeamRemoveRepositoryAuditEntry](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>action</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The action name</td>
  </tr>
  <tr>
    <td><strong>actor</strong> (<a href="http://example.com">AuditEntryActor</a>)</td> 
    <td>The user who initiated the action</td>
  </tr>
  <tr>
    <td><strong>actorIp</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The IP address of the actor</td>
  </tr>
  <tr>
    <td><strong>actorLocation</strong> (<a href="http://example.com">ActorLocation</a>)</td> 
    <td>A readable representation of the actor's location</td>
  </tr>
  <tr>
    <td><strong>actorLogin</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The username of the user who initiated the action</td>
  </tr>
  <tr>
    <td><strong>actorResourcePath</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP path for the actor.</td>
  </tr>
  <tr>
    <td><strong>actorUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP URL for the actor.</td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="http://example.com">PreciseDateTime!</a>)</td> 
    <td>The time the action was initiated</td>
  </tr>
  <tr>
    <td><strong>operationType</strong> (<a href="http://example.com">OperationType</a>)</td> 
    <td>The corresponding operation type for the action</td>
  </tr>
  <tr>
    <td><strong>user</strong> (<a href="http://example.com">User</a>)</td> 
    <td>The user affected by the action</td>
  </tr>
  <tr>
    <td><strong>userLogin</strong> (<a href="http://example.com">String</a>)</td> 
    <td>For actions involving two users, the actor is the initiator and the user is the affected user.</td>
  </tr>
  <tr>
    <td><strong>userResourcePath</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP path for the user.</td>
  </tr>
  <tr>
    <td><strong>userUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP URL for the user.</td>
  </tr>
</table>

---

### Comment

<p>Represents a comment.</p> 

#### Implemented by


- [CommitComment](http://example.com)
- [GistComment](http://example.com)
- [Issue](http://example.com)
- [IssueComment](http://example.com)
- [PullRequest](http://example.com)
- [PullRequestReview](http://example.com)
- [PullRequestReviewComment](http://example.com)
- [TeamDiscussion](http://example.com)
- [TeamDiscussionComment](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>author</strong> (<a href="http://example.com">Actor</a>)</td> 
    <td>The actor who authored the comment.</td>
  </tr>
  <tr>
    <td><strong>authorAssociation</strong> (<a href="http://example.com">CommentAuthorAssociation!</a>)</td> 
    <td>Author's association with the subject of the comment.</td>
  </tr>
  <tr>
    <td><strong>body</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The body as Markdown.</td>
  </tr>
  <tr>
    <td><strong>bodyHTML</strong> (<a href="http://example.com">HTML!</a>)</td> 
    <td>The body rendered to HTML.</td>
  </tr>
  <tr>
    <td><strong>bodyText</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The body rendered to text.</td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="http://example.com">DateTime!</a>)</td> 
    <td>Identifies the date and time when the object was created.</td>
  </tr>
  <tr>
    <td><strong>createdViaEmail</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Check if this comment was created via an email reply.</td>
  </tr>
  <tr>
    <td><strong>editor</strong> (<a href="http://example.com">Actor</a>)</td> 
    <td>The actor who edited the comment.</td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>includesCreatedEdit</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Check if this comment was edited and includes an edit with the creation data</td>
  </tr>
  <tr>
    <td><strong>lastEditedAt</strong> (<a href="http://example.com">DateTime</a>)</td> 
    <td>The moment the editor made the last edit</td>
  </tr>
  <tr>
    <td><strong>publishedAt</strong> (<a href="http://example.com">DateTime</a>)</td> 
    <td>Identifies when the comment was published at.</td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="http://example.com">DateTime!</a>)</td> 
    <td>Identifies the date and time when the object was last updated.</td>
  </tr>
  <tr>
    <td><strong>userContentEdits</strong> (<a href="http://example.com">UserContentEditConnection</a>)</td> 
    <td>
      <p>A list of edits to this content.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerDidAuthor</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Did the viewer author this comment.</td>
  </tr>
</table>

---

### Deletable

<p>Entities that can be deleted.</p> 

#### Implemented by


- [CommitComment](http://example.com)
- [GistComment](http://example.com)
- [IssueComment](http://example.com)
- [PullRequestReview](http://example.com)
- [PullRequestReviewComment](http://example.com)
- [TeamDiscussion](http://example.com)
- [TeamDiscussionComment](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>viewerCanDelete</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Check if the current viewer can delete this object.</td>
  </tr>
</table>

---

### GitObject

<p>Represents a Git object.</p> 

#### Implemented by


- [Blob](http://example.com)
- [Commit](http://example.com)
- [Tag](http://example.com)
- [Tree](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>abbreviatedOid</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>An abbreviated version of the Git object ID</td>
  </tr>
  <tr>
    <td><strong>commitResourcePath</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP path for this Git object</td>
  </tr>
  <tr>
    <td><strong>commitUrl</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP URL for this Git object</td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>oid</strong> (<a href="http://example.com">GitObjectID!</a>)</td> 
    <td>The Git object ID</td>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="http://example.com">Repository!</a>)</td> 
    <td>The Repository the Git object belongs to</td>
  </tr>
</table>

---

### EnterpriseAuditEntryData

<p>Metadata for an audit entry containing enterprise account information.</p> 

#### Implemented by


- [MembersCanDeleteReposClearAuditEntry](http://example.com)
- [MembersCanDeleteReposDisableAuditEntry](http://example.com)
- [MembersCanDeleteReposEnableAuditEntry](http://example.com)
- [OrgInviteToBusinessAuditEntry](http://example.com)
- [PrivateRepositoryForkingDisableAuditEntry](http://example.com)
- [PrivateRepositoryForkingEnableAuditEntry](http://example.com)
- [RepositoryVisibilityChangeDisableAuditEntry](http://example.com)
- [RepositoryVisibilityChangeEnableAuditEntry](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>enterpriseResourcePath</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP path for this enterprise.</td>
  </tr>
  <tr>
    <td><strong>enterpriseSlug</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The slug of the enterprise.</td>
  </tr>
  <tr>
    <td><strong>enterpriseUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP URL for this enterprise.</td>
  </tr>
</table>

---

### TeamAuditEntryData

<p>Metadata for an audit entry with action team.*</p> 

#### Implemented by


- [OrgRestoreMemberMembershipTeamAuditEntryData](http://example.com)
- [TeamAddMemberAuditEntry](http://example.com)
- [TeamAddRepositoryAuditEntry](http://example.com)
- [TeamChangeParentTeamAuditEntry](http://example.com)
- [TeamRemoveMemberAuditEntry](http://example.com)
- [TeamRemoveRepositoryAuditEntry](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>team</strong> (<a href="http://example.com">Team</a>)</td> 
    <td>The team associated with the action</td>
  </tr>
  <tr>
    <td><strong>teamName</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The name of the team</td>
  </tr>
  <tr>
    <td><strong>teamResourcePath</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP path for this team</td>
  </tr>
  <tr>
    <td><strong>teamUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP URL for this team</td>
  </tr>
</table>

---

### Actor

<p>Represents an object which can take actions on GitHub. Typically a User or Bot.</p> 

#### Implemented by


- [Bot](http://example.com)
- [EnterpriseUserAccount](http://example.com)
- [Mannequin](http://example.com)
- [Organization](http://example.com)
- [User](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>avatarUrl</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>
      <p>A URL pointing to the actor's public avatar.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>size (<a href="http://example.com">Int</a>)</p>
            <p><p>The size of the resulting square image.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>login</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The username of the actor.</td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP path for this actor.</td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP URL for this actor.</td>
  </tr>
</table>

---

### Node

<p>An object with an ID.</p> 

#### Implemented by


- [AddedToProjectEvent](http://example.com)
- [App](http://example.com)
- [AssignedEvent](http://example.com)
- [AutoMergeDisabledEvent](http://example.com)
- [AutoMergeEnabledEvent](http://example.com)
- [AutoRebaseEnabledEvent](http://example.com)
- [AutoSquashEnabledEvent](http://example.com)
- [AutomaticBaseChangeFailedEvent](http://example.com)
- [AutomaticBaseChangeSucceededEvent](http://example.com)
- [BaseRefChangedEvent](http://example.com)
- [BaseRefDeletedEvent](http://example.com)
- [BaseRefForcePushedEvent](http://example.com)
- [Blob](http://example.com)
- [Bot](http://example.com)
- [BranchProtectionRule](http://example.com)
- [CWE](http://example.com)
- [CheckRun](http://example.com)
- [CheckSuite](http://example.com)
- [ClosedEvent](http://example.com)
- [CodeOfConduct](http://example.com)
- [CommentDeletedEvent](http://example.com)
- [Commit](http://example.com)
- [CommitComment](http://example.com)
- [CommitCommentThread](http://example.com)
- [ConnectedEvent](http://example.com)
- [ConvertToDraftEvent](http://example.com)
- [ConvertedNoteToIssueEvent](http://example.com)
- [CrossReferencedEvent](http://example.com)
- [DemilestonedEvent](http://example.com)
- [DependencyGraphManifest](http://example.com)
- [DeployKey](http://example.com)
- [DeployedEvent](http://example.com)
- [Deployment](http://example.com)
- [DeploymentEnvironmentChangedEvent](http://example.com)
- [DeploymentStatus](http://example.com)
- [DisconnectedEvent](http://example.com)
- [Enterprise](http://example.com)
- [EnterpriseAdministratorInvitation](http://example.com)
- [EnterpriseIdentityProvider](http://example.com)
- [EnterpriseRepositoryInfo](http://example.com)
- [EnterpriseServerInstallation](http://example.com)
- [EnterpriseServerUserAccount](http://example.com)
- [EnterpriseServerUserAccountEmail](http://example.com)
- [EnterpriseServerUserAccountsUpload](http://example.com)
- [EnterpriseUserAccount](http://example.com)
- [ExternalIdentity](http://example.com)
- [Gist](http://example.com)
- [GistComment](http://example.com)
- [HeadRefDeletedEvent](http://example.com)
- [HeadRefForcePushedEvent](http://example.com)
- [HeadRefRestoredEvent](http://example.com)
- [IpAllowListEntry](http://example.com)
- [Issue](http://example.com)
- [IssueComment](http://example.com)
- [Label](http://example.com)
- [LabeledEvent](http://example.com)
- [Language](http://example.com)
- [License](http://example.com)
- [LockedEvent](http://example.com)
- [Mannequin](http://example.com)
- [MarkedAsDuplicateEvent](http://example.com)
- [MarketplaceCategory](http://example.com)
- [MarketplaceListing](http://example.com)
- [MembersCanDeleteReposClearAuditEntry](http://example.com)
- [MembersCanDeleteReposDisableAuditEntry](http://example.com)
- [MembersCanDeleteReposEnableAuditEntry](http://example.com)
- [MentionedEvent](http://example.com)
- [MergedEvent](http://example.com)
- [Milestone](http://example.com)
- [MilestonedEvent](http://example.com)
- [MovedColumnsInProjectEvent](http://example.com)
- [OauthApplicationCreateAuditEntry](http://example.com)
- [OrgAddBillingManagerAuditEntry](http://example.com)
- [OrgAddMemberAuditEntry](http://example.com)
- [OrgBlockUserAuditEntry](http://example.com)
- [OrgConfigDisableCollaboratorsOnlyAuditEntry](http://example.com)
- [OrgConfigEnableCollaboratorsOnlyAuditEntry](http://example.com)
- [OrgCreateAuditEntry](http://example.com)
- [OrgDisableOauthAppRestrictionsAuditEntry](http://example.com)
- [OrgDisableSamlAuditEntry](http://example.com)
- [OrgDisableTwoFactorRequirementAuditEntry](http://example.com)
- [OrgEnableOauthAppRestrictionsAuditEntry](http://example.com)
- [OrgEnableSamlAuditEntry](http://example.com)
- [OrgEnableTwoFactorRequirementAuditEntry](http://example.com)
- [OrgInviteMemberAuditEntry](http://example.com)
- [OrgInviteToBusinessAuditEntry](http://example.com)
- [OrgOauthAppAccessApprovedAuditEntry](http://example.com)
- [OrgOauthAppAccessDeniedAuditEntry](http://example.com)
- [OrgOauthAppAccessRequestedAuditEntry](http://example.com)
- [OrgRemoveBillingManagerAuditEntry](http://example.com)
- [OrgRemoveMemberAuditEntry](http://example.com)
- [OrgRemoveOutsideCollaboratorAuditEntry](http://example.com)
- [OrgRestoreMemberAuditEntry](http://example.com)
- [OrgUnblockUserAuditEntry](http://example.com)
- [OrgUpdateDefaultRepositoryPermissionAuditEntry](http://example.com)
- [OrgUpdateMemberAuditEntry](http://example.com)
- [OrgUpdateMemberRepositoryCreationPermissionAuditEntry](http://example.com)
- [OrgUpdateMemberRepositoryInvitationPermissionAuditEntry](http://example.com)
- [Organization](http://example.com)
- [OrganizationIdentityProvider](http://example.com)
- [OrganizationInvitation](http://example.com)
- [Package](http://example.com)
- [PackageFile](http://example.com)
- [PackageTag](http://example.com)
- [PackageVersion](http://example.com)
- [PinnedEvent](http://example.com)
- [PinnedIssue](http://example.com)
- [PrivateRepositoryForkingDisableAuditEntry](http://example.com)
- [PrivateRepositoryForkingEnableAuditEntry](http://example.com)
- [Project](http://example.com)
- [ProjectCard](http://example.com)
- [ProjectColumn](http://example.com)
- [PublicKey](http://example.com)
- [PullRequest](http://example.com)
- [PullRequestCommit](http://example.com)
- [PullRequestCommitCommentThread](http://example.com)
- [PullRequestReview](http://example.com)
- [PullRequestReviewComment](http://example.com)
- [PullRequestReviewThread](http://example.com)
- [Push](http://example.com)
- [PushAllowance](http://example.com)
- [Reaction](http://example.com)
- [ReadyForReviewEvent](http://example.com)
- [Ref](http://example.com)
- [ReferencedEvent](http://example.com)
- [Release](http://example.com)
- [ReleaseAsset](http://example.com)
- [RemovedFromProjectEvent](http://example.com)
- [RenamedTitleEvent](http://example.com)
- [ReopenedEvent](http://example.com)
- [RepoAccessAuditEntry](http://example.com)
- [RepoAddMemberAuditEntry](http://example.com)
- [RepoAddTopicAuditEntry](http://example.com)
- [RepoArchivedAuditEntry](http://example.com)
- [RepoChangeMergeSettingAuditEntry](http://example.com)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigLockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoCreateAuditEntry](http://example.com)
- [RepoDestroyAuditEntry](http://example.com)
- [RepoRemoveMemberAuditEntry](http://example.com)
- [RepoRemoveTopicAuditEntry](http://example.com)
- [Repository](http://example.com)
- [RepositoryInvitation](http://example.com)
- [RepositoryTopic](http://example.com)
- [RepositoryVisibilityChangeDisableAuditEntry](http://example.com)
- [RepositoryVisibilityChangeEnableAuditEntry](http://example.com)
- [RepositoryVulnerabilityAlert](http://example.com)
- [ReviewDismissalAllowance](http://example.com)
- [ReviewDismissedEvent](http://example.com)
- [ReviewRequest](http://example.com)
- [ReviewRequestRemovedEvent](http://example.com)
- [ReviewRequestedEvent](http://example.com)
- [SavedReply](http://example.com)
- [SecurityAdvisory](http://example.com)
- [SponsorsListing](http://example.com)
- [SponsorsTier](http://example.com)
- [Sponsorship](http://example.com)
- [Status](http://example.com)
- [StatusCheckRollup](http://example.com)
- [StatusContext](http://example.com)
- [SubscribedEvent](http://example.com)
- [Tag](http://example.com)
- [Team](http://example.com)
- [TeamAddMemberAuditEntry](http://example.com)
- [TeamAddRepositoryAuditEntry](http://example.com)
- [TeamChangeParentTeamAuditEntry](http://example.com)
- [TeamDiscussion](http://example.com)
- [TeamDiscussionComment](http://example.com)
- [TeamRemoveMemberAuditEntry](http://example.com)
- [TeamRemoveRepositoryAuditEntry](http://example.com)
- [Topic](http://example.com)
- [TransferredEvent](http://example.com)
- [Tree](http://example.com)
- [UnassignedEvent](http://example.com)
- [UnlabeledEvent](http://example.com)
- [UnlockedEvent](http://example.com)
- [UnmarkedAsDuplicateEvent](http://example.com)
- [UnpinnedEvent](http://example.com)
- [UnsubscribedEvent](http://example.com)
- [User](http://example.com)
- [UserBlockedEvent](http://example.com)
- [UserContentEdit](http://example.com)
- [UserStatus](http://example.com)
- [VerifiableDomain](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td>ID of the object.</td>
  </tr>
</table>

---

### OrganizationAuditEntryData

<p>Metadata for an audit entry with action org.*</p> 

#### Implemented by


- [MembersCanDeleteReposClearAuditEntry](http://example.com)
- [MembersCanDeleteReposDisableAuditEntry](http://example.com)
- [MembersCanDeleteReposEnableAuditEntry](http://example.com)
- [OauthApplicationCreateAuditEntry](http://example.com)
- [OrgAddBillingManagerAuditEntry](http://example.com)
- [OrgAddMemberAuditEntry](http://example.com)
- [OrgBlockUserAuditEntry](http://example.com)
- [OrgConfigDisableCollaboratorsOnlyAuditEntry](http://example.com)
- [OrgConfigEnableCollaboratorsOnlyAuditEntry](http://example.com)
- [OrgCreateAuditEntry](http://example.com)
- [OrgDisableOauthAppRestrictionsAuditEntry](http://example.com)
- [OrgDisableSamlAuditEntry](http://example.com)
- [OrgDisableTwoFactorRequirementAuditEntry](http://example.com)
- [OrgEnableOauthAppRestrictionsAuditEntry](http://example.com)
- [OrgEnableSamlAuditEntry](http://example.com)
- [OrgEnableTwoFactorRequirementAuditEntry](http://example.com)
- [OrgInviteMemberAuditEntry](http://example.com)
- [OrgInviteToBusinessAuditEntry](http://example.com)
- [OrgOauthAppAccessApprovedAuditEntry](http://example.com)
- [OrgOauthAppAccessDeniedAuditEntry](http://example.com)
- [OrgOauthAppAccessRequestedAuditEntry](http://example.com)
- [OrgRemoveBillingManagerAuditEntry](http://example.com)
- [OrgRemoveMemberAuditEntry](http://example.com)
- [OrgRemoveOutsideCollaboratorAuditEntry](http://example.com)
- [OrgRestoreMemberAuditEntry](http://example.com)
- [OrgRestoreMemberMembershipOrganizationAuditEntryData](http://example.com)
- [OrgUnblockUserAuditEntry](http://example.com)
- [OrgUpdateDefaultRepositoryPermissionAuditEntry](http://example.com)
- [OrgUpdateMemberAuditEntry](http://example.com)
- [OrgUpdateMemberRepositoryCreationPermissionAuditEntry](http://example.com)
- [OrgUpdateMemberRepositoryInvitationPermissionAuditEntry](http://example.com)
- [PrivateRepositoryForkingDisableAuditEntry](http://example.com)
- [PrivateRepositoryForkingEnableAuditEntry](http://example.com)
- [RepoAccessAuditEntry](http://example.com)
- [RepoAddMemberAuditEntry](http://example.com)
- [RepoAddTopicAuditEntry](http://example.com)
- [RepoArchivedAuditEntry](http://example.com)
- [RepoChangeMergeSettingAuditEntry](http://example.com)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigLockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoCreateAuditEntry](http://example.com)
- [RepoDestroyAuditEntry](http://example.com)
- [RepoRemoveMemberAuditEntry](http://example.com)
- [RepoRemoveTopicAuditEntry](http://example.com)
- [RepositoryVisibilityChangeDisableAuditEntry](http://example.com)
- [RepositoryVisibilityChangeEnableAuditEntry](http://example.com)
- [TeamAddMemberAuditEntry](http://example.com)
- [TeamAddRepositoryAuditEntry](http://example.com)
- [TeamChangeParentTeamAuditEntry](http://example.com)
- [TeamRemoveMemberAuditEntry](http://example.com)
- [TeamRemoveRepositoryAuditEntry](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>organization</strong> (<a href="http://example.com">Organization</a>)</td> 
    <td>The Organization associated with the Audit Entry.</td>
  </tr>
  <tr>
    <td><strong>organizationName</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The name of the Organization.</td>
  </tr>
  <tr>
    <td><strong>organizationResourcePath</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP path for the organization</td>
  </tr>
  <tr>
    <td><strong>organizationUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP URL for the organization</td>
  </tr>
</table>

---

### RepositoryInfo

<p>A subset of repository info.</p> 

#### Implemented by


- [Repository](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="http://example.com">DateTime!</a>)</td> 
    <td>Identifies the date and time when the object was created.</td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The description of the repository.</td>
  </tr>
  <tr>
    <td><strong>descriptionHTML</strong> (<a href="http://example.com">HTML!</a>)</td> 
    <td>The description of the repository rendered to HTML.</td>
  </tr>
  <tr>
    <td><strong>forkCount</strong> (<a href="http://example.com">Int!</a>)</td> 
    <td>Returns how many forks there are of this repository in the whole network.</td>
  </tr>
  <tr>
    <td><strong>hasIssuesEnabled</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Indicates if the repository has issues feature enabled.</td>
  </tr>
  <tr>
    <td><strong>hasProjectsEnabled</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Indicates if the repository has the Projects feature enabled.</td>
  </tr>
  <tr>
    <td><strong>hasWikiEnabled</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Indicates if the repository has wiki feature enabled.</td>
  </tr>
  <tr>
    <td><strong>homepageUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The repository's URL.</td>
  </tr>
  <tr>
    <td><strong>isArchived</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Indicates if the repository is unmaintained.</td>
  </tr>
  <tr>
    <td><strong>isFork</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Identifies if the repository is a fork.</td>
  </tr>
  <tr>
    <td><strong>isInOrganization</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Indicates if a repository is either owned by an organization, or is a private fork of an organization repository.</td>
  </tr>
  <tr>
    <td><strong>isLocked</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Indicates if the repository has been locked or not.</td>
  </tr>
  <tr>
    <td><strong>isMirror</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Identifies if the repository is a mirror.</td>
  </tr>
  <tr>
    <td><strong>isPrivate</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Identifies if the repository is private.</td>
  </tr>
  <tr>
    <td><strong>isTemplate</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Identifies if the repository is a template that can be used to generate new repositories.</td>
  </tr>
  <tr>
    <td><strong>licenseInfo</strong> (<a href="http://example.com">License</a>)</td> 
    <td>The license associated with the repository</td>
  </tr>
  <tr>
    <td><strong>lockReason</strong> (<a href="http://example.com">RepositoryLockReason</a>)</td> 
    <td>The reason the repository has been locked.</td>
  </tr>
  <tr>
    <td><strong>mirrorUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The repository's original mirror URL.</td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The name of the repository.</td>
  </tr>
  <tr>
    <td><strong>nameWithOwner</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The repository's name with owner.</td>
  </tr>
  <tr>
    <td><strong>openGraphImageUrl</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The image used to represent this repository in Open Graph data.</td>
  </tr>
  <tr>
    <td><strong>owner</strong> (<a href="http://example.com">RepositoryOwner!</a>)</td> 
    <td>The User owner of the repository.</td>
  </tr>
  <tr>
    <td><strong>pushedAt</strong> (<a href="http://example.com">DateTime</a>)</td> 
    <td>Identifies when the repository was last pushed to.</td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP path for this repository</td>
  </tr>
  <tr>
    <td><strong>shortDescriptionHTML</strong> (<a href="http://example.com">HTML!</a>)</td> 
    <td>
      <p>A description of the repository, rendered to HTML without any links in it.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>limit (<a href="http://example.com">Int</a>)</p>
            <p><p>How many characters to return.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="http://example.com">DateTime!</a>)</td> 
    <td>Identifies the date and time when the object was last updated.</td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP URL for this repository</td>
  </tr>
  <tr>
    <td><strong>usesCustomOpenGraphImage</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Whether this repository has a custom image to use with Open Graph as opposed to being represented by the owner's avatar.</td>
  </tr>
</table>

---

### Minimizable

<p>Entities that can be minimized.</p> 

#### Implemented by


- [CommitComment](http://example.com)
- [GistComment](http://example.com)
- [IssueComment](http://example.com)
- [PullRequestReviewComment](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>isMinimized</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Returns whether or not a comment has been minimized.</td>
  </tr>
  <tr>
    <td><strong>minimizedReason</strong> (<a href="http://example.com">String</a>)</td> 
    <td>Returns why the comment was minimized.</td>
  </tr>
  <tr>
    <td><strong>viewerCanMinimize</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Check if the current viewer can minimize this object.</td>
  </tr>
</table>

---

### RepositoryAuditEntryData

<p>Metadata for an audit entry with action repo.*</p> 

#### Implemented by


- [OrgRestoreMemberMembershipRepositoryAuditEntryData](http://example.com)
- [PrivateRepositoryForkingDisableAuditEntry](http://example.com)
- [PrivateRepositoryForkingEnableAuditEntry](http://example.com)
- [RepoAccessAuditEntry](http://example.com)
- [RepoAddMemberAuditEntry](http://example.com)
- [RepoAddTopicAuditEntry](http://example.com)
- [RepoArchivedAuditEntry](http://example.com)
- [RepoChangeMergeSettingAuditEntry](http://example.com)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableContributorsOnlyAuditEntry](http://example.com)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](http://example.com)
- [RepoConfigLockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](http://example.com)
- [RepoCreateAuditEntry](http://example.com)
- [RepoDestroyAuditEntry](http://example.com)
- [RepoRemoveMemberAuditEntry](http://example.com)
- [RepoRemoveTopicAuditEntry](http://example.com)
- [TeamAddRepositoryAuditEntry](http://example.com)
- [TeamRemoveRepositoryAuditEntry](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="http://example.com">Repository</a>)</td> 
    <td>The repository associated with the action</td>
  </tr>
  <tr>
    <td><strong>repositoryName</strong> (<a href="http://example.com">String</a>)</td> 
    <td>The name of the repository</td>
  </tr>
  <tr>
    <td><strong>repositoryResourcePath</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP path for the repository</td>
  </tr>
  <tr>
    <td><strong>repositoryUrl</strong> (<a href="http://example.com">URI</a>)</td> 
    <td>The HTTP URL for the repository</td>
  </tr>
</table>

---

### Reactable

<p>Represents a subject that can be reacted on.</p> 

#### Implemented by


- [CommitComment](http://example.com)
- [Issue](http://example.com)
- [IssueComment](http://example.com)
- [PullRequest](http://example.com)
- [PullRequestReview](http://example.com)
- [PullRequestReviewComment](http://example.com)
- [TeamDiscussion](http://example.com)
- [TeamDiscussionComment](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>databaseId</strong> (<a href="http://example.com">Int</a>)</td> 
    <td>Identifies the primary key from the database.</td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>reactionGroups</strong> (<a href="http://example.com">[ReactionGroup!]</a>)</td> 
    <td>A list of reactions grouped by content left on the subject.</td>
  </tr>
  <tr>
    <td><strong>reactions</strong> (<a href="http://example.com">ReactionConnection!</a>)</td> 
    <td>
      <p>A list of Reactions left on the Issue.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>content (<a href="http://example.com">ReactionContent</a>)</p>
            <p><p>Allows filtering Reactions by emoji.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">ReactionOrder</a>)</p>
            <p><p>Allows specifying the order in which reactions are returned.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerCanReact</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Can user react to this subject</td>
  </tr>
</table>

---

### Subscribable

<p>Entities that can be subscribed to for web and email notifications.</p> 

#### Implemented by


- [Commit](http://example.com)
- [Issue](http://example.com)
- [PullRequest](http://example.com)
- [Repository](http://example.com)
- [Team](http://example.com)
- [TeamDiscussion](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>viewerCanSubscribe</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>Check if the viewer is able to change their subscription status for the repository.</td>
  </tr>
  <tr>
    <td><strong>viewerSubscription</strong> (<a href="http://example.com">SubscriptionState</a>)</td> 
    <td>Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.</td>
  </tr>
</table>

---

### UpdatableComment

<p>Comments that can be updated.</p> 

#### Implemented by


- [CommitComment](http://example.com)
- [GistComment](http://example.com)
- [Issue](http://example.com)
- [IssueComment](http://example.com)
- [PullRequest](http://example.com)
- [PullRequestReview](http://example.com)
- [PullRequestReviewComment](http://example.com)
- [TeamDiscussion](http://example.com)
- [TeamDiscussionComment](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>viewerCannotUpdateReasons</strong> (<a href="http://example.com">[CommentCannotUpdateReason!]!</a>)</td> 
    <td>Reasons why the current viewer can not update this comment.</td>
  </tr>
</table>

---

### Labelable

<p>An object that can have labels assigned to it.</p> 

#### Implemented by


- [Issue](http://example.com)
- [PullRequest](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>labels</strong> (<a href="http://example.com">LabelConnection</a>)</td> 
    <td>
      <p>A list of labels associated with the object.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">LabelOrder</a>)</p>
            <p><p>Ordering options for labels returned from the connection.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### PackageOwner

<p>Represents an owner of a package.</p> 

#### Implemented by


- [Organization](http://example.com)
- [Repository](http://example.com)
- [User](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>packages</strong> (<a href="http://example.com">PackageConnection!</a>)</td> 
    <td>
      <p>A list of packages under the owner.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>names (<a href="http://example.com">[String]</a>)</p>
            <p><p>Find packages by their names.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">PackageOrder</a>)</p>
            <p><p>Ordering of the returned packages.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>packageType (<a href="http://example.com">PackageType</a>)</p>
            <p><p>Filter registry package by type.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>repositoryId (<a href="http://example.com">ID</a>)</p>
            <p><p>Find packages in a repository by ID.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### GitSignature

<p>Information about a signature (GPG or S/MIME) on a Commit or Tag.</p> 

#### Implemented by


- [GpgSignature](http://example.com)
- [SmimeSignature](http://example.com)
- [UnknownSignature](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>Email used to sign this object.</td>
  </tr>
  <tr>
    <td><strong>isValid</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>True if the signature is valid and verified by GitHub.</td>
  </tr>
  <tr>
    <td><strong>payload</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>Payload for GPG signing object. Raw ODB object without the signature header.</td>
  </tr>
  <tr>
    <td><strong>signature</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>ASCII-armored signature header from object.</td>
  </tr>
  <tr>
    <td><strong>signer</strong> (<a href="http://example.com">User</a>)</td> 
    <td>GitHub user corresponding to the email signing this commit.</td>
  </tr>
  <tr>
    <td><strong>state</strong> (<a href="http://example.com">GitSignatureState!</a>)</td> 
    <td>The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.</td>
  </tr>
  <tr>
    <td><strong>wasSignedByGitHub</strong> (<a href="http://example.com">Boolean!</a>)</td> 
    <td>True if the signature was made with GitHub's signing key.</td>
  </tr>
</table>

---

### RepositoryOwner

<p>Represents an owner of a Repository.</p> 

#### Implemented by


- [Organization](http://example.com)
- [User](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>avatarUrl</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>
      <p>A URL pointing to the owner's public avatar.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>size (<a href="http://example.com">Int</a>)</p>
            <p><p>The size of the resulting square image.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="http://example.com">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>login</strong> (<a href="http://example.com">String!</a>)</td> 
    <td>The username used to login.</td>
  </tr>
  <tr>
    <td><strong>repositories</strong> (<a href="http://example.com">RepositoryConnection!</a>)</td> 
    <td>
      <p>A list of repositories that the user owns.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>affiliations (<a href="http://example.com">[RepositoryAffiliation]</a>)</p>
            <p><p>Array of viewer&rsquo;s affiliation options for repositories returned from the
connection. For example, OWNER will include only repositories that the
current viewer owns.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="http://example.com">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>isFork (<a href="http://example.com">Boolean</a>)</p>
            <p><p>If non-null, filters repositories according to whether they are forks of another repository</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>isLocked (<a href="http://example.com">Boolean</a>)</p>
            <p><p>If non-null, filters repositories according to whether they have been locked</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="http://example.com">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="http://example.com">RepositoryOrder</a>)</p>
            <p><p>Ordering options for repositories returned from the connection</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>ownerAffiliations (<a href="http://example.com">[RepositoryAffiliation]</a>)</p>
            <p><p>Array of owner&rsquo;s affiliation options for repositories returned from the
connection. For example, OWNER will include only repositories that the
organization or user being viewed owns.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>privacy (<a href="http://example.com">RepositoryPrivacy</a>)</p>
            <p><p>If non-null, filters repositories according to privacy</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="http://example.com">Repository</a>)</td> 
    <td>
      <p>Find Repository.</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>name (<a href="http://example.com">String!</a>)</p>
            <p><p>Name of Repository to find.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP URL for the owner.</td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="http://example.com">URI!</a>)</td> 
    <td>The HTTP URL for the owner.</td>
  </tr>
</table>

---

### RepositoryNode

<p>Represents a object that belongs to a repository.</p> 

#### Implemented by


- [CommitComment](http://example.com)
- [CommitCommentThread](http://example.com)
- [Issue](http://example.com)
- [IssueComment](http://example.com)
- [PullRequest](http://example.com)
- [PullRequestCommitCommentThread](http://example.com)
- [PullRequestReview](http://example.com)
- [PullRequestReviewComment](http://example.com)
- [RepositoryVulnerabilityAlert](http://example.com) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="http://example.com">Repository!</a>)</td> 
    <td>The repository associated with this node.</td>
  </tr>
</table>

---