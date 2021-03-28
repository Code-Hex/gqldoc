# Interfaces

### About interfaces

[Interfaces](https://graphql.github.io/graphql-spec/June2018/#sec-Interfaces) serve as parent objects from which other objects can inherit.

### Minimizable

<p>Entities that can be minimized.</p> 

#### Implemented by


- [CommitComment](objects.md#commitcomment)
- [GistComment](objects.md#gistcomment)
- [IssueComment](objects.md#issuecomment)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>isMinimized</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Returns whether or not a comment has been minimized.</p></td>
  </tr>
  <tr>
    <td><strong>minimizedReason</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>Returns why the comment was minimized.</p></td>
  </tr>
  <tr>
    <td><strong>viewerCanMinimize</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Check if the current viewer can minimize this object.</p></td>
  </tr>
</table>

---

### Starrable

<p>Things that can be starred.</p> 

#### Implemented by


- [Gist](objects.md#gist)
- [Repository](objects.md#repository)
- [Topic](objects.md#topic) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>stargazerCount</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>Returns a count of how many stargazers there are on this object</p></td>
  </tr>
  <tr>
    <td><strong>stargazers</strong> (<a href="objects.md#stargazerconnection">StargazerConnection!</a>)</td> 
    <td>
      <p><p>A list of users who have starred this starrable.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#starorder">StarOrder</a>)</p>
            <p><p>Order for connection</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerHasStarred</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Returns a boolean indicating whether the viewing user has starred this starrable.</p></td>
  </tr>
</table>

---

### RepositoryAuditEntryData

<p>Metadata for an audit entry with action repo.*</p> 

#### Implemented by


- [OrgRestoreMemberMembershipRepositoryAuditEntryData](objects.md#orgrestoremembermembershiprepositoryauditentrydata)
- [PrivateRepositoryForkingDisableAuditEntry](objects.md#privaterepositoryforkingdisableauditentry)
- [PrivateRepositoryForkingEnableAuditEntry](objects.md#privaterepositoryforkingenableauditentry)
- [RepoAccessAuditEntry](objects.md#repoaccessauditentry)
- [RepoAddMemberAuditEntry](objects.md#repoaddmemberauditentry)
- [RepoAddTopicAuditEntry](objects.md#repoaddtopicauditentry)
- [RepoArchivedAuditEntry](objects.md#repoarchivedauditentry)
- [RepoChangeMergeSettingAuditEntry](objects.md#repochangemergesettingauditentry)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](objects.md#repoconfigdisableanonymousgitaccessauditentry)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](objects.md#repoconfigdisablecollaboratorsonlyauditentry)
- [RepoConfigDisableContributorsOnlyAuditEntry](objects.md#repoconfigdisablecontributorsonlyauditentry)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](objects.md#repoconfigdisablesockpuppetdisallowedauditentry)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](objects.md#repoconfigenableanonymousgitaccessauditentry)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](objects.md#repoconfigenablecollaboratorsonlyauditentry)
- [RepoConfigEnableContributorsOnlyAuditEntry](objects.md#repoconfigenablecontributorsonlyauditentry)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](objects.md#repoconfigenablesockpuppetdisallowedauditentry)
- [RepoConfigLockAnonymousGitAccessAuditEntry](objects.md#repoconfiglockanonymousgitaccessauditentry)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](objects.md#repoconfigunlockanonymousgitaccessauditentry)
- [RepoCreateAuditEntry](objects.md#repocreateauditentry)
- [RepoDestroyAuditEntry](objects.md#repodestroyauditentry)
- [RepoRemoveMemberAuditEntry](objects.md#reporemovememberauditentry)
- [RepoRemoveTopicAuditEntry](objects.md#reporemovetopicauditentry)
- [TeamAddRepositoryAuditEntry](objects.md#teamaddrepositoryauditentry)
- [TeamRemoveRepositoryAuditEntry](objects.md#teamremoverepositoryauditentry) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="objects.md#repository">Repository</a>)</td> 
    <td><p>The repository associated with the action</p></td>
  </tr>
  <tr>
    <td><strong>repositoryName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of the repository</p></td>
  </tr>
  <tr>
    <td><strong>repositoryResourcePath</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP path for the repository</p></td>
  </tr>
  <tr>
    <td><strong>repositoryUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP URL for the repository</p></td>
  </tr>
</table>

---

### MemberStatusable

<p>Entities that have members who can set status messages.</p> 

#### Implemented by


- [Organization](objects.md#organization)
- [Team](objects.md#team) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>memberStatuses</strong> (<a href="objects.md#userstatusconnection">UserStatusConnection!</a>)</td> 
    <td>
      <p><p>Get the status messages members of this entity have set that are either public or visible only to the organization.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#userstatusorder">UserStatusOrder</a>)</p>
            <p><p>Ordering options for user statuses returned from the connection.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### AuditEntry

<p>An entry in the audit log.</p> 

#### Implemented by


- [MembersCanDeleteReposClearAuditEntry](objects.md#memberscandeletereposclearauditentry)
- [MembersCanDeleteReposDisableAuditEntry](objects.md#memberscandeletereposdisableauditentry)
- [MembersCanDeleteReposEnableAuditEntry](objects.md#memberscandeletereposenableauditentry)
- [OauthApplicationCreateAuditEntry](objects.md#oauthapplicationcreateauditentry)
- [OrgAddBillingManagerAuditEntry](objects.md#orgaddbillingmanagerauditentry)
- [OrgAddMemberAuditEntry](objects.md#orgaddmemberauditentry)
- [OrgBlockUserAuditEntry](objects.md#orgblockuserauditentry)
- [OrgConfigDisableCollaboratorsOnlyAuditEntry](objects.md#orgconfigdisablecollaboratorsonlyauditentry)
- [OrgConfigEnableCollaboratorsOnlyAuditEntry](objects.md#orgconfigenablecollaboratorsonlyauditentry)
- [OrgCreateAuditEntry](objects.md#orgcreateauditentry)
- [OrgDisableOauthAppRestrictionsAuditEntry](objects.md#orgdisableoauthapprestrictionsauditentry)
- [OrgDisableSamlAuditEntry](objects.md#orgdisablesamlauditentry)
- [OrgDisableTwoFactorRequirementAuditEntry](objects.md#orgdisabletwofactorrequirementauditentry)
- [OrgEnableOauthAppRestrictionsAuditEntry](objects.md#orgenableoauthapprestrictionsauditentry)
- [OrgEnableSamlAuditEntry](objects.md#orgenablesamlauditentry)
- [OrgEnableTwoFactorRequirementAuditEntry](objects.md#orgenabletwofactorrequirementauditentry)
- [OrgInviteMemberAuditEntry](objects.md#orginvitememberauditentry)
- [OrgInviteToBusinessAuditEntry](objects.md#orginvitetobusinessauditentry)
- [OrgOauthAppAccessApprovedAuditEntry](objects.md#orgoauthappaccessapprovedauditentry)
- [OrgOauthAppAccessDeniedAuditEntry](objects.md#orgoauthappaccessdeniedauditentry)
- [OrgOauthAppAccessRequestedAuditEntry](objects.md#orgoauthappaccessrequestedauditentry)
- [OrgRemoveBillingManagerAuditEntry](objects.md#orgremovebillingmanagerauditentry)
- [OrgRemoveMemberAuditEntry](objects.md#orgremovememberauditentry)
- [OrgRemoveOutsideCollaboratorAuditEntry](objects.md#orgremoveoutsidecollaboratorauditentry)
- [OrgRestoreMemberAuditEntry](objects.md#orgrestorememberauditentry)
- [OrgUnblockUserAuditEntry](objects.md#orgunblockuserauditentry)
- [OrgUpdateDefaultRepositoryPermissionAuditEntry](objects.md#orgupdatedefaultrepositorypermissionauditentry)
- [OrgUpdateMemberAuditEntry](objects.md#orgupdatememberauditentry)
- [OrgUpdateMemberRepositoryCreationPermissionAuditEntry](objects.md#orgupdatememberrepositorycreationpermissionauditentry)
- [OrgUpdateMemberRepositoryInvitationPermissionAuditEntry](objects.md#orgupdatememberrepositoryinvitationpermissionauditentry)
- [PrivateRepositoryForkingDisableAuditEntry](objects.md#privaterepositoryforkingdisableauditentry)
- [PrivateRepositoryForkingEnableAuditEntry](objects.md#privaterepositoryforkingenableauditentry)
- [RepoAccessAuditEntry](objects.md#repoaccessauditentry)
- [RepoAddMemberAuditEntry](objects.md#repoaddmemberauditentry)
- [RepoAddTopicAuditEntry](objects.md#repoaddtopicauditentry)
- [RepoArchivedAuditEntry](objects.md#repoarchivedauditentry)
- [RepoChangeMergeSettingAuditEntry](objects.md#repochangemergesettingauditentry)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](objects.md#repoconfigdisableanonymousgitaccessauditentry)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](objects.md#repoconfigdisablecollaboratorsonlyauditentry)
- [RepoConfigDisableContributorsOnlyAuditEntry](objects.md#repoconfigdisablecontributorsonlyauditentry)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](objects.md#repoconfigdisablesockpuppetdisallowedauditentry)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](objects.md#repoconfigenableanonymousgitaccessauditentry)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](objects.md#repoconfigenablecollaboratorsonlyauditentry)
- [RepoConfigEnableContributorsOnlyAuditEntry](objects.md#repoconfigenablecontributorsonlyauditentry)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](objects.md#repoconfigenablesockpuppetdisallowedauditentry)
- [RepoConfigLockAnonymousGitAccessAuditEntry](objects.md#repoconfiglockanonymousgitaccessauditentry)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](objects.md#repoconfigunlockanonymousgitaccessauditentry)
- [RepoCreateAuditEntry](objects.md#repocreateauditentry)
- [RepoDestroyAuditEntry](objects.md#repodestroyauditentry)
- [RepoRemoveMemberAuditEntry](objects.md#reporemovememberauditentry)
- [RepoRemoveTopicAuditEntry](objects.md#reporemovetopicauditentry)
- [RepositoryVisibilityChangeDisableAuditEntry](objects.md#repositoryvisibilitychangedisableauditentry)
- [RepositoryVisibilityChangeEnableAuditEntry](objects.md#repositoryvisibilitychangeenableauditentry)
- [TeamAddMemberAuditEntry](objects.md#teamaddmemberauditentry)
- [TeamAddRepositoryAuditEntry](objects.md#teamaddrepositoryauditentry)
- [TeamChangeParentTeamAuditEntry](objects.md#teamchangeparentteamauditentry)
- [TeamRemoveMemberAuditEntry](objects.md#teamremovememberauditentry)
- [TeamRemoveRepositoryAuditEntry](objects.md#teamremoverepositoryauditentry) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>action</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The action name</p></td>
  </tr>
  <tr>
    <td><strong>actor</strong> (<a href="unions.md#auditentryactor">AuditEntryActor</a>)</td> 
    <td><p>The user who initiated the action</p></td>
  </tr>
  <tr>
    <td><strong>actorIp</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The IP address of the actor</p></td>
  </tr>
  <tr>
    <td><strong>actorLocation</strong> (<a href="objects.md#actorlocation">ActorLocation</a>)</td> 
    <td><p>A readable representation of the actor&rsquo;s location</p></td>
  </tr>
  <tr>
    <td><strong>actorLogin</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The username of the user who initiated the action</p></td>
  </tr>
  <tr>
    <td><strong>actorResourcePath</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP path for the actor.</p></td>
  </tr>
  <tr>
    <td><strong>actorUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP URL for the actor.</p></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#precisedatetime">PreciseDateTime!</a>)</td> 
    <td><p>The time the action was initiated</p></td>
  </tr>
  <tr>
    <td><strong>operationType</strong> (<a href="enums.md#operationtype">OperationType</a>)</td> 
    <td><p>The corresponding operation type for the action</p></td>
  </tr>
  <tr>
    <td><strong>user</strong> (<a href="objects.md#user">User</a>)</td> 
    <td><p>The user affected by the action</p></td>
  </tr>
  <tr>
    <td><strong>userLogin</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>For actions involving two users, the actor is the initiator and the user is the affected user.</p></td>
  </tr>
  <tr>
    <td><strong>userResourcePath</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP path for the user.</p></td>
  </tr>
  <tr>
    <td><strong>userUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP URL for the user.</p></td>
  </tr>
</table>

---

### ProjectOwner

<p>Represents an owner of a Project.</p> 

#### Implemented by


- [Organization](objects.md#organization)
- [Repository](objects.md#repository)
- [User](objects.md#user) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>project</strong> (<a href="objects.md#project">Project</a>)</td> 
    <td>
      <p><p>Find project by number.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>number (<a href="scalars.md#int">Int!</a>)</p>
            <p><p>The project number to find.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>projects</strong> (<a href="objects.md#projectconnection">ProjectConnection!</a>)</td> 
    <td>
      <p><p>A list of projects under the owner.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#projectorder">ProjectOrder</a>)</p>
            <p><p>Ordering options for projects returned from the connection</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>search (<a href="scalars.md#string">String</a>)</p>
            <p><p>Query to search projects by, currently only searching by name.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>states (<a href="enums.md#projectstate">[ProjectState!]</a>)</p>
            <p><p>A list of states to filter the projects by.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>projectsResourcePath</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP path listing owners projects</p></td>
  </tr>
  <tr>
    <td><strong>projectsUrl</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP URL listing owners projects</p></td>
  </tr>
  <tr>
    <td><strong>viewerCanCreateProjects</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Can the current viewer create new projects on this owner.</p></td>
  </tr>
</table>

---

### UpdatableComment

<p>Comments that can be updated.</p> 

#### Implemented by


- [CommitComment](objects.md#commitcomment)
- [GistComment](objects.md#gistcomment)
- [Issue](objects.md#issue)
- [IssueComment](objects.md#issuecomment)
- [PullRequest](objects.md#pullrequest)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [TeamDiscussion](objects.md#teamdiscussion)
- [TeamDiscussionComment](objects.md#teamdiscussioncomment) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>viewerCannotUpdateReasons</strong> (<a href="enums.md#commentcannotupdatereason">[CommentCannotUpdateReason!]!</a>)</td> 
    <td><p>Reasons why the current viewer can not update this comment.</p></td>
  </tr>
</table>

---

### Deletable

<p>Entities that can be deleted.</p> 

#### Implemented by


- [CommitComment](objects.md#commitcomment)
- [GistComment](objects.md#gistcomment)
- [IssueComment](objects.md#issuecomment)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [TeamDiscussion](objects.md#teamdiscussion)
- [TeamDiscussionComment](objects.md#teamdiscussioncomment) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>viewerCanDelete</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Check if the current viewer can delete this object.</p></td>
  </tr>
</table>

---

### Sponsorable

<p>Entities that can be sponsored through GitHub Sponsors</p> 

#### Implemented by


- [Organization](objects.md#organization)
- [User](objects.md#user) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>hasSponsorsListing</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>True if this user/organization has a GitHub Sponsors listing.</p></td>
  </tr>
  <tr>
    <td><strong>isSponsoredBy</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td>
      <p><p>Check if the given account is sponsoring this user/organization.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>accountLogin (<a href="scalars.md#string">String!</a>)</p>
            <p><p>The target account&rsquo;s login.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>isSponsoringViewer</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>True if the viewer is sponsored by this user/organization.</p></td>
  </tr>
  <tr>
    <td><strong>sponsorsListing</strong> (<a href="objects.md#sponsorslisting">SponsorsListing</a>)</td> 
    <td><p>The GitHub Sponsors listing for this user or organization.</p></td>
  </tr>
  <tr>
    <td><strong>sponsorshipForViewerAsSponsor</strong> (<a href="objects.md#sponsorship">Sponsorship</a>)</td> 
    <td><p>The viewer&rsquo;s sponsorship of this entity.</p></td>
  </tr>
  <tr>
    <td><strong>sponsorshipsAsMaintainer</strong> (<a href="objects.md#sponsorshipconnection">SponsorshipConnection!</a>)</td> 
    <td>
      <p><p>This object&rsquo;s sponsorships as the maintainer.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>includePrivate (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>Whether or not to include private sponsorships in the result set</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#sponsorshiporder">SponsorshipOrder</a>)</p>
            <p><p>Ordering options for sponsorships returned from this connection. If left
blank, the sponsorships will be ordered based on relevancy to the viewer.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>sponsorshipsAsSponsor</strong> (<a href="objects.md#sponsorshipconnection">SponsorshipConnection!</a>)</td> 
    <td>
      <p><p>This object&rsquo;s sponsorships as the sponsor.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#sponsorshiporder">SponsorshipOrder</a>)</p>
            <p><p>Ordering options for sponsorships returned from this connection. If left
blank, the sponsorships will be ordered based on relevancy to the viewer.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerCanSponsor</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether or not the viewer is able to sponsor this user/organization.</p></td>
  </tr>
  <tr>
    <td><strong>viewerIsSponsoring</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>True if the viewer is sponsoring this user/organization.</p></td>
  </tr>
</table>

---

### RepositoryOwner

<p>Represents an owner of a Repository.</p> 

#### Implemented by


- [Organization](objects.md#organization)
- [User](objects.md#user) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>avatarUrl</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td>
      <p><p>A URL pointing to the owner&rsquo;s public avatar.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>size (<a href="scalars.md#int">Int</a>)</p>
            <p><p>The size of the resulting square image.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>login</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The username used to login.</p></td>
  </tr>
  <tr>
    <td><strong>repositories</strong> (<a href="objects.md#repositoryconnection">RepositoryConnection!</a>)</td> 
    <td>
      <p><p>A list of repositories that the user owns.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>affiliations (<a href="enums.md#repositoryaffiliation">[RepositoryAffiliation]</a>)</p>
            <p><p>Array of viewer&rsquo;s affiliation options for repositories returned from the
connection. For example, OWNER will include only repositories that the
current viewer owns.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>isFork (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>If non-null, filters repositories according to whether they are forks of another repository</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>isLocked (<a href="scalars.md#boolean">Boolean</a>)</p>
            <p><p>If non-null, filters repositories according to whether they have been locked</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#repositoryorder">RepositoryOrder</a>)</p>
            <p><p>Ordering options for repositories returned from the connection</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>ownerAffiliations (<a href="enums.md#repositoryaffiliation">[RepositoryAffiliation]</a>)</p>
            <p><p>Array of owner&rsquo;s affiliation options for repositories returned from the
connection. For example, OWNER will include only repositories that the
organization or user being viewed owns.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>privacy (<a href="enums.md#repositoryprivacy">RepositoryPrivacy</a>)</p>
            <p><p>If non-null, filters repositories according to privacy</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="objects.md#repository">Repository</a>)</td> 
    <td>
      <p><p>Find Repository.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>name (<a href="scalars.md#string">String!</a>)</p>
            <p><p>Name of Repository to find.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP URL for the owner.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP URL for the owner.</p></td>
  </tr>
</table>

---

### Actor

<p>Represents an object which can take actions on GitHub. Typically a User or Bot.</p> 

#### Implemented by


- [Bot](objects.md#bot)
- [EnterpriseUserAccount](objects.md#enterpriseuseraccount)
- [Mannequin](objects.md#mannequin)
- [Organization](objects.md#organization)
- [User](objects.md#user) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>avatarUrl</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td>
      <p><p>A URL pointing to the actor&rsquo;s public avatar.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>size (<a href="scalars.md#int">Int</a>)</p>
            <p><p>The size of the resulting square image.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>login</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The username of the actor.</p></td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP path for this actor.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP URL for this actor.</p></td>
  </tr>
</table>

---

### ProfileOwner

<p>Represents any entity on GitHub that has a profile page.</p> 

#### Implemented by


- [Organization](objects.md#organization)
- [User](objects.md#user) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>anyPinnableItems</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td>
      <p><p>Determine if this repository owner has any items that can be pinned to their profile.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>type (<a href="enums.md#pinnableitemtype">PinnableItemType</a>)</p>
            <p><p>Filter to only a particular kind of pinnable item.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The public profile email.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>itemShowcase</strong> (<a href="objects.md#profileitemshowcase">ProfileItemShowcase!</a>)</td> 
    <td><p>Showcases a selection of repositories and gists that the profile owner has
either curated or that have been selected automatically based on popularity.</p></td>
  </tr>
  <tr>
    <td><strong>location</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The public profile location.</p></td>
  </tr>
  <tr>
    <td><strong>login</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The username used to login.</p></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The public profile name.</p></td>
  </tr>
  <tr>
    <td><strong>pinnableItems</strong> (<a href="objects.md#pinnableitemconnection">PinnableItemConnection!</a>)</td> 
    <td>
      <p><p>A list of repositories and gists this profile owner can pin to their profile.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>types (<a href="enums.md#pinnableitemtype">[PinnableItemType!]</a>)</p>
            <p><p>Filter the types of pinnable items that are returned.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>pinnedItems</strong> (<a href="objects.md#pinnableitemconnection">PinnableItemConnection!</a>)</td> 
    <td>
      <p><p>A list of repositories and gists this profile owner has pinned to their profile</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>types (<a href="enums.md#pinnableitemtype">[PinnableItemType!]</a>)</p>
            <p><p>Filter the types of pinned items that are returned.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>pinnedItemsRemaining</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>Returns how many more items this profile owner can pin to their profile.</p></td>
  </tr>
  <tr>
    <td><strong>viewerCanChangePinnedItems</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Can the viewer pin repositories and gists to the profile?</p></td>
  </tr>
  <tr>
    <td><strong>websiteUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The public profile website URL.</p></td>
  </tr>
</table>

---

### Contribution

<p>Represents a contribution a user made on GitHub, such as opening an issue.</p> 

#### Implemented by


- [CreatedCommitContribution](objects.md#createdcommitcontribution)
- [CreatedIssueContribution](objects.md#createdissuecontribution)
- [CreatedPullRequestContribution](objects.md#createdpullrequestcontribution)
- [CreatedPullRequestReviewContribution](objects.md#createdpullrequestreviewcontribution)
- [CreatedRepositoryContribution](objects.md#createdrepositorycontribution)
- [JoinedGitHubContribution](objects.md#joinedgithubcontribution)
- [RestrictedContribution](objects.md#restrictedcontribution) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>isRestricted</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether this contribution is associated with a record you do not have access to. For
example, your own &lsquo;first issue&rsquo; contribution may have been made on a repository you can no
longer access.</p></td>
  </tr>
  <tr>
    <td><strong>occurredAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>When this contribution was made.</p></td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP path for this contribution.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP URL for this contribution.</p></td>
  </tr>
  <tr>
    <td><strong>user</strong> (<a href="objects.md#user">User!</a>)</td> 
    <td><p>The user who made this contribution.</p></td>
  </tr>
</table>

---

### Comment

<p>Represents a comment.</p> 

#### Implemented by


- [CommitComment](objects.md#commitcomment)
- [GistComment](objects.md#gistcomment)
- [Issue](objects.md#issue)
- [IssueComment](objects.md#issuecomment)
- [PullRequest](objects.md#pullrequest)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [TeamDiscussion](objects.md#teamdiscussion)
- [TeamDiscussionComment](objects.md#teamdiscussioncomment) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>author</strong> (<a href="interfaces.md#actor">Actor</a>)</td> 
    <td><p>The actor who authored the comment.</p></td>
  </tr>
  <tr>
    <td><strong>authorAssociation</strong> (<a href="enums.md#commentauthorassociation">CommentAuthorAssociation!</a>)</td> 
    <td><p>Author&rsquo;s association with the subject of the comment.</p></td>
  </tr>
  <tr>
    <td><strong>body</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The body as Markdown.</p></td>
  </tr>
  <tr>
    <td><strong>bodyHTML</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td><p>The body rendered to HTML.</p></td>
  </tr>
  <tr>
    <td><strong>bodyText</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The body rendered to text.</p></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>Identifies the date and time when the object was created.</p></td>
  </tr>
  <tr>
    <td><strong>createdViaEmail</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Check if this comment was created via an email reply.</p></td>
  </tr>
  <tr>
    <td><strong>editor</strong> (<a href="interfaces.md#actor">Actor</a>)</td> 
    <td><p>The actor who edited the comment.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>includesCreatedEdit</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Check if this comment was edited and includes an edit with the creation data</p></td>
  </tr>
  <tr>
    <td><strong>lastEditedAt</strong> (<a href="scalars.md#datetime">DateTime</a>)</td> 
    <td><p>The moment the editor made the last edit</p></td>
  </tr>
  <tr>
    <td><strong>publishedAt</strong> (<a href="scalars.md#datetime">DateTime</a>)</td> 
    <td><p>Identifies when the comment was published at.</p></td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>Identifies the date and time when the object was last updated.</p></td>
  </tr>
  <tr>
    <td><strong>userContentEdits</strong> (<a href="objects.md#usercontenteditconnection">UserContentEditConnection</a>)</td> 
    <td>
      <p><p>A list of edits to this content.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerDidAuthor</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Did the viewer author this comment.</p></td>
  </tr>
</table>

---

### TopicAuditEntryData

<p>Metadata for an audit entry with a topic.</p> 

#### Implemented by


- [RepoAddTopicAuditEntry](objects.md#repoaddtopicauditentry)
- [RepoRemoveTopicAuditEntry](objects.md#reporemovetopicauditentry) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>topic</strong> (<a href="objects.md#topic">Topic</a>)</td> 
    <td><p>The name of the topic added to the repository</p></td>
  </tr>
  <tr>
    <td><strong>topicName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of the topic added to the repository</p></td>
  </tr>
</table>

---

### Reactable

<p>Represents a subject that can be reacted on.</p> 

#### Implemented by


- [CommitComment](objects.md#commitcomment)
- [Issue](objects.md#issue)
- [IssueComment](objects.md#issuecomment)
- [PullRequest](objects.md#pullrequest)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [TeamDiscussion](objects.md#teamdiscussion)
- [TeamDiscussionComment](objects.md#teamdiscussioncomment) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>databaseId</strong> (<a href="scalars.md#int">Int</a>)</td> 
    <td><p>Identifies the primary key from the database.</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>reactionGroups</strong> (<a href="objects.md#reactiongroup">[ReactionGroup!]</a>)</td> 
    <td><p>A list of reactions grouped by content left on the subject.</p></td>
  </tr>
  <tr>
    <td><strong>reactions</strong> (<a href="objects.md#reactionconnection">ReactionConnection!</a>)</td> 
    <td>
      <p><p>A list of Reactions left on the Issue.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>content (<a href="enums.md#reactioncontent">ReactionContent</a>)</p>
            <p><p>Allows filtering Reactions by emoji.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#reactionorder">ReactionOrder</a>)</p>
            <p><p>Allows specifying the order in which reactions are returned.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>viewerCanReact</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Can user react to this subject</p></td>
  </tr>
</table>

---

### PackageOwner

<p>Represents an owner of a package.</p> 

#### Implemented by


- [Organization](objects.md#organization)
- [Repository](objects.md#repository)
- [User](objects.md#user) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>packages</strong> (<a href="objects.md#packageconnection">PackageConnection!</a>)</td> 
    <td>
      <p><p>A list of packages under the owner.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>names (<a href="scalars.md#string">[String]</a>)</p>
            <p><p>Find packages by their names.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#packageorder">PackageOrder</a>)</p>
            <p><p>Ordering of the returned packages.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>packageType (<a href="enums.md#packagetype">PackageType</a>)</p>
            <p><p>Filter registry package by type.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>repositoryId (<a href="scalars.md#id">ID</a>)</p>
            <p><p>Find packages in a repository by ID.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### Lockable

<p>An object that can be locked.</p> 

#### Implemented by


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>activeLockReason</strong> (<a href="enums.md#lockreason">LockReason</a>)</td> 
    <td><p>Reason that the conversation was locked.</p></td>
  </tr>
  <tr>
    <td><strong>locked</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p><code>true</code> if the object is locked</p></td>
  </tr>
</table>

---

### Subscribable

<p>Entities that can be subscribed to for web and email notifications.</p> 

#### Implemented by


- [Commit](objects.md#commit)
- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest)
- [Repository](objects.md#repository)
- [Team](objects.md#team)
- [TeamDiscussion](objects.md#teamdiscussion) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>viewerCanSubscribe</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Check if the viewer is able to change their subscription status for the repository.</p></td>
  </tr>
  <tr>
    <td><strong>viewerSubscription</strong> (<a href="enums.md#subscriptionstate">SubscriptionState</a>)</td> 
    <td><p>Identifies if the viewer is watching, not watching, or ignoring the subscribable entity.</p></td>
  </tr>
</table>

---

### OauthApplicationAuditEntryData

<p>Metadata for an audit entry with action oauth_application.*</p> 

#### Implemented by


- [OauthApplicationCreateAuditEntry](objects.md#oauthapplicationcreateauditentry)
- [OrgOauthAppAccessApprovedAuditEntry](objects.md#orgoauthappaccessapprovedauditentry)
- [OrgOauthAppAccessDeniedAuditEntry](objects.md#orgoauthappaccessdeniedauditentry)
- [OrgOauthAppAccessRequestedAuditEntry](objects.md#orgoauthappaccessrequestedauditentry) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>oauthApplicationName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of the OAuth Application.</p></td>
  </tr>
  <tr>
    <td><strong>oauthApplicationResourcePath</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP path for the OAuth Application</p></td>
  </tr>
  <tr>
    <td><strong>oauthApplicationUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP URL for the OAuth Application</p></td>
  </tr>
</table>

---

### HovercardContext

<p>An individual line of a hovercard</p> 

#### Implemented by


- [GenericHovercardContext](objects.md#generichovercardcontext)
- [OrganizationTeamsHovercardContext](objects.md#organizationteamshovercardcontext)
- [OrganizationsHovercardContext](objects.md#organizationshovercardcontext)
- [ReviewStatusHovercardContext](objects.md#reviewstatushovercardcontext)
- [ViewerHovercardContext](objects.md#viewerhovercardcontext) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>message</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>A string describing this context</p></td>
  </tr>
  <tr>
    <td><strong>octicon</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>An octicon to accompany this context</p></td>
  </tr>
</table>

---

### Labelable

<p>An object that can have labels assigned to it.</p> 

#### Implemented by


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>labels</strong> (<a href="objects.md#labelconnection">LabelConnection</a>)</td> 
    <td>
      <p><p>A list of labels associated with the object.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>orderBy (<a href="input_objects.md#labelorder">LabelOrder</a>)</p>
            <p><p>Ordering options for labels returned from the connection.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### TeamAuditEntryData

<p>Metadata for an audit entry with action team.*</p> 

#### Implemented by


- [OrgRestoreMemberMembershipTeamAuditEntryData](objects.md#orgrestoremembermembershipteamauditentrydata)
- [TeamAddMemberAuditEntry](objects.md#teamaddmemberauditentry)
- [TeamAddRepositoryAuditEntry](objects.md#teamaddrepositoryauditentry)
- [TeamChangeParentTeamAuditEntry](objects.md#teamchangeparentteamauditentry)
- [TeamRemoveMemberAuditEntry](objects.md#teamremovememberauditentry)
- [TeamRemoveRepositoryAuditEntry](objects.md#teamremoverepositoryauditentry) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>team</strong> (<a href="objects.md#team">Team</a>)</td> 
    <td><p>The team associated with the action</p></td>
  </tr>
  <tr>
    <td><strong>teamName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of the team</p></td>
  </tr>
  <tr>
    <td><strong>teamResourcePath</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP path for this team</p></td>
  </tr>
  <tr>
    <td><strong>teamUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP URL for this team</p></td>
  </tr>
</table>

---

### Node

<p>An object with an ID.</p> 

#### Implemented by


- [AddedToProjectEvent](objects.md#addedtoprojectevent)
- [App](objects.md#app)
- [AssignedEvent](objects.md#assignedevent)
- [AutoMergeDisabledEvent](objects.md#automergedisabledevent)
- [AutoMergeEnabledEvent](objects.md#automergeenabledevent)
- [AutoRebaseEnabledEvent](objects.md#autorebaseenabledevent)
- [AutoSquashEnabledEvent](objects.md#autosquashenabledevent)
- [AutomaticBaseChangeFailedEvent](objects.md#automaticbasechangefailedevent)
- [AutomaticBaseChangeSucceededEvent](objects.md#automaticbasechangesucceededevent)
- [BaseRefChangedEvent](objects.md#baserefchangedevent)
- [BaseRefDeletedEvent](objects.md#baserefdeletedevent)
- [BaseRefForcePushedEvent](objects.md#baserefforcepushedevent)
- [Blob](objects.md#blob)
- [Bot](objects.md#bot)
- [BranchProtectionRule](objects.md#branchprotectionrule)
- [CWE](objects.md#cwe)
- [CheckRun](objects.md#checkrun)
- [CheckSuite](objects.md#checksuite)
- [ClosedEvent](objects.md#closedevent)
- [CodeOfConduct](objects.md#codeofconduct)
- [CommentDeletedEvent](objects.md#commentdeletedevent)
- [Commit](objects.md#commit)
- [CommitComment](objects.md#commitcomment)
- [CommitCommentThread](objects.md#commitcommentthread)
- [ConnectedEvent](objects.md#connectedevent)
- [ConvertToDraftEvent](objects.md#converttodraftevent)
- [ConvertedNoteToIssueEvent](objects.md#convertednotetoissueevent)
- [CrossReferencedEvent](objects.md#crossreferencedevent)
- [DemilestonedEvent](objects.md#demilestonedevent)
- [DependencyGraphManifest](objects.md#dependencygraphmanifest)
- [DeployKey](objects.md#deploykey)
- [DeployedEvent](objects.md#deployedevent)
- [Deployment](objects.md#deployment)
- [DeploymentEnvironmentChangedEvent](objects.md#deploymentenvironmentchangedevent)
- [DeploymentStatus](objects.md#deploymentstatus)
- [DisconnectedEvent](objects.md#disconnectedevent)
- [Enterprise](objects.md#enterprise)
- [EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)
- [EnterpriseIdentityProvider](objects.md#enterpriseidentityprovider)
- [EnterpriseRepositoryInfo](objects.md#enterpriserepositoryinfo)
- [EnterpriseServerInstallation](objects.md#enterpriseserverinstallation)
- [EnterpriseServerUserAccount](objects.md#enterpriseserveruseraccount)
- [EnterpriseServerUserAccountEmail](objects.md#enterpriseserveruseraccountemail)
- [EnterpriseServerUserAccountsUpload](objects.md#enterpriseserveruseraccountsupload)
- [EnterpriseUserAccount](objects.md#enterpriseuseraccount)
- [ExternalIdentity](objects.md#externalidentity)
- [Gist](objects.md#gist)
- [GistComment](objects.md#gistcomment)
- [HeadRefDeletedEvent](objects.md#headrefdeletedevent)
- [HeadRefForcePushedEvent](objects.md#headrefforcepushedevent)
- [HeadRefRestoredEvent](objects.md#headrefrestoredevent)
- [IpAllowListEntry](objects.md#ipallowlistentry)
- [Issue](objects.md#issue)
- [IssueComment](objects.md#issuecomment)
- [Label](objects.md#label)
- [LabeledEvent](objects.md#labeledevent)
- [Language](objects.md#language)
- [License](objects.md#license)
- [LockedEvent](objects.md#lockedevent)
- [Mannequin](objects.md#mannequin)
- [MarkedAsDuplicateEvent](objects.md#markedasduplicateevent)
- [MarketplaceCategory](objects.md#marketplacecategory)
- [MarketplaceListing](objects.md#marketplacelisting)
- [MembersCanDeleteReposClearAuditEntry](objects.md#memberscandeletereposclearauditentry)
- [MembersCanDeleteReposDisableAuditEntry](objects.md#memberscandeletereposdisableauditentry)
- [MembersCanDeleteReposEnableAuditEntry](objects.md#memberscandeletereposenableauditentry)
- [MentionedEvent](objects.md#mentionedevent)
- [MergedEvent](objects.md#mergedevent)
- [Milestone](objects.md#milestone)
- [MilestonedEvent](objects.md#milestonedevent)
- [MovedColumnsInProjectEvent](objects.md#movedcolumnsinprojectevent)
- [OauthApplicationCreateAuditEntry](objects.md#oauthapplicationcreateauditentry)
- [OrgAddBillingManagerAuditEntry](objects.md#orgaddbillingmanagerauditentry)
- [OrgAddMemberAuditEntry](objects.md#orgaddmemberauditentry)
- [OrgBlockUserAuditEntry](objects.md#orgblockuserauditentry)
- [OrgConfigDisableCollaboratorsOnlyAuditEntry](objects.md#orgconfigdisablecollaboratorsonlyauditentry)
- [OrgConfigEnableCollaboratorsOnlyAuditEntry](objects.md#orgconfigenablecollaboratorsonlyauditentry)
- [OrgCreateAuditEntry](objects.md#orgcreateauditentry)
- [OrgDisableOauthAppRestrictionsAuditEntry](objects.md#orgdisableoauthapprestrictionsauditentry)
- [OrgDisableSamlAuditEntry](objects.md#orgdisablesamlauditentry)
- [OrgDisableTwoFactorRequirementAuditEntry](objects.md#orgdisabletwofactorrequirementauditentry)
- [OrgEnableOauthAppRestrictionsAuditEntry](objects.md#orgenableoauthapprestrictionsauditentry)
- [OrgEnableSamlAuditEntry](objects.md#orgenablesamlauditentry)
- [OrgEnableTwoFactorRequirementAuditEntry](objects.md#orgenabletwofactorrequirementauditentry)
- [OrgInviteMemberAuditEntry](objects.md#orginvitememberauditentry)
- [OrgInviteToBusinessAuditEntry](objects.md#orginvitetobusinessauditentry)
- [OrgOauthAppAccessApprovedAuditEntry](objects.md#orgoauthappaccessapprovedauditentry)
- [OrgOauthAppAccessDeniedAuditEntry](objects.md#orgoauthappaccessdeniedauditentry)
- [OrgOauthAppAccessRequestedAuditEntry](objects.md#orgoauthappaccessrequestedauditentry)
- [OrgRemoveBillingManagerAuditEntry](objects.md#orgremovebillingmanagerauditentry)
- [OrgRemoveMemberAuditEntry](objects.md#orgremovememberauditentry)
- [OrgRemoveOutsideCollaboratorAuditEntry](objects.md#orgremoveoutsidecollaboratorauditentry)
- [OrgRestoreMemberAuditEntry](objects.md#orgrestorememberauditentry)
- [OrgUnblockUserAuditEntry](objects.md#orgunblockuserauditentry)
- [OrgUpdateDefaultRepositoryPermissionAuditEntry](objects.md#orgupdatedefaultrepositorypermissionauditentry)
- [OrgUpdateMemberAuditEntry](objects.md#orgupdatememberauditentry)
- [OrgUpdateMemberRepositoryCreationPermissionAuditEntry](objects.md#orgupdatememberrepositorycreationpermissionauditentry)
- [OrgUpdateMemberRepositoryInvitationPermissionAuditEntry](objects.md#orgupdatememberrepositoryinvitationpermissionauditentry)
- [Organization](objects.md#organization)
- [OrganizationIdentityProvider](objects.md#organizationidentityprovider)
- [OrganizationInvitation](objects.md#organizationinvitation)
- [Package](objects.md#package)
- [PackageFile](objects.md#packagefile)
- [PackageTag](objects.md#packagetag)
- [PackageVersion](objects.md#packageversion)
- [PinnedEvent](objects.md#pinnedevent)
- [PinnedIssue](objects.md#pinnedissue)
- [PrivateRepositoryForkingDisableAuditEntry](objects.md#privaterepositoryforkingdisableauditentry)
- [PrivateRepositoryForkingEnableAuditEntry](objects.md#privaterepositoryforkingenableauditentry)
- [Project](objects.md#project)
- [ProjectCard](objects.md#projectcard)
- [ProjectColumn](objects.md#projectcolumn)
- [PublicKey](objects.md#publickey)
- [PullRequest](objects.md#pullrequest)
- [PullRequestCommit](objects.md#pullrequestcommit)
- [PullRequestCommitCommentThread](objects.md#pullrequestcommitcommentthread)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [PullRequestReviewThread](objects.md#pullrequestreviewthread)
- [Push](objects.md#push)
- [PushAllowance](objects.md#pushallowance)
- [Reaction](objects.md#reaction)
- [ReadyForReviewEvent](objects.md#readyforreviewevent)
- [Ref](objects.md#ref)
- [ReferencedEvent](objects.md#referencedevent)
- [Release](objects.md#release)
- [ReleaseAsset](objects.md#releaseasset)
- [RemovedFromProjectEvent](objects.md#removedfromprojectevent)
- [RenamedTitleEvent](objects.md#renamedtitleevent)
- [ReopenedEvent](objects.md#reopenedevent)
- [RepoAccessAuditEntry](objects.md#repoaccessauditentry)
- [RepoAddMemberAuditEntry](objects.md#repoaddmemberauditentry)
- [RepoAddTopicAuditEntry](objects.md#repoaddtopicauditentry)
- [RepoArchivedAuditEntry](objects.md#repoarchivedauditentry)
- [RepoChangeMergeSettingAuditEntry](objects.md#repochangemergesettingauditentry)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](objects.md#repoconfigdisableanonymousgitaccessauditentry)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](objects.md#repoconfigdisablecollaboratorsonlyauditentry)
- [RepoConfigDisableContributorsOnlyAuditEntry](objects.md#repoconfigdisablecontributorsonlyauditentry)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](objects.md#repoconfigdisablesockpuppetdisallowedauditentry)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](objects.md#repoconfigenableanonymousgitaccessauditentry)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](objects.md#repoconfigenablecollaboratorsonlyauditentry)
- [RepoConfigEnableContributorsOnlyAuditEntry](objects.md#repoconfigenablecontributorsonlyauditentry)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](objects.md#repoconfigenablesockpuppetdisallowedauditentry)
- [RepoConfigLockAnonymousGitAccessAuditEntry](objects.md#repoconfiglockanonymousgitaccessauditentry)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](objects.md#repoconfigunlockanonymousgitaccessauditentry)
- [RepoCreateAuditEntry](objects.md#repocreateauditentry)
- [RepoDestroyAuditEntry](objects.md#repodestroyauditentry)
- [RepoRemoveMemberAuditEntry](objects.md#reporemovememberauditentry)
- [RepoRemoveTopicAuditEntry](objects.md#reporemovetopicauditentry)
- [Repository](objects.md#repository)
- [RepositoryInvitation](objects.md#repositoryinvitation)
- [RepositoryTopic](objects.md#repositorytopic)
- [RepositoryVisibilityChangeDisableAuditEntry](objects.md#repositoryvisibilitychangedisableauditentry)
- [RepositoryVisibilityChangeEnableAuditEntry](objects.md#repositoryvisibilitychangeenableauditentry)
- [RepositoryVulnerabilityAlert](objects.md#repositoryvulnerabilityalert)
- [ReviewDismissalAllowance](objects.md#reviewdismissalallowance)
- [ReviewDismissedEvent](objects.md#reviewdismissedevent)
- [ReviewRequest](objects.md#reviewrequest)
- [ReviewRequestRemovedEvent](objects.md#reviewrequestremovedevent)
- [ReviewRequestedEvent](objects.md#reviewrequestedevent)
- [SavedReply](objects.md#savedreply)
- [SecurityAdvisory](objects.md#securityadvisory)
- [SponsorsListing](objects.md#sponsorslisting)
- [SponsorsTier](objects.md#sponsorstier)
- [Sponsorship](objects.md#sponsorship)
- [Status](objects.md#status)
- [StatusCheckRollup](objects.md#statuscheckrollup)
- [StatusContext](objects.md#statuscontext)
- [SubscribedEvent](objects.md#subscribedevent)
- [Tag](objects.md#tag)
- [Team](objects.md#team)
- [TeamAddMemberAuditEntry](objects.md#teamaddmemberauditentry)
- [TeamAddRepositoryAuditEntry](objects.md#teamaddrepositoryauditentry)
- [TeamChangeParentTeamAuditEntry](objects.md#teamchangeparentteamauditentry)
- [TeamDiscussion](objects.md#teamdiscussion)
- [TeamDiscussionComment](objects.md#teamdiscussioncomment)
- [TeamRemoveMemberAuditEntry](objects.md#teamremovememberauditentry)
- [TeamRemoveRepositoryAuditEntry](objects.md#teamremoverepositoryauditentry)
- [Topic](objects.md#topic)
- [TransferredEvent](objects.md#transferredevent)
- [Tree](objects.md#tree)
- [UnassignedEvent](objects.md#unassignedevent)
- [UnlabeledEvent](objects.md#unlabeledevent)
- [UnlockedEvent](objects.md#unlockedevent)
- [UnmarkedAsDuplicateEvent](objects.md#unmarkedasduplicateevent)
- [UnpinnedEvent](objects.md#unpinnedevent)
- [UnsubscribedEvent](objects.md#unsubscribedevent)
- [User](objects.md#user)
- [UserBlockedEvent](objects.md#userblockedevent)
- [UserContentEdit](objects.md#usercontentedit)
- [UserStatus](objects.md#userstatus)
- [VerifiableDomain](objects.md#verifiabledomain) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td><p>ID of the object.</p></td>
  </tr>
</table>

---

### Updatable

<p>Entities that can be updated.</p> 

#### Implemented by


- [CommitComment](objects.md#commitcomment)
- [GistComment](objects.md#gistcomment)
- [Issue](objects.md#issue)
- [IssueComment](objects.md#issuecomment)
- [Project](objects.md#project)
- [PullRequest](objects.md#pullrequest)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [TeamDiscussion](objects.md#teamdiscussion)
- [TeamDiscussionComment](objects.md#teamdiscussioncomment) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>viewerCanUpdate</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Check if the current viewer can update this object.</p></td>
  </tr>
</table>

---

### Assignable

<p>An object that can have users assigned to it.</p> 

#### Implemented by


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>assignees</strong> (<a href="objects.md#userconnection">UserConnection!</a>)</td> 
    <td>
      <p><p>A list of Users assigned to this object.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come after the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>before (<a href="scalars.md#string">String</a>)</p>
            <p><p>Returns the elements in the list that come before the specified cursor.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the first <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>last (<a href="scalars.md#int">Int</a>)</p>
            <p><p>Returns the last <em>n</em> elements from the list.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### EnterpriseAuditEntryData

<p>Metadata for an audit entry containing enterprise account information.</p> 

#### Implemented by


- [MembersCanDeleteReposClearAuditEntry](objects.md#memberscandeletereposclearauditentry)
- [MembersCanDeleteReposDisableAuditEntry](objects.md#memberscandeletereposdisableauditentry)
- [MembersCanDeleteReposEnableAuditEntry](objects.md#memberscandeletereposenableauditentry)
- [OrgInviteToBusinessAuditEntry](objects.md#orginvitetobusinessauditentry)
- [PrivateRepositoryForkingDisableAuditEntry](objects.md#privaterepositoryforkingdisableauditentry)
- [PrivateRepositoryForkingEnableAuditEntry](objects.md#privaterepositoryforkingenableauditentry)
- [RepositoryVisibilityChangeDisableAuditEntry](objects.md#repositoryvisibilitychangedisableauditentry)
- [RepositoryVisibilityChangeEnableAuditEntry](objects.md#repositoryvisibilitychangeenableauditentry) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>enterpriseResourcePath</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP path for this enterprise.</p></td>
  </tr>
  <tr>
    <td><strong>enterpriseSlug</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The slug of the enterprise.</p></td>
  </tr>
  <tr>
    <td><strong>enterpriseUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP URL for this enterprise.</p></td>
  </tr>
</table>

---

### GitObject

<p>Represents a Git object.</p> 

#### Implemented by


- [Blob](objects.md#blob)
- [Commit](objects.md#commit)
- [Tag](objects.md#tag)
- [Tree](objects.md#tree) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>abbreviatedOid</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>An abbreviated version of the Git object ID</p></td>
  </tr>
  <tr>
    <td><strong>commitResourcePath</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP path for this Git object</p></td>
  </tr>
  <tr>
    <td><strong>commitUrl</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP URL for this Git object</p></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>oid</strong> (<a href="scalars.md#gitobjectid">GitObjectID!</a>)</td> 
    <td><p>The Git object ID</p></td>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="objects.md#repository">Repository!</a>)</td> 
    <td><p>The Repository the Git object belongs to</p></td>
  </tr>
</table>

---

### RepositoryNode

<p>Represents a object that belongs to a repository.</p> 

#### Implemented by


- [CommitComment](objects.md#commitcomment)
- [CommitCommentThread](objects.md#commitcommentthread)
- [Issue](objects.md#issue)
- [IssueComment](objects.md#issuecomment)
- [PullRequest](objects.md#pullrequest)
- [PullRequestCommitCommentThread](objects.md#pullrequestcommitcommentthread)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [RepositoryVulnerabilityAlert](objects.md#repositoryvulnerabilityalert) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>repository</strong> (<a href="objects.md#repository">Repository!</a>)</td> 
    <td><p>The repository associated with this node.</p></td>
  </tr>
</table>

---

### Closable

<p>An object that can be closed</p> 

#### Implemented by


- [Issue](objects.md#issue)
- [Milestone](objects.md#milestone)
- [Project](objects.md#project)
- [PullRequest](objects.md#pullrequest) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>closed</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p><code>true</code> if the object is closed (definition of closed may depend on type)</p></td>
  </tr>
  <tr>
    <td><strong>closedAt</strong> (<a href="scalars.md#datetime">DateTime</a>)</td> 
    <td><p>Identifies the date and time when the object was closed.</p></td>
  </tr>
</table>

---

### UniformResourceLocatable

<p>Represents a type that can be retrieved by a URL.</p> 

#### Implemented by


- [Bot](objects.md#bot)
- [CheckRun](objects.md#checkrun)
- [ClosedEvent](objects.md#closedevent)
- [Commit](objects.md#commit)
- [ConvertToDraftEvent](objects.md#converttodraftevent)
- [CrossReferencedEvent](objects.md#crossreferencedevent)
- [Gist](objects.md#gist)
- [Issue](objects.md#issue)
- [Mannequin](objects.md#mannequin)
- [MergedEvent](objects.md#mergedevent)
- [Milestone](objects.md#milestone)
- [Organization](objects.md#organization)
- [PullRequest](objects.md#pullrequest)
- [PullRequestCommit](objects.md#pullrequestcommit)
- [ReadyForReviewEvent](objects.md#readyforreviewevent)
- [Release](objects.md#release)
- [Repository](objects.md#repository)
- [RepositoryTopic](objects.md#repositorytopic)
- [ReviewDismissedEvent](objects.md#reviewdismissedevent)
- [TeamDiscussion](objects.md#teamdiscussion)
- [TeamDiscussionComment](objects.md#teamdiscussioncomment)
- [User](objects.md#user) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTML path to this resource.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The URL to this resource.</p></td>
  </tr>
</table>

---

### OrganizationAuditEntryData

<p>Metadata for an audit entry with action org.*</p> 

#### Implemented by


- [MembersCanDeleteReposClearAuditEntry](objects.md#memberscandeletereposclearauditentry)
- [MembersCanDeleteReposDisableAuditEntry](objects.md#memberscandeletereposdisableauditentry)
- [MembersCanDeleteReposEnableAuditEntry](objects.md#memberscandeletereposenableauditentry)
- [OauthApplicationCreateAuditEntry](objects.md#oauthapplicationcreateauditentry)
- [OrgAddBillingManagerAuditEntry](objects.md#orgaddbillingmanagerauditentry)
- [OrgAddMemberAuditEntry](objects.md#orgaddmemberauditentry)
- [OrgBlockUserAuditEntry](objects.md#orgblockuserauditentry)
- [OrgConfigDisableCollaboratorsOnlyAuditEntry](objects.md#orgconfigdisablecollaboratorsonlyauditentry)
- [OrgConfigEnableCollaboratorsOnlyAuditEntry](objects.md#orgconfigenablecollaboratorsonlyauditentry)
- [OrgCreateAuditEntry](objects.md#orgcreateauditentry)
- [OrgDisableOauthAppRestrictionsAuditEntry](objects.md#orgdisableoauthapprestrictionsauditentry)
- [OrgDisableSamlAuditEntry](objects.md#orgdisablesamlauditentry)
- [OrgDisableTwoFactorRequirementAuditEntry](objects.md#orgdisabletwofactorrequirementauditentry)
- [OrgEnableOauthAppRestrictionsAuditEntry](objects.md#orgenableoauthapprestrictionsauditentry)
- [OrgEnableSamlAuditEntry](objects.md#orgenablesamlauditentry)
- [OrgEnableTwoFactorRequirementAuditEntry](objects.md#orgenabletwofactorrequirementauditentry)
- [OrgInviteMemberAuditEntry](objects.md#orginvitememberauditentry)
- [OrgInviteToBusinessAuditEntry](objects.md#orginvitetobusinessauditentry)
- [OrgOauthAppAccessApprovedAuditEntry](objects.md#orgoauthappaccessapprovedauditentry)
- [OrgOauthAppAccessDeniedAuditEntry](objects.md#orgoauthappaccessdeniedauditentry)
- [OrgOauthAppAccessRequestedAuditEntry](objects.md#orgoauthappaccessrequestedauditentry)
- [OrgRemoveBillingManagerAuditEntry](objects.md#orgremovebillingmanagerauditentry)
- [OrgRemoveMemberAuditEntry](objects.md#orgremovememberauditentry)
- [OrgRemoveOutsideCollaboratorAuditEntry](objects.md#orgremoveoutsidecollaboratorauditentry)
- [OrgRestoreMemberAuditEntry](objects.md#orgrestorememberauditentry)
- [OrgRestoreMemberMembershipOrganizationAuditEntryData](objects.md#orgrestoremembermembershiporganizationauditentrydata)
- [OrgUnblockUserAuditEntry](objects.md#orgunblockuserauditentry)
- [OrgUpdateDefaultRepositoryPermissionAuditEntry](objects.md#orgupdatedefaultrepositorypermissionauditentry)
- [OrgUpdateMemberAuditEntry](objects.md#orgupdatememberauditentry)
- [OrgUpdateMemberRepositoryCreationPermissionAuditEntry](objects.md#orgupdatememberrepositorycreationpermissionauditentry)
- [OrgUpdateMemberRepositoryInvitationPermissionAuditEntry](objects.md#orgupdatememberrepositoryinvitationpermissionauditentry)
- [PrivateRepositoryForkingDisableAuditEntry](objects.md#privaterepositoryforkingdisableauditentry)
- [PrivateRepositoryForkingEnableAuditEntry](objects.md#privaterepositoryforkingenableauditentry)
- [RepoAccessAuditEntry](objects.md#repoaccessauditentry)
- [RepoAddMemberAuditEntry](objects.md#repoaddmemberauditentry)
- [RepoAddTopicAuditEntry](objects.md#repoaddtopicauditentry)
- [RepoArchivedAuditEntry](objects.md#repoarchivedauditentry)
- [RepoChangeMergeSettingAuditEntry](objects.md#repochangemergesettingauditentry)
- [RepoConfigDisableAnonymousGitAccessAuditEntry](objects.md#repoconfigdisableanonymousgitaccessauditentry)
- [RepoConfigDisableCollaboratorsOnlyAuditEntry](objects.md#repoconfigdisablecollaboratorsonlyauditentry)
- [RepoConfigDisableContributorsOnlyAuditEntry](objects.md#repoconfigdisablecontributorsonlyauditentry)
- [RepoConfigDisableSockpuppetDisallowedAuditEntry](objects.md#repoconfigdisablesockpuppetdisallowedauditentry)
- [RepoConfigEnableAnonymousGitAccessAuditEntry](objects.md#repoconfigenableanonymousgitaccessauditentry)
- [RepoConfigEnableCollaboratorsOnlyAuditEntry](objects.md#repoconfigenablecollaboratorsonlyauditentry)
- [RepoConfigEnableContributorsOnlyAuditEntry](objects.md#repoconfigenablecontributorsonlyauditentry)
- [RepoConfigEnableSockpuppetDisallowedAuditEntry](objects.md#repoconfigenablesockpuppetdisallowedauditentry)
- [RepoConfigLockAnonymousGitAccessAuditEntry](objects.md#repoconfiglockanonymousgitaccessauditentry)
- [RepoConfigUnlockAnonymousGitAccessAuditEntry](objects.md#repoconfigunlockanonymousgitaccessauditentry)
- [RepoCreateAuditEntry](objects.md#repocreateauditentry)
- [RepoDestroyAuditEntry](objects.md#repodestroyauditentry)
- [RepoRemoveMemberAuditEntry](objects.md#reporemovememberauditentry)
- [RepoRemoveTopicAuditEntry](objects.md#reporemovetopicauditentry)
- [RepositoryVisibilityChangeDisableAuditEntry](objects.md#repositoryvisibilitychangedisableauditentry)
- [RepositoryVisibilityChangeEnableAuditEntry](objects.md#repositoryvisibilitychangeenableauditentry)
- [TeamAddMemberAuditEntry](objects.md#teamaddmemberauditentry)
- [TeamAddRepositoryAuditEntry](objects.md#teamaddrepositoryauditentry)
- [TeamChangeParentTeamAuditEntry](objects.md#teamchangeparentteamauditentry)
- [TeamRemoveMemberAuditEntry](objects.md#teamremovememberauditentry)
- [TeamRemoveRepositoryAuditEntry](objects.md#teamremoverepositoryauditentry) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>organization</strong> (<a href="objects.md#organization">Organization</a>)</td> 
    <td><p>The Organization associated with the Audit Entry.</p></td>
  </tr>
  <tr>
    <td><strong>organizationName</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The name of the Organization.</p></td>
  </tr>
  <tr>
    <td><strong>organizationResourcePath</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP path for the organization</p></td>
  </tr>
  <tr>
    <td><strong>organizationUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The HTTP URL for the organization</p></td>
  </tr>
</table>

---

### GitSignature

<p>Information about a signature (GPG or S/MIME) on a Commit or Tag.</p> 

#### Implemented by


- [GpgSignature](objects.md#gpgsignature)
- [SmimeSignature](objects.md#smimesignature)
- [UnknownSignature](objects.md#unknownsignature) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Email used to sign this object.</p></td>
  </tr>
  <tr>
    <td><strong>isValid</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>True if the signature is valid and verified by GitHub.</p></td>
  </tr>
  <tr>
    <td><strong>payload</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>Payload for GPG signing object. Raw ODB object without the signature header.</p></td>
  </tr>
  <tr>
    <td><strong>signature</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>ASCII-armored signature header from object.</p></td>
  </tr>
  <tr>
    <td><strong>signer</strong> (<a href="objects.md#user">User</a>)</td> 
    <td><p>GitHub user corresponding to the email signing this commit.</p></td>
  </tr>
  <tr>
    <td><strong>state</strong> (<a href="enums.md#gitsignaturestate">GitSignatureState!</a>)</td> 
    <td><p>The state of this signature. <code>VALID</code> if signature is valid and verified by
GitHub, otherwise represents reason why signature is considered invalid.</p></td>
  </tr>
  <tr>
    <td><strong>wasSignedByGitHub</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>True if the signature was made with GitHub&rsquo;s signing key.</p></td>
  </tr>
</table>

---

### RepositoryInfo

<p>A subset of repository info.</p> 

#### Implemented by


- [Repository](objects.md#repository) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>Identifies the date and time when the object was created.</p></td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td><p>The description of the repository.</p></td>
  </tr>
  <tr>
    <td><strong>descriptionHTML</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td><p>The description of the repository rendered to HTML.</p></td>
  </tr>
  <tr>
    <td><strong>forkCount</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td><p>Returns how many forks there are of this repository in the whole network.</p></td>
  </tr>
  <tr>
    <td><strong>hasIssuesEnabled</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if the repository has issues feature enabled.</p></td>
  </tr>
  <tr>
    <td><strong>hasProjectsEnabled</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if the repository has the Projects feature enabled.</p></td>
  </tr>
  <tr>
    <td><strong>hasWikiEnabled</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if the repository has wiki feature enabled.</p></td>
  </tr>
  <tr>
    <td><strong>homepageUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The repository&rsquo;s URL.</p></td>
  </tr>
  <tr>
    <td><strong>isArchived</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if the repository is unmaintained.</p></td>
  </tr>
  <tr>
    <td><strong>isFork</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Identifies if the repository is a fork.</p></td>
  </tr>
  <tr>
    <td><strong>isInOrganization</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if a repository is either owned by an organization, or is a private fork of an organization repository.</p></td>
  </tr>
  <tr>
    <td><strong>isLocked</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Indicates if the repository has been locked or not.</p></td>
  </tr>
  <tr>
    <td><strong>isMirror</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Identifies if the repository is a mirror.</p></td>
  </tr>
  <tr>
    <td><strong>isPrivate</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Identifies if the repository is private.</p></td>
  </tr>
  <tr>
    <td><strong>isTemplate</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Identifies if the repository is a template that can be used to generate new repositories.</p></td>
  </tr>
  <tr>
    <td><strong>licenseInfo</strong> (<a href="objects.md#license">License</a>)</td> 
    <td><p>The license associated with the repository</p></td>
  </tr>
  <tr>
    <td><strong>lockReason</strong> (<a href="enums.md#repositorylockreason">RepositoryLockReason</a>)</td> 
    <td><p>The reason the repository has been locked.</p></td>
  </tr>
  <tr>
    <td><strong>mirrorUrl</strong> (<a href="scalars.md#uri">URI</a>)</td> 
    <td><p>The repository&rsquo;s original mirror URL.</p></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The name of the repository.</p></td>
  </tr>
  <tr>
    <td><strong>nameWithOwner</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td><p>The repository&rsquo;s name with owner.</p></td>
  </tr>
  <tr>
    <td><strong>openGraphImageUrl</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The image used to represent this repository in Open Graph data.</p></td>
  </tr>
  <tr>
    <td><strong>owner</strong> (<a href="interfaces.md#repositoryowner">RepositoryOwner!</a>)</td> 
    <td><p>The User owner of the repository.</p></td>
  </tr>
  <tr>
    <td><strong>pushedAt</strong> (<a href="scalars.md#datetime">DateTime</a>)</td> 
    <td><p>Identifies when the repository was last pushed to.</p></td>
  </tr>
  <tr>
    <td><strong>resourcePath</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP path for this repository</p></td>
  </tr>
  <tr>
    <td><strong>shortDescriptionHTML</strong> (<a href="scalars.md#html">HTML!</a>)</td> 
    <td>
      <p><p>A description of the repository, rendered to HTML without any links in it.</p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>limit (<a href="scalars.md#int">Int</a>)</p>
            <p><p>How many characters to return.</p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
  <tr>
    <td><strong>updatedAt</strong> (<a href="scalars.md#datetime">DateTime!</a>)</td> 
    <td><p>Identifies the date and time when the object was last updated.</p></td>
  </tr>
  <tr>
    <td><strong>url</strong> (<a href="scalars.md#uri">URI!</a>)</td> 
    <td><p>The HTTP URL for this repository</p></td>
  </tr>
  <tr>
    <td><strong>usesCustomOpenGraphImage</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td><p>Whether this repository has a custom image to use with Open Graph as opposed to being represented by the owner&rsquo;s avatar.</p></td>
  </tr>
</table>

---