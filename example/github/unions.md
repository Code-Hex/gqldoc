# Unions

### About unions

A [union](https://graphql.github.io/graphql-spec/June2018/#sec-Unions) is a type of object representing many objects.

### Assignee

<p>Types that can be assigned to issues.</p>

### Possible types


- [Bot](objects.md#bot)
- [Mannequin](objects.md#mannequin)
- [Organization](objects.md#organization)
- [User](objects.md#user)

---

### AuditEntryActor

<p>Types that can initiate an audit log event.</p>

### Possible types


- [Bot](objects.md#bot)
- [Organization](objects.md#organization)
- [User](objects.md#user)

---

### BranchActorAllowanceActor

<p>Types which can be actors for <code>BranchActorAllowance</code> objects.</p>

### Possible types


- [Team](objects.md#team)
- [User](objects.md#user)

---

### Closer

<p>The object which triggered a <code>ClosedEvent</code>.</p>

### Possible types


- [Commit](objects.md#commit)
- [PullRequest](objects.md#pullrequest)

---

### CreatedIssueOrRestrictedContribution

<p>Represents either a issue the viewer can access or a restricted contribution.</p>

### Possible types


- [CreatedIssueContribution](objects.md#createdissuecontribution)
- [RestrictedContribution](objects.md#restrictedcontribution)

---

### CreatedPullRequestOrRestrictedContribution

<p>Represents either a pull request the viewer can access or a restricted contribution.</p>

### Possible types


- [CreatedPullRequestContribution](objects.md#createdpullrequestcontribution)
- [RestrictedContribution](objects.md#restrictedcontribution)

---

### CreatedRepositoryOrRestrictedContribution

<p>Represents either a repository the viewer can access or a restricted contribution.</p>

### Possible types


- [CreatedRepositoryContribution](objects.md#createdrepositorycontribution)
- [RestrictedContribution](objects.md#restrictedcontribution)

---

### DeploymentReviewer

<p>Users and teams.</p>

### Possible types


- [Team](objects.md#team)
- [User](objects.md#user)

---

### EnterpriseMember

<p>An object that is a member of an enterprise.</p>

### Possible types


- [EnterpriseUserAccount](objects.md#enterpriseuseraccount)
- [User](objects.md#user)

---

### IpAllowListOwner

<p>Types that can own an IP allow list.</p>

### Possible types


- [App](objects.md#app)
- [Enterprise](objects.md#enterprise)
- [Organization](objects.md#organization)

---

### IssueOrPullRequest

<p>Used for return value of Repository.issueOrPullRequest.</p>

### Possible types


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest)

---

### IssueTimelineItem

<p>An item in an issue timeline</p>

### Possible types


- [AssignedEvent](objects.md#assignedevent)
- [ClosedEvent](objects.md#closedevent)
- [Commit](objects.md#commit)
- [CrossReferencedEvent](objects.md#crossreferencedevent)
- [DemilestonedEvent](objects.md#demilestonedevent)
- [IssueComment](objects.md#issuecomment)
- [LabeledEvent](objects.md#labeledevent)
- [LockedEvent](objects.md#lockedevent)
- [MilestonedEvent](objects.md#milestonedevent)
- [ReferencedEvent](objects.md#referencedevent)
- [RenamedTitleEvent](objects.md#renamedtitleevent)
- [ReopenedEvent](objects.md#reopenedevent)
- [SubscribedEvent](objects.md#subscribedevent)
- [TransferredEvent](objects.md#transferredevent)
- [UnassignedEvent](objects.md#unassignedevent)
- [UnlabeledEvent](objects.md#unlabeledevent)
- [UnlockedEvent](objects.md#unlockedevent)
- [UnsubscribedEvent](objects.md#unsubscribedevent)
- [UserBlockedEvent](objects.md#userblockedevent)

---

### IssueTimelineItems

<p>An item in an issue timeline</p>

### Possible types


- [AddedToProjectEvent](objects.md#addedtoprojectevent)
- [AssignedEvent](objects.md#assignedevent)
- [ClosedEvent](objects.md#closedevent)
- [CommentDeletedEvent](objects.md#commentdeletedevent)
- [ConnectedEvent](objects.md#connectedevent)
- [ConvertedNoteToIssueEvent](objects.md#convertednotetoissueevent)
- [ConvertedToDiscussionEvent](objects.md#convertedtodiscussionevent)
- [CrossReferencedEvent](objects.md#crossreferencedevent)
- [DemilestonedEvent](objects.md#demilestonedevent)
- [DisconnectedEvent](objects.md#disconnectedevent)
- [IssueComment](objects.md#issuecomment)
- [LabeledEvent](objects.md#labeledevent)
- [LockedEvent](objects.md#lockedevent)
- [MarkedAsDuplicateEvent](objects.md#markedasduplicateevent)
- [MentionedEvent](objects.md#mentionedevent)
- [MilestonedEvent](objects.md#milestonedevent)
- [MovedColumnsInProjectEvent](objects.md#movedcolumnsinprojectevent)
- [PinnedEvent](objects.md#pinnedevent)
- [ReferencedEvent](objects.md#referencedevent)
- [RemovedFromProjectEvent](objects.md#removedfromprojectevent)
- [RenamedTitleEvent](objects.md#renamedtitleevent)
- [ReopenedEvent](objects.md#reopenedevent)
- [SubscribedEvent](objects.md#subscribedevent)
- [TransferredEvent](objects.md#transferredevent)
- [UnassignedEvent](objects.md#unassignedevent)
- [UnlabeledEvent](objects.md#unlabeledevent)
- [UnlockedEvent](objects.md#unlockedevent)
- [UnmarkedAsDuplicateEvent](objects.md#unmarkedasduplicateevent)
- [UnpinnedEvent](objects.md#unpinnedevent)
- [UnsubscribedEvent](objects.md#unsubscribedevent)
- [UserBlockedEvent](objects.md#userblockedevent)

---

### MilestoneItem

<p>Types that can be inside a Milestone.</p>

### Possible types


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest)

---

### OrgRestoreMemberAuditEntryMembership

<p>Types of memberships that can be restored for an Organization member.</p>

### Possible types


- [OrgRestoreMemberMembershipOrganizationAuditEntryData](objects.md#orgrestoremembermembershiporganizationauditentrydata)
- [OrgRestoreMemberMembershipRepositoryAuditEntryData](objects.md#orgrestoremembermembershiprepositoryauditentrydata)
- [OrgRestoreMemberMembershipTeamAuditEntryData](objects.md#orgrestoremembermembershipteamauditentrydata)

---

### OrganizationAuditEntry

<p>An audit entry in an organization audit log.</p>

### Possible types


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

---

### PermissionGranter

<p>Types that can grant permissions on a repository to a user</p>

### Possible types


- [Organization](objects.md#organization)
- [Repository](objects.md#repository)
- [Team](objects.md#team)

---

### PinnableItem

<p>Types that can be pinned to a profile page.</p>

### Possible types


- [Gist](objects.md#gist)
- [Repository](objects.md#repository)

---

### ProjectCardItem

<p>Types that can be inside Project Cards.</p>

### Possible types


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest)

---

### ProjectNextItemContent

<p>Types that can be inside Project Items.</p>

### Possible types


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest)

---

### PullRequestTimelineItem

<p>An item in a pull request timeline</p>

### Possible types


- [AssignedEvent](objects.md#assignedevent)
- [BaseRefDeletedEvent](objects.md#baserefdeletedevent)
- [BaseRefForcePushedEvent](objects.md#baserefforcepushedevent)
- [ClosedEvent](objects.md#closedevent)
- [Commit](objects.md#commit)
- [CommitCommentThread](objects.md#commitcommentthread)
- [CrossReferencedEvent](objects.md#crossreferencedevent)
- [DemilestonedEvent](objects.md#demilestonedevent)
- [DeployedEvent](objects.md#deployedevent)
- [DeploymentEnvironmentChangedEvent](objects.md#deploymentenvironmentchangedevent)
- [HeadRefDeletedEvent](objects.md#headrefdeletedevent)
- [HeadRefForcePushedEvent](objects.md#headrefforcepushedevent)
- [HeadRefRestoredEvent](objects.md#headrefrestoredevent)
- [IssueComment](objects.md#issuecomment)
- [LabeledEvent](objects.md#labeledevent)
- [LockedEvent](objects.md#lockedevent)
- [MergedEvent](objects.md#mergedevent)
- [MilestonedEvent](objects.md#milestonedevent)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewComment](objects.md#pullrequestreviewcomment)
- [PullRequestReviewThread](objects.md#pullrequestreviewthread)
- [ReferencedEvent](objects.md#referencedevent)
- [RenamedTitleEvent](objects.md#renamedtitleevent)
- [ReopenedEvent](objects.md#reopenedevent)
- [ReviewDismissedEvent](objects.md#reviewdismissedevent)
- [ReviewRequestRemovedEvent](objects.md#reviewrequestremovedevent)
- [ReviewRequestedEvent](objects.md#reviewrequestedevent)
- [SubscribedEvent](objects.md#subscribedevent)
- [UnassignedEvent](objects.md#unassignedevent)
- [UnlabeledEvent](objects.md#unlabeledevent)
- [UnlockedEvent](objects.md#unlockedevent)
- [UnsubscribedEvent](objects.md#unsubscribedevent)
- [UserBlockedEvent](objects.md#userblockedevent)

---

### PullRequestTimelineItems

<p>An item in a pull request timeline</p>

### Possible types


- [AddedToProjectEvent](objects.md#addedtoprojectevent)
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
- [ClosedEvent](objects.md#closedevent)
- [CommentDeletedEvent](objects.md#commentdeletedevent)
- [ConnectedEvent](objects.md#connectedevent)
- [ConvertToDraftEvent](objects.md#converttodraftevent)
- [ConvertedNoteToIssueEvent](objects.md#convertednotetoissueevent)
- [ConvertedToDiscussionEvent](objects.md#convertedtodiscussionevent)
- [CrossReferencedEvent](objects.md#crossreferencedevent)
- [DemilestonedEvent](objects.md#demilestonedevent)
- [DeployedEvent](objects.md#deployedevent)
- [DeploymentEnvironmentChangedEvent](objects.md#deploymentenvironmentchangedevent)
- [DisconnectedEvent](objects.md#disconnectedevent)
- [HeadRefDeletedEvent](objects.md#headrefdeletedevent)
- [HeadRefForcePushedEvent](objects.md#headrefforcepushedevent)
- [HeadRefRestoredEvent](objects.md#headrefrestoredevent)
- [IssueComment](objects.md#issuecomment)
- [LabeledEvent](objects.md#labeledevent)
- [LockedEvent](objects.md#lockedevent)
- [MarkedAsDuplicateEvent](objects.md#markedasduplicateevent)
- [MentionedEvent](objects.md#mentionedevent)
- [MergedEvent](objects.md#mergedevent)
- [MilestonedEvent](objects.md#milestonedevent)
- [MovedColumnsInProjectEvent](objects.md#movedcolumnsinprojectevent)
- [PinnedEvent](objects.md#pinnedevent)
- [PullRequestCommit](objects.md#pullrequestcommit)
- [PullRequestCommitCommentThread](objects.md#pullrequestcommitcommentthread)
- [PullRequestReview](objects.md#pullrequestreview)
- [PullRequestReviewThread](objects.md#pullrequestreviewthread)
- [PullRequestRevisionMarker](objects.md#pullrequestrevisionmarker)
- [ReadyForReviewEvent](objects.md#readyforreviewevent)
- [ReferencedEvent](objects.md#referencedevent)
- [RemovedFromProjectEvent](objects.md#removedfromprojectevent)
- [RenamedTitleEvent](objects.md#renamedtitleevent)
- [ReopenedEvent](objects.md#reopenedevent)
- [ReviewDismissedEvent](objects.md#reviewdismissedevent)
- [ReviewRequestRemovedEvent](objects.md#reviewrequestremovedevent)
- [ReviewRequestedEvent](objects.md#reviewrequestedevent)
- [SubscribedEvent](objects.md#subscribedevent)
- [TransferredEvent](objects.md#transferredevent)
- [UnassignedEvent](objects.md#unassignedevent)
- [UnlabeledEvent](objects.md#unlabeledevent)
- [UnlockedEvent](objects.md#unlockedevent)
- [UnmarkedAsDuplicateEvent](objects.md#unmarkedasduplicateevent)
- [UnpinnedEvent](objects.md#unpinnedevent)
- [UnsubscribedEvent](objects.md#unsubscribedevent)
- [UserBlockedEvent](objects.md#userblockedevent)

---

### PushAllowanceActor

<p>Types that can be an actor.</p>

### Possible types


- [App](objects.md#app)
- [Team](objects.md#team)
- [User](objects.md#user)

---

### Reactor

<p>Types that can be assigned to reactions.</p>

### Possible types


- [Bot](objects.md#bot)
- [Mannequin](objects.md#mannequin)
- [Organization](objects.md#organization)
- [User](objects.md#user)

---

### ReferencedSubject

<p>Any referencable object</p>

### Possible types


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest)

---

### RenamedTitleSubject

<p>An object which has a renamable title</p>

### Possible types


- [Issue](objects.md#issue)
- [PullRequest](objects.md#pullrequest)

---

### RequestedReviewer

<p>Types that can be requested reviewers.</p>

### Possible types


- [Mannequin](objects.md#mannequin)
- [Team](objects.md#team)
- [User](objects.md#user)

---

### ReviewDismissalAllowanceActor

<p>Types that can be an actor.</p>

### Possible types


- [Team](objects.md#team)
- [User](objects.md#user)

---

### SearchResultItem

<p>The results of a search.</p>

### Possible types


- [App](objects.md#app)
- [Discussion](objects.md#discussion)
- [Issue](objects.md#issue)
- [MarketplaceListing](objects.md#marketplacelisting)
- [Organization](objects.md#organization)
- [PullRequest](objects.md#pullrequest)
- [Repository](objects.md#repository)
- [User](objects.md#user)

---

### Sponsor

<p>Entities that can sponsor others via GitHub Sponsors</p>

### Possible types


- [Organization](objects.md#organization)
- [User](objects.md#user)

---

### SponsorableItem

<p>Entities that can be sponsored via GitHub Sponsors</p>

### Possible types


- [Organization](objects.md#organization)
- [User](objects.md#user)

---

### StatusCheckRollupContext

<p>Types that can be inside a StatusCheckRollup context.</p>

### Possible types


- [CheckRun](objects.md#checkrun)
- [StatusContext](objects.md#statuscontext)

---

### VerifiableDomainOwner

<p>Types that can own a verifiable domain.</p>

### Possible types


- [Enterprise](objects.md#enterprise)
- [Organization](objects.md#organization)

---