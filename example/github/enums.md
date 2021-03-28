# Enums

### About enums

[Enums](https://graphql.github.io/graphql-spec/June2018/#sec-Enums) represent possible sets of values for a field.

### StarOrderField

<p>Properties by which star connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>STARRED_AT</strong></td>
    <td><p>Allows ordering a list of stars by when they were created.</p></td>
  </tr>
</table>

---

### MilestoneOrderField

<p>Properties by which milestone connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order milestones by when they were created.</p></td>
  </tr>
  <tr>
    <td><strong>DUE_DATE</strong></td>
    <td><p>Order milestones by when they are due.</p></td>
  </tr>
  <tr>
    <td><strong>NUMBER</strong></td>
    <td><p>Order milestones by their number.</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order milestones by when they were last updated.</p></td>
  </tr>
</table>

---

### ProjectOrderField

<p>Properties by which project connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order projects by creation time</p></td>
  </tr>
  <tr>
    <td><strong>NAME</strong></td>
    <td><p>Order projects by name</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order projects by update time</p></td>
  </tr>
</table>

---

### RepositoryInvitationOrderField

<p>Properties by which repository invitation connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order repository invitations by creation time</p></td>
  </tr>
  <tr>
    <td><strong>INVITEE_LOGIN</strong></td>
    <td><p>Order repository invitations by invitee login</p></td>
  </tr>
</table>

---

### PullRequestUpdateState

<p>The possible target states when updating a pull request.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CLOSED</strong></td>
    <td><p>A pull request that has been closed without being merged.</p></td>
  </tr>
  <tr>
    <td><strong>OPEN</strong></td>
    <td><p>A pull request that is still open.</p></td>
  </tr>
</table>

---

### TeamMemberOrderField

<p>Properties by which team member connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order team members by creation time</p></td>
  </tr>
  <tr>
    <td><strong>LOGIN</strong></td>
    <td><p>Order team members by login</p></td>
  </tr>
</table>

---

### SubscriptionState

<p>The possible states of a subscription.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>IGNORED</strong></td>
    <td><p>The User is never notified.</p></td>
  </tr>
  <tr>
    <td><strong>SUBSCRIBED</strong></td>
    <td><p>The User is notified of all conversations.</p></td>
  </tr>
  <tr>
    <td><strong>UNSUBSCRIBED</strong></td>
    <td><p>The User is only notified when participating or @mentioned.</p></td>
  </tr>
</table>

---

### TeamMemberRole

<p>The possible team member roles; either &lsquo;maintainer&rsquo; or &lsquo;member&rsquo;.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>MAINTAINER</strong></td>
    <td><p>A team maintainer has permission to add and remove team members.</p></td>
  </tr>
  <tr>
    <td><strong>MEMBER</strong></td>
    <td><p>A team member has no administrative permissions on the team.</p></td>
  </tr>
</table>

---

### PullRequestTimelineItemsItemType

<p>The possible item types found in a timeline.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADDED_TO_PROJECT_EVENT</strong></td>
    <td><p>Represents a &lsquo;added_to_project&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>ASSIGNED_EVENT</strong></td>
    <td><p>Represents an &lsquo;assigned&rsquo; event on any assignable object.</p></td>
  </tr>
  <tr>
    <td><strong>AUTOMATIC_BASE_CHANGE_FAILED_EVENT</strong></td>
    <td><p>Represents a &lsquo;automatic_base_change_failed&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>AUTOMATIC_BASE_CHANGE_SUCCEEDED_EVENT</strong></td>
    <td><p>Represents a &lsquo;automatic_base_change_succeeded&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>AUTO_MERGE_DISABLED_EVENT</strong></td>
    <td><p>Represents a &lsquo;auto_merge_disabled&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>AUTO_MERGE_ENABLED_EVENT</strong></td>
    <td><p>Represents a &lsquo;auto_merge_enabled&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>AUTO_REBASE_ENABLED_EVENT</strong></td>
    <td><p>Represents a &lsquo;auto_rebase_enabled&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>AUTO_SQUASH_ENABLED_EVENT</strong></td>
    <td><p>Represents a &lsquo;auto_squash_enabled&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>BASE_REF_CHANGED_EVENT</strong></td>
    <td><p>Represents a &lsquo;base_ref_changed&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>BASE_REF_DELETED_EVENT</strong></td>
    <td><p>Represents a &lsquo;base_ref_deleted&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>BASE_REF_FORCE_PUSHED_EVENT</strong></td>
    <td><p>Represents a &lsquo;base_ref_force_pushed&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CLOSED_EVENT</strong></td>
    <td><p>Represents a &lsquo;closed&rsquo; event on any <code>Closable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>COMMENT_DELETED_EVENT</strong></td>
    <td><p>Represents a &lsquo;comment_deleted&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CONNECTED_EVENT</strong></td>
    <td><p>Represents a &lsquo;connected&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CONVERTED_NOTE_TO_ISSUE_EVENT</strong></td>
    <td><p>Represents a &lsquo;converted_note_to_issue&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CONVERT_TO_DRAFT_EVENT</strong></td>
    <td><p>Represents a &lsquo;convert_to_draft&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CROSS_REFERENCED_EVENT</strong></td>
    <td><p>Represents a mention made by one issue or pull request to another.</p></td>
  </tr>
  <tr>
    <td><strong>DEMILESTONED_EVENT</strong></td>
    <td><p>Represents a &lsquo;demilestoned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>DEPLOYED_EVENT</strong></td>
    <td><p>Represents a &lsquo;deployed&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>DEPLOYMENT_ENVIRONMENT_CHANGED_EVENT</strong></td>
    <td><p>Represents a &lsquo;deployment_environment_changed&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>DISCONNECTED_EVENT</strong></td>
    <td><p>Represents a &lsquo;disconnected&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>HEAD_REF_DELETED_EVENT</strong></td>
    <td><p>Represents a &lsquo;head_ref_deleted&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>HEAD_REF_FORCE_PUSHED_EVENT</strong></td>
    <td><p>Represents a &lsquo;head_ref_force_pushed&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>HEAD_REF_RESTORED_EVENT</strong></td>
    <td><p>Represents a &lsquo;head_ref_restored&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>ISSUE_COMMENT</strong></td>
    <td><p>Represents a comment on an Issue.</p></td>
  </tr>
  <tr>
    <td><strong>LABELED_EVENT</strong></td>
    <td><p>Represents a &lsquo;labeled&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>LOCKED_EVENT</strong></td>
    <td><p>Represents a &lsquo;locked&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MARKED_AS_DUPLICATE_EVENT</strong></td>
    <td><p>Represents a &lsquo;marked_as_duplicate&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MENTIONED_EVENT</strong></td>
    <td><p>Represents a &lsquo;mentioned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MERGED_EVENT</strong></td>
    <td><p>Represents a &lsquo;merged&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MILESTONED_EVENT</strong></td>
    <td><p>Represents a &lsquo;milestoned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MOVED_COLUMNS_IN_PROJECT_EVENT</strong></td>
    <td><p>Represents a &lsquo;moved_columns_in_project&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>PINNED_EVENT</strong></td>
    <td><p>Represents a &lsquo;pinned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST_COMMIT</strong></td>
    <td><p>Represents a Git commit part of a pull request.</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST_COMMIT_COMMENT_THREAD</strong></td>
    <td><p>Represents a commit comment thread part of a pull request.</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST_REVIEW</strong></td>
    <td><p>A review object for a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST_REVIEW_THREAD</strong></td>
    <td><p>A threaded list of comments for a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST_REVISION_MARKER</strong></td>
    <td><p>Represents the latest point in the pull request timeline for which the viewer has seen the pull request&rsquo;s commits.</p></td>
  </tr>
  <tr>
    <td><strong>READY_FOR_REVIEW_EVENT</strong></td>
    <td><p>Represents a &lsquo;ready_for_review&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>REFERENCED_EVENT</strong></td>
    <td><p>Represents a &lsquo;referenced&rsquo; event on a given <code>ReferencedSubject</code>.</p></td>
  </tr>
  <tr>
    <td><strong>REMOVED_FROM_PROJECT_EVENT</strong></td>
    <td><p>Represents a &lsquo;removed_from_project&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>RENAMED_TITLE_EVENT</strong></td>
    <td><p>Represents a &lsquo;renamed&rsquo; event on a given issue or pull request</p></td>
  </tr>
  <tr>
    <td><strong>REOPENED_EVENT</strong></td>
    <td><p>Represents a &lsquo;reopened&rsquo; event on any <code>Closable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>REVIEW_DISMISSED_EVENT</strong></td>
    <td><p>Represents a &lsquo;review_dismissed&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>REVIEW_REQUESTED_EVENT</strong></td>
    <td><p>Represents an &lsquo;review_requested&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>REVIEW_REQUEST_REMOVED_EVENT</strong></td>
    <td><p>Represents an &lsquo;review_request_removed&rsquo; event on a given pull request.</p></td>
  </tr>
  <tr>
    <td><strong>SUBSCRIBED_EVENT</strong></td>
    <td><p>Represents a &lsquo;subscribed&rsquo; event on a given <code>Subscribable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>TRANSFERRED_EVENT</strong></td>
    <td><p>Represents a &lsquo;transferred&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNASSIGNED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unassigned&rsquo; event on any assignable object.</p></td>
  </tr>
  <tr>
    <td><strong>UNLABELED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unlabeled&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNLOCKED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unlocked&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNMARKED_AS_DUPLICATE_EVENT</strong></td>
    <td><p>Represents an &lsquo;unmarked_as_duplicate&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNPINNED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unpinned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNSUBSCRIBED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unsubscribed&rsquo; event on a given <code>Subscribable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>USER_BLOCKED_EVENT</strong></td>
    <td><p>Represents a &lsquo;user_blocked&rsquo; event on a given user.</p></td>
  </tr>
</table>

---

### EnterpriseEnabledSettingValue

<p>The possible values for an enabled/no policy enterprise setting.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ENABLED</strong></td>
    <td><p>The setting is enabled for organizations in the enterprise.</p></td>
  </tr>
  <tr>
    <td><strong>NO_POLICY</strong></td>
    <td><p>There is no policy set for organizations in the enterprise.</p></td>
  </tr>
</table>

---

### LockReason

<p>The possible reasons that an issue or pull request was locked.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>OFF_TOPIC</strong></td>
    <td><p>The issue or pull request was locked because the conversation was off-topic.</p></td>
  </tr>
  <tr>
    <td><strong>RESOLVED</strong></td>
    <td><p>The issue or pull request was locked because the conversation was resolved.</p></td>
  </tr>
  <tr>
    <td><strong>SPAM</strong></td>
    <td><p>The issue or pull request was locked because the conversation was spam.</p></td>
  </tr>
  <tr>
    <td><strong>TOO_HEATED</strong></td>
    <td><p>The issue or pull request was locked because the conversation was too heated.</p></td>
  </tr>
</table>

---

### RefOrderField

<p>Properties by which ref connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALPHABETICAL</strong></td>
    <td><p>Order refs by their alphanumeric name</p></td>
  </tr>
  <tr>
    <td><strong>TAG_COMMIT_DATE</strong></td>
    <td><p>Order refs by underlying commit date if the ref prefix is refs/tags/</p></td>
  </tr>
</table>

---

### TeamRole

<p>The role of a user on a team.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>User has admin rights on the team.</p></td>
  </tr>
  <tr>
    <td><strong>MEMBER</strong></td>
    <td><p>User is a member of the team.</p></td>
  </tr>
</table>

---

### OrgCreateAuditEntryBillingPlan

<p>The billing plans available for organizations.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BUSINESS</strong></td>
    <td><p>Team Plan</p></td>
  </tr>
  <tr>
    <td><strong>BUSINESS_PLUS</strong></td>
    <td><p>Enterprise Cloud Plan</p></td>
  </tr>
  <tr>
    <td><strong>FREE</strong></td>
    <td><p>Free Plan</p></td>
  </tr>
  <tr>
    <td><strong>TIERED_PER_SEAT</strong></td>
    <td><p>Tiered Per Seat Plan</p></td>
  </tr>
  <tr>
    <td><strong>UNLIMITED</strong></td>
    <td><p>Legacy Unlimited Plan</p></td>
  </tr>
</table>

---

### EnterpriseServerUserAccountsUploadSyncState

<p>Synchronization state of the Enterprise Server user accounts upload</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>FAILURE</strong></td>
    <td><p>The synchronization of the upload failed.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>The synchronization of the upload is pending.</p></td>
  </tr>
  <tr>
    <td><strong>SUCCESS</strong></td>
    <td><p>The synchronization of the upload succeeded.</p></td>
  </tr>
</table>

---

### PullRequestReviewEvent

<p>The possible events to perform on a pull request review.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>APPROVE</strong></td>
    <td><p>Submit feedback and approve merging these changes.</p></td>
  </tr>
  <tr>
    <td><strong>COMMENT</strong></td>
    <td><p>Submit general feedback without explicit approval.</p></td>
  </tr>
  <tr>
    <td><strong>DISMISS</strong></td>
    <td><p>Dismiss review so it now longer effects merging.</p></td>
  </tr>
  <tr>
    <td><strong>REQUEST_CHANGES</strong></td>
    <td><p>Submit feedback that must be addressed before merging.</p></td>
  </tr>
</table>

---

### RepoAccessAuditEntryVisibility

<p>The privacy of a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>The repository is visible only to users in the same business.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>The repository is visible only to those with explicit access.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>The repository is visible to everyone.</p></td>
  </tr>
</table>

---

### ProjectColumnPurpose

<p>The semantic purpose of the column - todo, in progress, or done.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>DONE</strong></td>
    <td><p>The column contains cards which are complete</p></td>
  </tr>
  <tr>
    <td><strong>IN_PROGRESS</strong></td>
    <td><p>The column contains cards which are currently being worked on</p></td>
  </tr>
  <tr>
    <td><strong>TODO</strong></td>
    <td><p>The column contains cards still to be worked on</p></td>
  </tr>
</table>

---

### VerifiableDomainOrderField

<p>Properties by which verifiable domain connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order verifiable domains by their creation date.</p></td>
  </tr>
  <tr>
    <td><strong>DOMAIN</strong></td>
    <td><p>Order verifiable domains by the domain name.</p></td>
  </tr>
</table>

---

### IpAllowListEntryOrderField

<p>Properties by which IP allow list entry connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALLOW_LIST_VALUE</strong></td>
    <td><p>Order IP allow list entries by the allow list value.</p></td>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order IP allow list entries by creation time.</p></td>
  </tr>
</table>

---

### MergeableState

<p>Whether or not a PullRequest can be merged.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CONFLICTING</strong></td>
    <td><p>The pull request cannot be merged due to merge conflicts.</p></td>
  </tr>
  <tr>
    <td><strong>MERGEABLE</strong></td>
    <td><p>The pull request can be merged.</p></td>
  </tr>
  <tr>
    <td><strong>UNKNOWN</strong></td>
    <td><p>The mergeability of the pull request is still being calculated.</p></td>
  </tr>
</table>

---

### OrganizationMemberRole

<p>The possible roles within an organization for its members.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>The user is an administrator of the organization.</p></td>
  </tr>
  <tr>
    <td><strong>MEMBER</strong></td>
    <td><p>The user is a member of the organization.</p></td>
  </tr>
</table>

---

### ProjectCardState

<p>Various content states of a ProjectCard</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CONTENT_ONLY</strong></td>
    <td><p>The card has content only.</p></td>
  </tr>
  <tr>
    <td><strong>NOTE_ONLY</strong></td>
    <td><p>The card has a note only.</p></td>
  </tr>
  <tr>
    <td><strong>REDACTED</strong></td>
    <td><p>The card is redacted.</p></td>
  </tr>
</table>

---

### OperationType

<p>The corresponding operation type for the action</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ACCESS</strong></td>
    <td><p>An existing resource was accessed</p></td>
  </tr>
  <tr>
    <td><strong>AUTHENTICATION</strong></td>
    <td><p>A resource performed an authentication event</p></td>
  </tr>
  <tr>
    <td><strong>CREATE</strong></td>
    <td><p>A new resource was created</p></td>
  </tr>
  <tr>
    <td><strong>MODIFY</strong></td>
    <td><p>An existing resource was modified</p></td>
  </tr>
  <tr>
    <td><strong>REMOVE</strong></td>
    <td><p>An existing resource was removed</p></td>
  </tr>
  <tr>
    <td><strong>RESTORE</strong></td>
    <td><p>An existing resource was restored</p></td>
  </tr>
  <tr>
    <td><strong>TRANSFER</strong></td>
    <td><p>An existing resource was transferred between multiple resources</p></td>
  </tr>
</table>

---

### PackageVersionOrderField

<p>Properties by which package version connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order package versions by creation time</p></td>
  </tr>
</table>

---

### EnterpriseServerUserAccountEmailOrderField

<p>Properties by which Enterprise Server user account email connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>EMAIL</strong></td>
    <td><p>Order emails by email</p></td>
  </tr>
</table>

---

### PackageType

<p>The possible types of a package.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>DEBIAN</strong></td>
    <td><p>A debian package.</p></td>
  </tr>
  <tr>
    <td><strong>DOCKER</strong></td>
    <td><p>A docker image.</p></td>
  </tr>
  <tr>
    <td><strong>MAVEN</strong></td>
    <td><p>A maven package.</p></td>
  </tr>
  <tr>
    <td><strong>NPM</strong></td>
    <td><p>An npm package.</p></td>
  </tr>
  <tr>
    <td><strong>NUGET</strong></td>
    <td><p>A nuget package.</p></td>
  </tr>
  <tr>
    <td><strong>PYPI</strong></td>
    <td><p>A python package.</p></td>
  </tr>
  <tr>
    <td><strong>RUBYGEMS</strong></td>
    <td><p>A rubygems package.</p></td>
  </tr>
</table>

---

### SecurityAdvisoryOrderField

<p>Properties by which security advisory connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>PUBLISHED_AT</strong></td>
    <td><p>Order advisories by publication time</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order advisories by update time</p></td>
  </tr>
</table>

---

### DiffSide

<p>The possible sides of a diff.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>LEFT</strong></td>
    <td><p>The left side of the diff.</p></td>
  </tr>
  <tr>
    <td><strong>RIGHT</strong></td>
    <td><p>The right side of the diff.</p></td>
  </tr>
</table>

---

### RepositoryInteractionLimitExpiry

<p>The length for a repository interaction limit to be enabled for.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ONE_DAY</strong></td>
    <td><p>The interaction limit will expire after 1 day.</p></td>
  </tr>
  <tr>
    <td><strong>ONE_MONTH</strong></td>
    <td><p>The interaction limit will expire after 1 month.</p></td>
  </tr>
  <tr>
    <td><strong>ONE_WEEK</strong></td>
    <td><p>The interaction limit will expire after 1 week.</p></td>
  </tr>
  <tr>
    <td><strong>SIX_MONTHS</strong></td>
    <td><p>The interaction limit will expire after 6 months.</p></td>
  </tr>
  <tr>
    <td><strong>THREE_DAYS</strong></td>
    <td><p>The interaction limit will expire after 3 days.</p></td>
  </tr>
</table>

---

### DeploymentState

<p>The possible states in which a deployment can be.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ABANDONED</strong></td>
    <td><p>The pending deployment was not updated after 30 minutes.</p></td>
  </tr>
  <tr>
    <td><strong>ACTIVE</strong></td>
    <td><p>The deployment is currently active.</p></td>
  </tr>
  <tr>
    <td><strong>DESTROYED</strong></td>
    <td><p>An inactive transient deployment.</p></td>
  </tr>
  <tr>
    <td><strong>ERROR</strong></td>
    <td><p>The deployment experienced an error.</p></td>
  </tr>
  <tr>
    <td><strong>FAILURE</strong></td>
    <td><p>The deployment has failed.</p></td>
  </tr>
  <tr>
    <td><strong>INACTIVE</strong></td>
    <td><p>The deployment is inactive.</p></td>
  </tr>
  <tr>
    <td><strong>IN_PROGRESS</strong></td>
    <td><p>The deployment is in progress.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>The deployment is pending.</p></td>
  </tr>
  <tr>
    <td><strong>QUEUED</strong></td>
    <td><p>The deployment has queued</p></td>
  </tr>
  <tr>
    <td><strong>WAITING</strong></td>
    <td><p>The deployment is waiting.</p></td>
  </tr>
</table>

---

### SavedReplyOrderField

<p>Properties by which saved reply connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order saved reply by when they were updated.</p></td>
  </tr>
</table>

---

### PackageFileOrderField

<p>Properties by which package file connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order package files by creation time</p></td>
  </tr>
</table>

---

### ReactionOrderField

<p>A list of fields that reactions can be ordered by.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Allows ordering a list of reactions by when they were created.</p></td>
  </tr>
</table>

---

### RepoChangeMergeSettingAuditEntryMergeType

<p>The merge options available for pull requests to this repository.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>MERGE</strong></td>
    <td><p>The pull request is added to the base branch in a merge commit.</p></td>
  </tr>
  <tr>
    <td><strong>REBASE</strong></td>
    <td><p>Commits from the pull request are added onto the base branch individually without a merge commit.</p></td>
  </tr>
  <tr>
    <td><strong>SQUASH</strong></td>
    <td><p>The pull request&rsquo;s commits are squashed into a single commit before they are merged to the base branch.</p></td>
  </tr>
</table>

---

### IssueState

<p>The possible states of an issue.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CLOSED</strong></td>
    <td><p>An issue that has been closed</p></td>
  </tr>
  <tr>
    <td><strong>OPEN</strong></td>
    <td><p>An issue that is still open</p></td>
  </tr>
</table>

---

### IssueTimelineItemsItemType

<p>The possible item types found in a timeline.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADDED_TO_PROJECT_EVENT</strong></td>
    <td><p>Represents a &lsquo;added_to_project&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>ASSIGNED_EVENT</strong></td>
    <td><p>Represents an &lsquo;assigned&rsquo; event on any assignable object.</p></td>
  </tr>
  <tr>
    <td><strong>CLOSED_EVENT</strong></td>
    <td><p>Represents a &lsquo;closed&rsquo; event on any <code>Closable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>COMMENT_DELETED_EVENT</strong></td>
    <td><p>Represents a &lsquo;comment_deleted&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CONNECTED_EVENT</strong></td>
    <td><p>Represents a &lsquo;connected&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CONVERTED_NOTE_TO_ISSUE_EVENT</strong></td>
    <td><p>Represents a &lsquo;converted_note_to_issue&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>CROSS_REFERENCED_EVENT</strong></td>
    <td><p>Represents a mention made by one issue or pull request to another.</p></td>
  </tr>
  <tr>
    <td><strong>DEMILESTONED_EVENT</strong></td>
    <td><p>Represents a &lsquo;demilestoned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>DISCONNECTED_EVENT</strong></td>
    <td><p>Represents a &lsquo;disconnected&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>ISSUE_COMMENT</strong></td>
    <td><p>Represents a comment on an Issue.</p></td>
  </tr>
  <tr>
    <td><strong>LABELED_EVENT</strong></td>
    <td><p>Represents a &lsquo;labeled&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>LOCKED_EVENT</strong></td>
    <td><p>Represents a &lsquo;locked&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MARKED_AS_DUPLICATE_EVENT</strong></td>
    <td><p>Represents a &lsquo;marked_as_duplicate&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MENTIONED_EVENT</strong></td>
    <td><p>Represents a &lsquo;mentioned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MILESTONED_EVENT</strong></td>
    <td><p>Represents a &lsquo;milestoned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>MOVED_COLUMNS_IN_PROJECT_EVENT</strong></td>
    <td><p>Represents a &lsquo;moved_columns_in_project&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>PINNED_EVENT</strong></td>
    <td><p>Represents a &lsquo;pinned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>REFERENCED_EVENT</strong></td>
    <td><p>Represents a &lsquo;referenced&rsquo; event on a given <code>ReferencedSubject</code>.</p></td>
  </tr>
  <tr>
    <td><strong>REMOVED_FROM_PROJECT_EVENT</strong></td>
    <td><p>Represents a &lsquo;removed_from_project&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>RENAMED_TITLE_EVENT</strong></td>
    <td><p>Represents a &lsquo;renamed&rsquo; event on a given issue or pull request</p></td>
  </tr>
  <tr>
    <td><strong>REOPENED_EVENT</strong></td>
    <td><p>Represents a &lsquo;reopened&rsquo; event on any <code>Closable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>SUBSCRIBED_EVENT</strong></td>
    <td><p>Represents a &lsquo;subscribed&rsquo; event on a given <code>Subscribable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>TRANSFERRED_EVENT</strong></td>
    <td><p>Represents a &lsquo;transferred&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNASSIGNED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unassigned&rsquo; event on any assignable object.</p></td>
  </tr>
  <tr>
    <td><strong>UNLABELED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unlabeled&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNLOCKED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unlocked&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNMARKED_AS_DUPLICATE_EVENT</strong></td>
    <td><p>Represents an &lsquo;unmarked_as_duplicate&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNPINNED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unpinned&rsquo; event on a given issue or pull request.</p></td>
  </tr>
  <tr>
    <td><strong>UNSUBSCRIBED_EVENT</strong></td>
    <td><p>Represents an &lsquo;unsubscribed&rsquo; event on a given <code>Subscribable</code>.</p></td>
  </tr>
  <tr>
    <td><strong>USER_BLOCKED_EVENT</strong></td>
    <td><p>Represents a &lsquo;user_blocked&rsquo; event on a given user.</p></td>
  </tr>
</table>

---

### ContributionLevel

<p>Varying levels of contributions from none to many.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>FIRST_QUARTILE</strong></td>
    <td><p>Lowest 25% of days of contributions.</p></td>
  </tr>
  <tr>
    <td><strong>FOURTH_QUARTILE</strong></td>
    <td><p>Highest 25% of days of contributions. More contributions than the third quartile.</p></td>
  </tr>
  <tr>
    <td><strong>NONE</strong></td>
    <td><p>No contributions occurred.</p></td>
  </tr>
  <tr>
    <td><strong>SECOND_QUARTILE</strong></td>
    <td><p>Second lowest 25% of days of contributions. More contributions than the first quartile.</p></td>
  </tr>
  <tr>
    <td><strong>THIRD_QUARTILE</strong></td>
    <td><p>Second highest 25% of days of contributions. More contributions than second quartile, less than the fourth quartile.</p></td>
  </tr>
</table>

---

### AuditLogOrderField

<p>Properties by which Audit Log connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order audit log entries by timestamp</p></td>
  </tr>
</table>

---

### LabelOrderField

<p>Properties by which label connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order labels by creation time</p></td>
  </tr>
  <tr>
    <td><strong>NAME</strong></td>
    <td><p>Order labels by name</p></td>
  </tr>
</table>

---

### PullRequestState

<p>The possible states of a pull request.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CLOSED</strong></td>
    <td><p>A pull request that has been closed without being merged.</p></td>
  </tr>
  <tr>
    <td><strong>MERGED</strong></td>
    <td><p>A pull request that has been closed by being merged.</p></td>
  </tr>
  <tr>
    <td><strong>OPEN</strong></td>
    <td><p>A pull request that is still open.</p></td>
  </tr>
</table>

---

### EnterpriseEnabledDisabledSettingValue

<p>The possible values for an enabled/disabled enterprise setting.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>DISABLED</strong></td>
    <td><p>The setting is disabled for organizations in the enterprise.</p></td>
  </tr>
  <tr>
    <td><strong>ENABLED</strong></td>
    <td><p>The setting is enabled for organizations in the enterprise.</p></td>
  </tr>
  <tr>
    <td><strong>NO_POLICY</strong></td>
    <td><p>There is no policy set for organizations in the enterprise.</p></td>
  </tr>
</table>

---

### OrganizationInvitationType

<p>The possible organization invitation types.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>EMAIL</strong></td>
    <td><p>The invitation was to an email address.</p></td>
  </tr>
  <tr>
    <td><strong>USER</strong></td>
    <td><p>The invitation was to an existing user.</p></td>
  </tr>
</table>

---

### RepositoryOrderField

<p>Properties by which repository connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order repositories by creation time</p></td>
  </tr>
  <tr>
    <td><strong>NAME</strong></td>
    <td><p>Order repositories by name</p></td>
  </tr>
  <tr>
    <td><strong>PUSHED_AT</strong></td>
    <td><p>Order repositories by push time</p></td>
  </tr>
  <tr>
    <td><strong>STARGAZERS</strong></td>
    <td><p>Order repositories by number of stargazers</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order repositories by update time</p></td>
  </tr>
</table>

---

### EnterpriseAdministratorRole

<p>The possible administrator roles in an enterprise account.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BILLING_MANAGER</strong></td>
    <td><p>Represents a billing manager of the enterprise account.</p></td>
  </tr>
  <tr>
    <td><strong>OWNER</strong></td>
    <td><p>Represents an owner of the enterprise account.</p></td>
  </tr>
</table>

---

### EnterpriseAdministratorInvitationOrderField

<p>Properties by which enterprise administrator invitation connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order enterprise administrator member invitations by creation time</p></td>
  </tr>
</table>

---

### OrgUpdateDefaultRepositoryPermissionAuditEntryPermission

<p>The default permission a repository can have in an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>Can read, clone, push, and add collaborators to repositories.</p></td>
  </tr>
  <tr>
    <td><strong>NONE</strong></td>
    <td><p>No default permission value.</p></td>
  </tr>
  <tr>
    <td><strong>READ</strong></td>
    <td><p>Can read and clone repositories.</p></td>
  </tr>
  <tr>
    <td><strong>WRITE</strong></td>
    <td><p>Can read, clone and push to repositories.</p></td>
  </tr>
</table>

---

### ReportedContentClassifiers

<p>The reasons a piece of content can be reported or minimized.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ABUSE</strong></td>
    <td><p>An abusive or harassing piece of content</p></td>
  </tr>
  <tr>
    <td><strong>DUPLICATE</strong></td>
    <td><p>A duplicated piece of content</p></td>
  </tr>
  <tr>
    <td><strong>OFF_TOPIC</strong></td>
    <td><p>An irrelevant piece of content</p></td>
  </tr>
  <tr>
    <td><strong>OUTDATED</strong></td>
    <td><p>An outdated piece of content</p></td>
  </tr>
  <tr>
    <td><strong>RESOLVED</strong></td>
    <td><p>The content has been resolved</p></td>
  </tr>
  <tr>
    <td><strong>SPAM</strong></td>
    <td><p>A spammy piece of content</p></td>
  </tr>
</table>

---

### RepoAddMemberAuditEntryVisibility

<p>The privacy of a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>The repository is visible only to users in the same business.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>The repository is visible only to those with explicit access.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>The repository is visible to everyone.</p></td>
  </tr>
</table>

---

### GistOrderField

<p>Properties by which gist connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order gists by creation time</p></td>
  </tr>
  <tr>
    <td><strong>PUSHED_AT</strong></td>
    <td><p>Order gists by push time</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order gists by update time</p></td>
  </tr>
</table>

---

### SecurityAdvisorySeverity

<p>Severity of the vulnerability.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CRITICAL</strong></td>
    <td><p>Critical.</p></td>
  </tr>
  <tr>
    <td><strong>HIGH</strong></td>
    <td><p>High.</p></td>
  </tr>
  <tr>
    <td><strong>LOW</strong></td>
    <td><p>Low.</p></td>
  </tr>
  <tr>
    <td><strong>MODERATE</strong></td>
    <td><p>Moderate.</p></td>
  </tr>
</table>

---

### SecurityVulnerabilityOrderField

<p>Properties by which security vulnerability connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order vulnerability by update time</p></td>
  </tr>
</table>

---

### SamlSignatureAlgorithm

<p>The possible signature algorithms used to sign SAML requests for a Identity Provider.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>RSA_SHA1</strong></td>
    <td><p>RSA-SHA1</p></td>
  </tr>
  <tr>
    <td><strong>RSA_SHA256</strong></td>
    <td><p>RSA-SHA256</p></td>
  </tr>
  <tr>
    <td><strong>RSA_SHA384</strong></td>
    <td><p>RSA-SHA384</p></td>
  </tr>
  <tr>
    <td><strong>RSA_SHA512</strong></td>
    <td><p>RSA-SHA512</p></td>
  </tr>
</table>

---

### GitSignatureState

<p>The state of a Git signature.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BAD_CERT</strong></td>
    <td><p>The signing certificate or its chain could not be verified</p></td>
  </tr>
  <tr>
    <td><strong>BAD_EMAIL</strong></td>
    <td><p>Invalid email used for signing</p></td>
  </tr>
  <tr>
    <td><strong>EXPIRED_KEY</strong></td>
    <td><p>Signing key expired</p></td>
  </tr>
  <tr>
    <td><strong>GPGVERIFY_ERROR</strong></td>
    <td><p>Internal error - the GPG verification service misbehaved</p></td>
  </tr>
  <tr>
    <td><strong>GPGVERIFY_UNAVAILABLE</strong></td>
    <td><p>Internal error - the GPG verification service is unavailable at the moment</p></td>
  </tr>
  <tr>
    <td><strong>INVALID</strong></td>
    <td><p>Invalid signature</p></td>
  </tr>
  <tr>
    <td><strong>MALFORMED_SIG</strong></td>
    <td><p>Malformed signature</p></td>
  </tr>
  <tr>
    <td><strong>NOT_SIGNING_KEY</strong></td>
    <td><p>The usage flags for the key that signed this don&rsquo;t allow signing</p></td>
  </tr>
  <tr>
    <td><strong>NO_USER</strong></td>
    <td><p>Email used for signing not known to GitHub</p></td>
  </tr>
  <tr>
    <td><strong>OCSP_ERROR</strong></td>
    <td><p>Valid signature, though certificate revocation check failed</p></td>
  </tr>
  <tr>
    <td><strong>OCSP_PENDING</strong></td>
    <td><p>Valid signature, pending certificate revocation checking</p></td>
  </tr>
  <tr>
    <td><strong>OCSP_REVOKED</strong></td>
    <td><p>One or more certificates in chain has been revoked</p></td>
  </tr>
  <tr>
    <td><strong>UNKNOWN_KEY</strong></td>
    <td><p>Key used for signing not known to GitHub</p></td>
  </tr>
  <tr>
    <td><strong>UNKNOWN_SIG_TYPE</strong></td>
    <td><p>Unknown signature type</p></td>
  </tr>
  <tr>
    <td><strong>UNSIGNED</strong></td>
    <td><p>Unsigned</p></td>
  </tr>
  <tr>
    <td><strong>UNVERIFIED_EMAIL</strong></td>
    <td><p>Email used for signing unverified on GitHub</p></td>
  </tr>
  <tr>
    <td><strong>VALID</strong></td>
    <td><p>Valid signature and verified by GitHub</p></td>
  </tr>
</table>

---

### SponsorsGoalKind

<p>The different kinds of goals a GitHub Sponsors member can have.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>MONTHLY_SPONSORSHIP_AMOUNT</strong></td>
    <td><p>The goal is about getting a certain dollar amount from sponsorships each month.</p></td>
  </tr>
  <tr>
    <td><strong>TOTAL_SPONSORS_COUNT</strong></td>
    <td><p>The goal is about reaching a certain number of sponsors.</p></td>
  </tr>
</table>

---

### EnterpriseMembersCanMakePurchasesSettingValue

<p>The possible values for the members can make purchases setting.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>DISABLED</strong></td>
    <td><p>The setting is disabled for organizations in the enterprise.</p></td>
  </tr>
  <tr>
    <td><strong>ENABLED</strong></td>
    <td><p>The setting is enabled for organizations in the enterprise.</p></td>
  </tr>
</table>

---

### OrganizationOrderField

<p>Properties by which organization connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order organizations by creation time</p></td>
  </tr>
  <tr>
    <td><strong>LOGIN</strong></td>
    <td><p>Order organizations by login</p></td>
  </tr>
</table>

---

### CheckStatusState

<p>The possible states for a check suite or run status.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COMPLETED</strong></td>
    <td><p>The check suite or run has been completed.</p></td>
  </tr>
  <tr>
    <td><strong>IN_PROGRESS</strong></td>
    <td><p>The check suite or run is in progress.</p></td>
  </tr>
  <tr>
    <td><strong>QUEUED</strong></td>
    <td><p>The check suite or run has been queued.</p></td>
  </tr>
  <tr>
    <td><strong>REQUESTED</strong></td>
    <td><p>The check suite or run has been requested.</p></td>
  </tr>
  <tr>
    <td><strong>WAITING</strong></td>
    <td><p>The check suite or run is in waiting state.</p></td>
  </tr>
</table>

---

### EnterpriseServerInstallationOrderField

<p>Properties by which Enterprise Server installation connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order Enterprise Server installations by creation time</p></td>
  </tr>
  <tr>
    <td><strong>CUSTOMER_NAME</strong></td>
    <td><p>Order Enterprise Server installations by customer name</p></td>
  </tr>
  <tr>
    <td><strong>HOST_NAME</strong></td>
    <td><p>Order Enterprise Server installations by host name</p></td>
  </tr>
</table>

---

### RepositoryLockReason

<p>The possible reasons a given repository could be in a locked state.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BILLING</strong></td>
    <td><p>The repository is locked due to a billing related reason.</p></td>
  </tr>
  <tr>
    <td><strong>MIGRATING</strong></td>
    <td><p>The repository is locked due to a migration.</p></td>
  </tr>
  <tr>
    <td><strong>MOVING</strong></td>
    <td><p>The repository is locked due to a move.</p></td>
  </tr>
  <tr>
    <td><strong>RENAME</strong></td>
    <td><p>The repository is locked due to a rename.</p></td>
  </tr>
</table>

---

### OrgRemoveMemberAuditEntryMembershipType

<p>The type of membership a user has with an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>Organization administrators have full access and can change several settings,
including the names of repositories that belong to the Organization and Owners
team membership. In addition, organization admins can delete the organization
and all of its repositories.</p></td>
  </tr>
  <tr>
    <td><strong>BILLING_MANAGER</strong></td>
    <td><p>A billing manager is a user who manages the billing settings for the Organization, such as updating payment information.</p></td>
  </tr>
  <tr>
    <td><strong>DIRECT_MEMBER</strong></td>
    <td><p>A direct member is a user that is a member of the Organization.</p></td>
  </tr>
  <tr>
    <td><strong>OUTSIDE_COLLABORATOR</strong></td>
    <td><p>An outside collaborator is a person who isn&rsquo;t explicitly a member of the
Organization, but who has Read, Write, or Admin permissions to one or more
repositories in the organization.</p></td>
  </tr>
  <tr>
    <td><strong>UNAFFILIATED</strong></td>
    <td><p>An unaffiliated collaborator is a person who is not a member of the
Organization and does not have access to any repositories in the Organization.</p></td>
  </tr>
</table>

---

### PackageOrderField

<p>Properties by which package connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order packages by creation time</p></td>
  </tr>
</table>

---

### TeamOrderField

<p>Properties by which team connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>NAME</strong></td>
    <td><p>Allows ordering a list of teams by name.</p></td>
  </tr>
</table>

---

### TeamRepositoryOrderField

<p>Properties by which team repository connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order repositories by creation time</p></td>
  </tr>
  <tr>
    <td><strong>NAME</strong></td>
    <td><p>Order repositories by name</p></td>
  </tr>
  <tr>
    <td><strong>PERMISSION</strong></td>
    <td><p>Order repositories by permission</p></td>
  </tr>
  <tr>
    <td><strong>PUSHED_AT</strong></td>
    <td><p>Order repositories by push time</p></td>
  </tr>
  <tr>
    <td><strong>STARGAZERS</strong></td>
    <td><p>Order repositories by number of stargazers</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order repositories by update time</p></td>
  </tr>
</table>

---

### OrgRemoveBillingManagerAuditEntryReason

<p>The reason a billing manager was removed from an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>SAML_EXTERNAL_IDENTITY_MISSING</strong></td>
    <td><p>SAML external identity missing</p></td>
  </tr>
  <tr>
    <td><strong>SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY</strong></td>
    <td><p>SAML SSO enforcement requires an external identity</p></td>
  </tr>
  <tr>
    <td><strong>TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE</strong></td>
    <td><p>The organization required 2FA of its billing managers and this user did not have 2FA enabled.</p></td>
  </tr>
</table>

---

### PullRequestMergeMethod

<p>Represents available types of methods to use when merging a pull request.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>MERGE</strong></td>
    <td><p>Add all commits from the head branch to the base branch with a merge commit.</p></td>
  </tr>
  <tr>
    <td><strong>REBASE</strong></td>
    <td><p>Add all commits from the head branch onto the base branch individually.</p></td>
  </tr>
  <tr>
    <td><strong>SQUASH</strong></td>
    <td><p>Combine all commits from the head branch into a single commit in the base branch.</p></td>
  </tr>
</table>

---

### SponsorsTierOrderField

<p>Properties by which Sponsors tiers connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order tiers by creation time.</p></td>
  </tr>
  <tr>
    <td><strong>MONTHLY_PRICE_IN_CENTS</strong></td>
    <td><p>Order tiers by their monthly price in cents</p></td>
  </tr>
</table>

---

### RepositoryVisibility

<p>The repository&rsquo;s visibility level.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>The repository is visible only to users in the same business.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>The repository is visible only to those with explicit access.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>The repository is visible to everyone.</p></td>
  </tr>
</table>

---

### FundingPlatform

<p>The possible funding platforms for repository funding links.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COMMUNITY_BRIDGE</strong></td>
    <td><p>Community Bridge funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>CUSTOM</strong></td>
    <td><p>Custom funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>GITHUB</strong></td>
    <td><p>GitHub funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>ISSUEHUNT</strong></td>
    <td><p>IssueHunt funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>KO_FI</strong></td>
    <td><p>Ko-fi funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>LIBERAPAY</strong></td>
    <td><p>Liberapay funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>OPEN_COLLECTIVE</strong></td>
    <td><p>Open Collective funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>OTECHIE</strong></td>
    <td><p>Otechie funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>PATREON</strong></td>
    <td><p>Patreon funding platform.</p></td>
  </tr>
  <tr>
    <td><strong>TIDELIFT</strong></td>
    <td><p>Tidelift funding platform.</p></td>
  </tr>
</table>

---

### PinnableItemType

<p>Represents items that can be pinned to a profile page or dashboard.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>GIST</strong></td>
    <td><p>A gist.</p></td>
  </tr>
  <tr>
    <td><strong>ISSUE</strong></td>
    <td><p>An issue.</p></td>
  </tr>
  <tr>
    <td><strong>ORGANIZATION</strong></td>
    <td><p>An organization.</p></td>
  </tr>
  <tr>
    <td><strong>PROJECT</strong></td>
    <td><p>A project.</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST</strong></td>
    <td><p>A pull request.</p></td>
  </tr>
  <tr>
    <td><strong>REPOSITORY</strong></td>
    <td><p>A repository.</p></td>
  </tr>
  <tr>
    <td><strong>TEAM</strong></td>
    <td><p>A team.</p></td>
  </tr>
  <tr>
    <td><strong>USER</strong></td>
    <td><p>A user.</p></td>
  </tr>
</table>

---

### SecurityAdvisoryIdentifierType

<p>Identifier formats available for advisories.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CVE</strong></td>
    <td><p>Common Vulnerabilities and Exposures Identifier.</p></td>
  </tr>
  <tr>
    <td><strong>GHSA</strong></td>
    <td><p>GitHub Security Advisory ID.</p></td>
  </tr>
</table>

---

### OrgRemoveOutsideCollaboratorAuditEntryReason

<p>The reason an outside collaborator was removed from an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>SAML_EXTERNAL_IDENTITY_MISSING</strong></td>
    <td><p>SAML external identity missing</p></td>
  </tr>
  <tr>
    <td><strong>TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE</strong></td>
    <td><p>The organization required 2FA of its billing managers and this user did not have 2FA enabled.</p></td>
  </tr>
</table>

---

### OrderDirection

<p>Possible directions in which to order a list of items when provided an <code>orderBy</code> argument.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ASC</strong></td>
    <td><p>Specifies an ascending order for a given <code>orderBy</code> argument.</p></td>
  </tr>
  <tr>
    <td><strong>DESC</strong></td>
    <td><p>Specifies a descending order for a given <code>orderBy</code> argument.</p></td>
  </tr>
</table>

---

### PullRequestReviewCommentState

<p>The possible states of a pull request review comment.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>A comment that is part of a pending review</p></td>
  </tr>
  <tr>
    <td><strong>SUBMITTED</strong></td>
    <td><p>A comment that is part of a submitted review</p></td>
  </tr>
</table>

---

### PullRequestReviewState

<p>The possible states of a pull request review.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>APPROVED</strong></td>
    <td><p>A review allowing the pull request to merge.</p></td>
  </tr>
  <tr>
    <td><strong>CHANGES_REQUESTED</strong></td>
    <td><p>A review blocking the pull request from merging.</p></td>
  </tr>
  <tr>
    <td><strong>COMMENTED</strong></td>
    <td><p>An informational review.</p></td>
  </tr>
  <tr>
    <td><strong>DISMISSED</strong></td>
    <td><p>A review that has been dismissed.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>A review that has not yet been submitted.</p></td>
  </tr>
</table>

---

### RepositoryPrivacy

<p>The privacy of a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>Private</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>Public</p></td>
  </tr>
</table>

---

### RepoCreateAuditEntryVisibility

<p>The privacy of a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>The repository is visible only to users in the same business.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>The repository is visible only to those with explicit access.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>The repository is visible to everyone.</p></td>
  </tr>
</table>

---

### CommitContributionOrderField

<p>Properties by which commit contribution connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COMMIT_COUNT</strong></td>
    <td><p>Order commit contributions by how many commits they represent.</p></td>
  </tr>
  <tr>
    <td><strong>OCCURRED_AT</strong></td>
    <td><p>Order commit contributions by when they were made.</p></td>
  </tr>
</table>

---

### OauthApplicationCreateAuditEntryState

<p>The state of an OAuth Application when it was created.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ACTIVE</strong></td>
    <td><p>The OAuth Application was active and allowed to have OAuth Accesses.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING_DELETION</strong></td>
    <td><p>The OAuth Application was in the process of being deleted.</p></td>
  </tr>
  <tr>
    <td><strong>SUSPENDED</strong></td>
    <td><p>The OAuth Application was suspended from generating OAuth Accesses due to abuse or security concerns.</p></td>
  </tr>
</table>

---

### RepoArchivedAuditEntryVisibility

<p>The privacy of a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>The repository is visible only to users in the same business.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>The repository is visible only to those with explicit access.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>The repository is visible to everyone.</p></td>
  </tr>
</table>

---

### DefaultRepositoryPermissionField

<p>The possible default permissions for repositories.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>Can read, write, and administrate repos by default</p></td>
  </tr>
  <tr>
    <td><strong>NONE</strong></td>
    <td><p>No access</p></td>
  </tr>
  <tr>
    <td><strong>READ</strong></td>
    <td><p>Can read repos by default</p></td>
  </tr>
  <tr>
    <td><strong>WRITE</strong></td>
    <td><p>Can read and write repos by default</p></td>
  </tr>
</table>

---

### UserBlockDuration

<p>The possible durations that a user can be blocked for.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ONE_DAY</strong></td>
    <td><p>The user was blocked for 1 day</p></td>
  </tr>
  <tr>
    <td><strong>ONE_MONTH</strong></td>
    <td><p>The user was blocked for 30 days</p></td>
  </tr>
  <tr>
    <td><strong>ONE_WEEK</strong></td>
    <td><p>The user was blocked for 7 days</p></td>
  </tr>
  <tr>
    <td><strong>PERMANENT</strong></td>
    <td><p>The user was blocked permanently</p></td>
  </tr>
  <tr>
    <td><strong>THREE_DAYS</strong></td>
    <td><p>The user was blocked for 3 days</p></td>
  </tr>
</table>

---

### OrgRemoveMemberAuditEntryReason

<p>The reason a member was removed from an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>SAML_EXTERNAL_IDENTITY_MISSING</strong></td>
    <td><p>SAML external identity missing</p></td>
  </tr>
  <tr>
    <td><strong>SAML_SSO_ENFORCEMENT_REQUIRES_EXTERNAL_IDENTITY</strong></td>
    <td><p>SAML SSO enforcement requires an external identity</p></td>
  </tr>
  <tr>
    <td><strong>TWO_FACTOR_ACCOUNT_RECOVERY</strong></td>
    <td><p>User was removed from organization during account recovery</p></td>
  </tr>
  <tr>
    <td><strong>TWO_FACTOR_REQUIREMENT_NON_COMPLIANCE</strong></td>
    <td><p>The organization required 2FA of its billing managers and this user did not have 2FA enabled.</p></td>
  </tr>
  <tr>
    <td><strong>USER_ACCOUNT_DELETED</strong></td>
    <td><p>User account has been deleted</p></td>
  </tr>
</table>

---

### LanguageOrderField

<p>Properties by which language connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>SIZE</strong></td>
    <td><p>Order languages by the size of all files containing the language</p></td>
  </tr>
</table>

---

### ReleaseOrderField

<p>Properties by which release connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order releases by creation time</p></td>
  </tr>
  <tr>
    <td><strong>NAME</strong></td>
    <td><p>Order releases alphabetically by name</p></td>
  </tr>
</table>

---

### OrganizationInvitationRole

<p>The possible organization invitation roles.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>The user is invited to be an admin of the organization.</p></td>
  </tr>
  <tr>
    <td><strong>BILLING_MANAGER</strong></td>
    <td><p>The user is invited to be a billing manager of the organization.</p></td>
  </tr>
  <tr>
    <td><strong>DIRECT_MEMBER</strong></td>
    <td><p>The user is invited to be a direct member of the organization.</p></td>
  </tr>
  <tr>
    <td><strong>REINSTATE</strong></td>
    <td><p>The user&rsquo;s previous role will be reinstated.</p></td>
  </tr>
</table>

---

### SamlDigestAlgorithm

<p>The possible digest algorithms used to sign SAML requests for an identity provider.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>SHA1</strong></td>
    <td><p>SHA1</p></td>
  </tr>
  <tr>
    <td><strong>SHA256</strong></td>
    <td><p>SHA256</p></td>
  </tr>
  <tr>
    <td><strong>SHA384</strong></td>
    <td><p>SHA384</p></td>
  </tr>
  <tr>
    <td><strong>SHA512</strong></td>
    <td><p>SHA512</p></td>
  </tr>
</table>

---

### SponsorableOrderField

<p>Properties by which sponsorable connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>LOGIN</strong></td>
    <td><p>Order sponsorable entities by login (username).</p></td>
  </tr>
</table>

---

### IpAllowListEnabledSettingValue

<p>The possible values for the IP allow list enabled setting.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>DISABLED</strong></td>
    <td><p>The setting is disabled for the owner.</p></td>
  </tr>
  <tr>
    <td><strong>ENABLED</strong></td>
    <td><p>The setting is enabled for the owner.</p></td>
  </tr>
</table>

---

### OrganizationMembersCanCreateRepositoriesSettingValue

<p>The possible values for the members can create repositories setting on an organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>Members will be able to create public and private repositories.</p></td>
  </tr>
  <tr>
    <td><strong>DISABLED</strong></td>
    <td><p>Members will not be able to create public or private repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>Members will be able to create only private repositories.</p></td>
  </tr>
</table>

---

### SearchType

<p>Represents the individual results of a search.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ISSUE</strong></td>
    <td><p>Returns results matching issues in repositories.</p></td>
  </tr>
  <tr>
    <td><strong>REPOSITORY</strong></td>
    <td><p>Returns results matching repositories.</p></td>
  </tr>
  <tr>
    <td><strong>USER</strong></td>
    <td><p>Returns results matching users and organizations on GitHub.</p></td>
  </tr>
</table>

---

### EnterpriseMembersCanCreateRepositoriesSettingValue

<p>The possible values for the enterprise members can create repositories setting.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>Members will be able to create public and private repositories.</p></td>
  </tr>
  <tr>
    <td><strong>DISABLED</strong></td>
    <td><p>Members will not be able to create public or private repositories.</p></td>
  </tr>
  <tr>
    <td><strong>NO_POLICY</strong></td>
    <td><p>Organization administrators choose whether to allow members to create repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>Members will be able to create only private repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>Members will be able to create only public repositories.</p></td>
  </tr>
</table>

---

### PullRequestOrderField

<p>Properties by which pull_requests connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order pull_requests by creation time</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order pull_requests by update time</p></td>
  </tr>
</table>

---

### OrgAddMemberAuditEntryPermission

<p>The permissions available to members on an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>Can read, clone, push, and add collaborators to repositories.</p></td>
  </tr>
  <tr>
    <td><strong>READ</strong></td>
    <td><p>Can read and clone repositories.</p></td>
  </tr>
</table>

---

### OrgUpdateMemberAuditEntryPermission

<p>The permissions available to members on an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>Can read, clone, push, and add collaborators to repositories.</p></td>
  </tr>
  <tr>
    <td><strong>READ</strong></td>
    <td><p>Can read and clone repositories.</p></td>
  </tr>
</table>

---

### RepositoryInteractionLimit

<p>A repository interaction limit.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COLLABORATORS_ONLY</strong></td>
    <td><p>Users that are not collaborators will not be able to interact with the repository.</p></td>
  </tr>
  <tr>
    <td><strong>CONTRIBUTORS_ONLY</strong></td>
    <td><p>Users that have not previously committed to a repository’s default branch will be unable to interact with the repository.</p></td>
  </tr>
  <tr>
    <td><strong>EXISTING_USERS</strong></td>
    <td><p>Users that have recently created their account will be unable to interact with the repository.</p></td>
  </tr>
  <tr>
    <td><strong>NO_LIMIT</strong></td>
    <td><p>No interaction limits are enabled.</p></td>
  </tr>
</table>

---

### TeamDiscussionOrderField

<p>Properties by which team discussion connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Allows chronological ordering of team discussions.</p></td>
  </tr>
</table>

---

### RepositoryContributionType

<p>The reason a repository is listed as &lsquo;contributed&rsquo;.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COMMIT</strong></td>
    <td><p>Created a commit</p></td>
  </tr>
  <tr>
    <td><strong>ISSUE</strong></td>
    <td><p>Created an issue</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST</strong></td>
    <td><p>Created a pull request</p></td>
  </tr>
  <tr>
    <td><strong>PULL_REQUEST_REVIEW</strong></td>
    <td><p>Reviewed a pull request</p></td>
  </tr>
  <tr>
    <td><strong>REPOSITORY</strong></td>
    <td><p>Created the repository</p></td>
  </tr>
</table>

---

### OrgUpdateMemberRepositoryCreationPermissionAuditEntryVisibility

<p>The permissions available for repository creation on an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>All organization members are restricted from creating any repositories.</p></td>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>All organization members are restricted from creating internal repositories.</p></td>
  </tr>
  <tr>
    <td><strong>NONE</strong></td>
    <td><p>All organization members are allowed to create any repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>All organization members are restricted from creating private repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE_INTERNAL</strong></td>
    <td><p>All organization members are restricted from creating private or internal repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>All organization members are restricted from creating public repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC_INTERNAL</strong></td>
    <td><p>All organization members are restricted from creating public or internal repositories.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC_PRIVATE</strong></td>
    <td><p>All organization members are restricted from creating public or private repositories.</p></td>
  </tr>
</table>

---

### CommentCannotUpdateReason

<p>The possible errors that will prevent a user from updating a comment.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ARCHIVED</strong></td>
    <td><p>Unable to create comment because repository is archived.</p></td>
  </tr>
  <tr>
    <td><strong>DENIED</strong></td>
    <td><p>You cannot update this comment</p></td>
  </tr>
  <tr>
    <td><strong>INSUFFICIENT_ACCESS</strong></td>
    <td><p>You must be the author or have write access to this repository to update this comment.</p></td>
  </tr>
  <tr>
    <td><strong>LOCKED</strong></td>
    <td><p>Unable to create comment because issue is locked.</p></td>
  </tr>
  <tr>
    <td><strong>LOGIN_REQUIRED</strong></td>
    <td><p>You must be logged in to update this comment.</p></td>
  </tr>
  <tr>
    <td><strong>MAINTENANCE</strong></td>
    <td><p>Repository is under maintenance.</p></td>
  </tr>
  <tr>
    <td><strong>VERIFIED_EMAIL_REQUIRED</strong></td>
    <td><p>At least one email address must be verified to update this comment.</p></td>
  </tr>
</table>

---

### EnterpriseDefaultRepositoryPermissionSettingValue

<p>The possible values for the enterprise default repository permission setting.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>Organization members will be able to clone, pull, push, and add new collaborators to all organization repositories.</p></td>
  </tr>
  <tr>
    <td><strong>NONE</strong></td>
    <td><p>Organization members will only be able to clone and pull public repositories.</p></td>
  </tr>
  <tr>
    <td><strong>NO_POLICY</strong></td>
    <td><p>Organizations in the enterprise choose default repository permissions for their members.</p></td>
  </tr>
  <tr>
    <td><strong>READ</strong></td>
    <td><p>Organization members will be able to clone and pull all organization repositories.</p></td>
  </tr>
  <tr>
    <td><strong>WRITE</strong></td>
    <td><p>Organization members will be able to clone, pull, and push all organization repositories.</p></td>
  </tr>
</table>

---

### SponsorshipPrivacy

<p>The privacy of a sponsorship</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>Private</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>Public</p></td>
  </tr>
</table>

---

### FileViewedState

<p>The possible viewed states of a file .</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>DISMISSED</strong></td>
    <td><p>The file has new changes since last viewed.</p></td>
  </tr>
  <tr>
    <td><strong>UNVIEWED</strong></td>
    <td><p>The file has not been marked as viewed.</p></td>
  </tr>
  <tr>
    <td><strong>VIEWED</strong></td>
    <td><p>The file has been marked as viewed.</p></td>
  </tr>
</table>

---

### TeamReviewAssignmentAlgorithm

<p>The possible team review assignment algorithms</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>LOAD_BALANCE</strong></td>
    <td><p>Balance review load across the entire team</p></td>
  </tr>
  <tr>
    <td><strong>ROUND_ROBIN</strong></td>
    <td><p>Alternate reviews between each team member</p></td>
  </tr>
</table>

---

### TopicSuggestionDeclineReason

<p>Reason that the suggested topic is declined.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>NOT_RELEVANT</strong></td>
    <td><p>The suggested topic is not relevant to the repository.</p></td>
  </tr>
  <tr>
    <td><strong>PERSONAL_PREFERENCE</strong></td>
    <td><p>The viewer does not like the suggested topic.</p></td>
  </tr>
  <tr>
    <td><strong>TOO_GENERAL</strong></td>
    <td><p>The suggested topic is too general for the repository.</p></td>
  </tr>
  <tr>
    <td><strong>TOO_SPECIFIC</strong></td>
    <td><p>The suggested topic is too specific for the repository (e.g. #ruby-on-rails-version-4-2-1).</p></td>
  </tr>
</table>

---

### IssueCommentOrderField

<p>Properties by which issue comment connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order issue comments by update time</p></td>
  </tr>
</table>

---

### PullRequestReviewDecision

<p>The review status of a pull request.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>APPROVED</strong></td>
    <td><p>The pull request has received an approving review.</p></td>
  </tr>
  <tr>
    <td><strong>CHANGES_REQUESTED</strong></td>
    <td><p>Changes have been requested on the pull request.</p></td>
  </tr>
  <tr>
    <td><strong>REVIEW_REQUIRED</strong></td>
    <td><p>A review is required before the pull request can be merged.</p></td>
  </tr>
</table>

---

### TeamPrivacy

<p>The possible team privacy values.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>SECRET</strong></td>
    <td><p>A secret team can only be seen by its members.</p></td>
  </tr>
  <tr>
    <td><strong>VISIBLE</strong></td>
    <td><p>A visible team can be seen and @mentioned by every member of the organization.</p></td>
  </tr>
</table>

---

### NotificationRestrictionSettingValue

<p>The possible values for the notification restriction setting.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>DISABLED</strong></td>
    <td><p>The setting is disabled for the owner.</p></td>
  </tr>
  <tr>
    <td><strong>ENABLED</strong></td>
    <td><p>The setting is enabled for the owner.</p></td>
  </tr>
</table>

---

### MilestoneState

<p>The possible states of a milestone.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CLOSED</strong></td>
    <td><p>A milestone that has been closed.</p></td>
  </tr>
  <tr>
    <td><strong>OPEN</strong></td>
    <td><p>A milestone that is still open.</p></td>
  </tr>
</table>

---

### SecurityAdvisoryEcosystem

<p>The possible ecosystems of a security vulnerability&rsquo;s package.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COMPOSER</strong></td>
    <td><p>PHP packages hosted at packagist.org</p></td>
  </tr>
  <tr>
    <td><strong>MAVEN</strong></td>
    <td><p>Java artifacts hosted at the Maven central repository</p></td>
  </tr>
  <tr>
    <td><strong>NPM</strong></td>
    <td><p>JavaScript packages hosted at npmjs.com</p></td>
  </tr>
  <tr>
    <td><strong>NUGET</strong></td>
    <td><p>.NET packages hosted at the NuGet Gallery</p></td>
  </tr>
  <tr>
    <td><strong>PIP</strong></td>
    <td><p>Python packages hosted at PyPI.org</p></td>
  </tr>
  <tr>
    <td><strong>RUBYGEMS</strong></td>
    <td><p>Ruby gems hosted at RubyGems.org</p></td>
  </tr>
</table>

---

### ProjectState

<p>State of the project; either &lsquo;open&rsquo; or &lsquo;closed&rsquo;</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CLOSED</strong></td>
    <td><p>The project is closed.</p></td>
  </tr>
  <tr>
    <td><strong>OPEN</strong></td>
    <td><p>The project is open.</p></td>
  </tr>
</table>

---

### RepositoryInteractionLimitOrigin

<p>Indicates where an interaction limit is configured.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ORGANIZATION</strong></td>
    <td><p>A limit that is configured at the organization level.</p></td>
  </tr>
  <tr>
    <td><strong>REPOSITORY</strong></td>
    <td><p>A limit that is configured at the repository level.</p></td>
  </tr>
  <tr>
    <td><strong>USER</strong></td>
    <td><p>A limit that is configured at the user-wide level.</p></td>
  </tr>
</table>

---

### RepositoryAffiliation

<p>The affiliation of a user to a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COLLABORATOR</strong></td>
    <td><p>Repositories that the user has been added to as a collaborator.</p></td>
  </tr>
  <tr>
    <td><strong>ORGANIZATION_MEMBER</strong></td>
    <td><p>Repositories that the user has access to through being a member of an
organization. This includes every repository on every team that the user is on.</p></td>
  </tr>
  <tr>
    <td><strong>OWNER</strong></td>
    <td><p>Repositories that are owned by the authenticated user.</p></td>
  </tr>
</table>

---

### DeploymentOrderField

<p>Properties by which deployment connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order collection by creation time</p></td>
  </tr>
</table>

---

### ProjectCardArchivedState

<p>The possible archived states of a project card.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ARCHIVED</strong></td>
    <td><p>A project card that is archived</p></td>
  </tr>
  <tr>
    <td><strong>NOT_ARCHIVED</strong></td>
    <td><p>A project card that is not archived</p></td>
  </tr>
</table>

---

### CheckConclusionState

<p>The possible states for a check suite or run conclusion.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ACTION_REQUIRED</strong></td>
    <td><p>The check suite or run requires action.</p></td>
  </tr>
  <tr>
    <td><strong>CANCELLED</strong></td>
    <td><p>The check suite or run has been cancelled.</p></td>
  </tr>
  <tr>
    <td><strong>FAILURE</strong></td>
    <td><p>The check suite or run has failed.</p></td>
  </tr>
  <tr>
    <td><strong>NEUTRAL</strong></td>
    <td><p>The check suite or run was neutral.</p></td>
  </tr>
  <tr>
    <td><strong>SKIPPED</strong></td>
    <td><p>The check suite or run was skipped.</p></td>
  </tr>
  <tr>
    <td><strong>STALE</strong></td>
    <td><p>The check suite or run was marked stale by GitHub. Only GitHub can use this conclusion.</p></td>
  </tr>
  <tr>
    <td><strong>STARTUP_FAILURE</strong></td>
    <td><p>The check suite or run has failed at startup.</p></td>
  </tr>
  <tr>
    <td><strong>SUCCESS</strong></td>
    <td><p>The check suite or run has succeeded.</p></td>
  </tr>
  <tr>
    <td><strong>TIMED_OUT</strong></td>
    <td><p>The check suite or run has timed out.</p></td>
  </tr>
</table>

---

### IssueOrderField

<p>Properties by which issue connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COMMENTS</strong></td>
    <td><p>Order issues by comment count</p></td>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order issues by creation time</p></td>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order issues by update time</p></td>
  </tr>
</table>

---

### RepoRemoveMemberAuditEntryVisibility

<p>The privacy of a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>The repository is visible only to users in the same business.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>The repository is visible only to those with explicit access.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>The repository is visible to everyone.</p></td>
  </tr>
</table>

---

### RepositoryPermission

<p>The access level to a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ADMIN</strong></td>
    <td><p>Can read, clone, and push to this repository. Can also manage issues, pull
requests, and repository settings, including adding collaborators</p></td>
  </tr>
  <tr>
    <td><strong>MAINTAIN</strong></td>
    <td><p>Can read, clone, and push to this repository. They can also manage issues, pull requests, and some repository settings</p></td>
  </tr>
  <tr>
    <td><strong>READ</strong></td>
    <td><p>Can read and clone this repository. Can also open and comment on issues and pull requests</p></td>
  </tr>
  <tr>
    <td><strong>TRIAGE</strong></td>
    <td><p>Can read and clone this repository. Can also manage issues and pull requests</p></td>
  </tr>
  <tr>
    <td><strong>WRITE</strong></td>
    <td><p>Can read, clone, and push to this repository. Can also manage issues and pull requests</p></td>
  </tr>
</table>

---

### CheckRunType

<p>The possible types of check runs.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>Every check run available.</p></td>
  </tr>
  <tr>
    <td><strong>LATEST</strong></td>
    <td><p>The latest check run.</p></td>
  </tr>
</table>

---

### TeamMembershipType

<p>Defines which types of team members are included in the returned list. Can be one of IMMEDIATE, CHILD_TEAM or ALL.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>Includes immediate and child team members for the team.</p></td>
  </tr>
  <tr>
    <td><strong>CHILD_TEAM</strong></td>
    <td><p>Includes only child team members for the team.</p></td>
  </tr>
  <tr>
    <td><strong>IMMEDIATE</strong></td>
    <td><p>Includes only immediate members of the team.</p></td>
  </tr>
</table>

---

### CheckAnnotationLevel

<p>Represents an annotation&rsquo;s information level.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>FAILURE</strong></td>
    <td><p>An annotation indicating an inescapable error.</p></td>
  </tr>
  <tr>
    <td><strong>NOTICE</strong></td>
    <td><p>An annotation indicating some information.</p></td>
  </tr>
  <tr>
    <td><strong>WARNING</strong></td>
    <td><p>An annotation indicating an ignorable error.</p></td>
  </tr>
</table>

---

### EnterpriseUserDeployment

<p>The possible GitHub Enterprise deployments where this user can exist.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CLOUD</strong></td>
    <td><p>The user is part of a GitHub Enterprise Cloud deployment.</p></td>
  </tr>
  <tr>
    <td><strong>SERVER</strong></td>
    <td><p>The user is part of a GitHub Enterprise Server deployment.</p></td>
  </tr>
</table>

---

### IdentityProviderConfigurationState

<p>The possible states in which authentication can be configured with an identity provider.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CONFIGURED</strong></td>
    <td><p>Authentication with an identity provider is configured but not enforced.</p></td>
  </tr>
  <tr>
    <td><strong>ENFORCED</strong></td>
    <td><p>Authentication with an identity provider is configured and enforced.</p></td>
  </tr>
  <tr>
    <td><strong>UNCONFIGURED</strong></td>
    <td><p>Authentication with an identity provider is not configured.</p></td>
  </tr>
</table>

---

### ReactionContent

<p>Emojis that can be attached to Issues, Pull Requests and Comments.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CONFUSED</strong></td>
    <td><p>Represents the <code>:confused:</code> emoji.</p></td>
  </tr>
  <tr>
    <td><strong>EYES</strong></td>
    <td><p>Represents the <code>:eyes:</code> emoji.</p></td>
  </tr>
  <tr>
    <td><strong>HEART</strong></td>
    <td><p>Represents the <code>:heart:</code> emoji.</p></td>
  </tr>
  <tr>
    <td><strong>HOORAY</strong></td>
    <td><p>Represents the <code>:hooray:</code> emoji.</p></td>
  </tr>
  <tr>
    <td><strong>LAUGH</strong></td>
    <td><p>Represents the <code>:laugh:</code> emoji.</p></td>
  </tr>
  <tr>
    <td><strong>ROCKET</strong></td>
    <td><p>Represents the <code>:rocket:</code> emoji.</p></td>
  </tr>
  <tr>
    <td><strong>THUMBS_DOWN</strong></td>
    <td><p>Represents the <code>:-1:</code> emoji.</p></td>
  </tr>
  <tr>
    <td><strong>THUMBS_UP</strong></td>
    <td><p>Represents the <code>:+1:</code> emoji.</p></td>
  </tr>
</table>

---

### StatusState

<p>The possible commit status states.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ERROR</strong></td>
    <td><p>Status is errored.</p></td>
  </tr>
  <tr>
    <td><strong>EXPECTED</strong></td>
    <td><p>Status is expected.</p></td>
  </tr>
  <tr>
    <td><strong>FAILURE</strong></td>
    <td><p>Status is failing.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>Status is pending.</p></td>
  </tr>
  <tr>
    <td><strong>SUCCESS</strong></td>
    <td><p>Status is successful.</p></td>
  </tr>
</table>

---

### EnterpriseUserAccountMembershipRole

<p>The possible roles for enterprise membership.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>MEMBER</strong></td>
    <td><p>The user is a member of the enterprise membership.</p></td>
  </tr>
  <tr>
    <td><strong>OWNER</strong></td>
    <td><p>The user is an owner of the enterprise membership.</p></td>
  </tr>
</table>

---

### CollaboratorAffiliation

<p>Collaborators affiliation level with a subject.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>All collaborators the authenticated user can see.</p></td>
  </tr>
  <tr>
    <td><strong>DIRECT</strong></td>
    <td><p>All collaborators with permissions to an organization-owned subject, regardless of organization membership status.</p></td>
  </tr>
  <tr>
    <td><strong>OUTSIDE</strong></td>
    <td><p>All outside collaborators of an organization-owned subject.</p></td>
  </tr>
</table>

---

### ProjectTemplate

<p>GitHub-provided templates for Projects</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>AUTOMATED_KANBAN_V2</strong></td>
    <td><p>Create a board with v2 triggers to automatically move cards across To do, In progress and Done columns.</p></td>
  </tr>
  <tr>
    <td><strong>AUTOMATED_REVIEWS_KANBAN</strong></td>
    <td><p>Create a board with triggers to automatically move cards across columns with review automation.</p></td>
  </tr>
  <tr>
    <td><strong>BASIC_KANBAN</strong></td>
    <td><p>Create a board with columns for To do, In progress and Done.</p></td>
  </tr>
  <tr>
    <td><strong>BUG_TRIAGE</strong></td>
    <td><p>Create a board to triage and prioritize bugs with To do, priority, and Done columns.</p></td>
  </tr>
</table>

---

### UserStatusOrderField

<p>Properties by which user status connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>UPDATED_AT</strong></td>
    <td><p>Order user statuses by when they were updated.</p></td>
  </tr>
</table>

---

### TeamDiscussionCommentOrderField

<p>Properties by which team discussion comment connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>NUMBER</strong></td>
    <td><p>Allows sequential ordering of team discussion comments (which is equivalent to chronological ordering).</p></td>
  </tr>
</table>

---

### EnterpriseServerUserAccountsUploadOrderField

<p>Properties by which Enterprise Server user accounts upload connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order user accounts uploads by creation time</p></td>
  </tr>
</table>

---

### RequestableCheckStatusState

<p>The possible states that can be requested when creating a check run.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COMPLETED</strong></td>
    <td><p>The check suite or run has been completed.</p></td>
  </tr>
  <tr>
    <td><strong>IN_PROGRESS</strong></td>
    <td><p>The check suite or run is in progress.</p></td>
  </tr>
  <tr>
    <td><strong>QUEUED</strong></td>
    <td><p>The check suite or run has been queued.</p></td>
  </tr>
  <tr>
    <td><strong>WAITING</strong></td>
    <td><p>The check suite or run is in waiting state.</p></td>
  </tr>
</table>

---

### OrgRemoveOutsideCollaboratorAuditEntryMembershipType

<p>The type of membership a user has with an Organization.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BILLING_MANAGER</strong></td>
    <td><p>A billing manager is a user who manages the billing settings for the Organization, such as updating payment information.</p></td>
  </tr>
  <tr>
    <td><strong>OUTSIDE_COLLABORATOR</strong></td>
    <td><p>An outside collaborator is a person who isn&rsquo;t explicitly a member of the
Organization, but who has Read, Write, or Admin permissions to one or more
repositories in the organization.</p></td>
  </tr>
  <tr>
    <td><strong>UNAFFILIATED</strong></td>
    <td><p>An unaffiliated collaborator is a person who is not a member of the
Organization and does not have access to any repositories in the organization.</p></td>
  </tr>
</table>

---

### CommentAuthorAssociation

<p>A comment author association with repository.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>COLLABORATOR</strong></td>
    <td><p>Author has been invited to collaborate on the repository.</p></td>
  </tr>
  <tr>
    <td><strong>CONTRIBUTOR</strong></td>
    <td><p>Author has previously committed to the repository.</p></td>
  </tr>
  <tr>
    <td><strong>FIRST_TIMER</strong></td>
    <td><p>Author has not previously committed to GitHub.</p></td>
  </tr>
  <tr>
    <td><strong>FIRST_TIME_CONTRIBUTOR</strong></td>
    <td><p>Author has not previously committed to the repository.</p></td>
  </tr>
  <tr>
    <td><strong>MANNEQUIN</strong></td>
    <td><p>Author is a placeholder for an unclaimed user.</p></td>
  </tr>
  <tr>
    <td><strong>MEMBER</strong></td>
    <td><p>Author is a member of the organization that owns the repository.</p></td>
  </tr>
  <tr>
    <td><strong>NONE</strong></td>
    <td><p>Author has no association with the repository.</p></td>
  </tr>
  <tr>
    <td><strong>OWNER</strong></td>
    <td><p>Author is the owner of the repository.</p></td>
  </tr>
</table>

---

### GistPrivacy

<p>The privacy of a Gist</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ALL</strong></td>
    <td><p>Gists that are public and secret</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>Public</p></td>
  </tr>
  <tr>
    <td><strong>SECRET</strong></td>
    <td><p>Secret</p></td>
  </tr>
</table>

---

### RepoDestroyAuditEntryVisibility

<p>The privacy of a repository</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>INTERNAL</strong></td>
    <td><p>The repository is visible only to users in the same business.</p></td>
  </tr>
  <tr>
    <td><strong>PRIVATE</strong></td>
    <td><p>The repository is visible only to those with explicit access.</p></td>
  </tr>
  <tr>
    <td><strong>PUBLIC</strong></td>
    <td><p>The repository is visible to everyone.</p></td>
  </tr>
</table>

---

### MergeStateStatus

<p>Detailed status information about a pull request merge.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>BEHIND</strong></td>
    <td><p>The head ref is out of date.</p></td>
  </tr>
  <tr>
    <td><strong>BLOCKED</strong></td>
    <td><p>The merge is blocked.</p></td>
  </tr>
  <tr>
    <td><strong>CLEAN</strong></td>
    <td><p>Mergeable and passing commit status.</p></td>
  </tr>
  <tr>
    <td><strong>DIRTY</strong></td>
    <td><p>The merge commit cannot be cleanly created.</p></td>
  </tr>
  <tr>
    <td><strong>DRAFT</strong></td>
    <td><p>The merge is blocked due to the pull request being a draft.</p></td>
  </tr>
  <tr>
    <td><strong>HAS_HOOKS</strong></td>
    <td><p>Mergeable with passing commit status and pre-receive hooks.</p></td>
  </tr>
  <tr>
    <td><strong>UNKNOWN</strong></td>
    <td><p>The state cannot currently be determined.</p></td>
  </tr>
  <tr>
    <td><strong>UNSTABLE</strong></td>
    <td><p>Mergeable with non-passing commit status.</p></td>
  </tr>
</table>

---

### EnterpriseServerUserAccountOrderField

<p>Properties by which Enterprise Server user account connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>LOGIN</strong></td>
    <td><p>Order user accounts by login</p></td>
  </tr>
  <tr>
    <td><strong>REMOTE_CREATED_AT</strong></td>
    <td><p>Order user accounts by creation time on the Enterprise Server installation</p></td>
  </tr>
</table>

---

### EnterpriseMemberOrderField

<p>Properties by which enterprise member connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order enterprise members by creation time</p></td>
  </tr>
  <tr>
    <td><strong>LOGIN</strong></td>
    <td><p>Order enterprise members by login</p></td>
  </tr>
</table>

---

### SponsorshipOrderField

<p>Properties by which sponsorship connections can be ordered.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>CREATED_AT</strong></td>
    <td><p>Order sponsorship by creation time.</p></td>
  </tr>
</table>

---

### DeploymentStatusState

<p>The possible states for a deployment status.</p>

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>ERROR</strong></td>
    <td><p>The deployment experienced an error.</p></td>
  </tr>
  <tr>
    <td><strong>FAILURE</strong></td>
    <td><p>The deployment has failed.</p></td>
  </tr>
  <tr>
    <td><strong>INACTIVE</strong></td>
    <td><p>The deployment is inactive.</p></td>
  </tr>
  <tr>
    <td><strong>IN_PROGRESS</strong></td>
    <td><p>The deployment is in progress.</p></td>
  </tr>
  <tr>
    <td><strong>PENDING</strong></td>
    <td><p>The deployment is pending.</p></td>
  </tr>
  <tr>
    <td><strong>QUEUED</strong></td>
    <td><p>The deployment is queued</p></td>
  </tr>
  <tr>
    <td><strong>SUCCESS</strong></td>
    <td><p>The deployment was successful.</p></td>
  </tr>
  <tr>
    <td><strong>WAITING</strong></td>
    <td><p>The deployment is waiting.</p></td>
  </tr>
</table>

---