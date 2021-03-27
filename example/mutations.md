# Mutations

### About mutations

The root query for implementing GraphQL mutations.

### acceptEnterpriseAdministratorInvitation

Accepts a pending invitation for a user to become an administrator of an enterprise.

#### Input fields

- input ([AcceptEnterpriseAdministratorInvitationInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| invitation ([EnterpriseAdministratorInvitation](http://example.com)) | The invitation that was accepted. |
| message ([String](http://example.com)) | A message confirming the result of accepting an administrator invitation. |

---

### acceptTopicSuggestion

Applies a suggested topic to the repository.

#### Input fields

- input ([AcceptTopicSuggestionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| topic ([Topic](http://example.com)) | The accepted topic. |

---

### addAssigneesToAssignable

Adds assignees to an assignable object.

#### Input fields

- input ([AddAssigneesToAssignableInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| assignable ([Assignable](http://example.com)) | The item that was assigned. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### addComment

Adds a comment to an Issue or Pull Request.

#### Input fields

- input ([AddCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| commentEdge ([IssueCommentEdge](http://example.com)) | The edge from the subject's comment connection. |
| subject ([Node](http://example.com)) | The subject |
| timelineEdge ([IssueTimelineItemEdge](http://example.com)) | The edge from the subject's timeline connection. |

---

### addEnterpriseSupportEntitlement

Adds a support entitlement to an enterprise member.

#### Input fields

- input ([AddEnterpriseSupportEntitlementInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| message ([String](http://example.com)) | A message confirming the result of adding the support entitlement. |

---

### addLabelsToLabelable

Adds labels to a labelable object.

#### Input fields

- input ([AddLabelsToLabelableInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| labelable ([Labelable](http://example.com)) | The item that was labeled. |

---

### addProjectCard

Adds a card to a ProjectColumn. Either `contentId` or `note` must be provided but **not** both.

#### Input fields

- input ([AddProjectCardInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| cardEdge ([ProjectCardEdge](http://example.com)) | The edge from the ProjectColumn's card connection. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| projectColumn ([ProjectColumn](http://example.com)) | The ProjectColumn |

---

### addProjectColumn

Adds a column to a Project.

#### Input fields

- input ([AddProjectColumnInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| columnEdge ([ProjectColumnEdge](http://example.com)) | The edge from the project's column connection. |
| project ([Project](http://example.com)) | The project |

---

### addPullRequestReview

Adds a review to a Pull Request.

#### Input fields

- input ([AddPullRequestReviewInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequestReview ([PullRequestReview](http://example.com)) | The newly created pull request review. |
| reviewEdge ([PullRequestReviewEdge](http://example.com)) | The edge from the pull request's review connection. |

---

### addPullRequestReviewComment

Adds a comment to a review.

#### Input fields

- input ([AddPullRequestReviewCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| comment ([PullRequestReviewComment](http://example.com)) | The newly created comment. |
| commentEdge ([PullRequestReviewCommentEdge](http://example.com)) | The edge from the review's comment connection. |

---

### addPullRequestReviewThread

Adds a new thread to a pending Pull Request Review.

#### Input fields

- input ([AddPullRequestReviewThreadInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| thread ([PullRequestReviewThread](http://example.com)) | The newly created thread. |

---

### addReaction

Adds a reaction to a subject.

#### Input fields

- input ([AddReactionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| reaction ([Reaction](http://example.com)) | The reaction object. |
| subject ([Reactable](http://example.com)) | The reactable subject. |

---

### addStar

Adds a star to a Starrable.

#### Input fields

- input ([AddStarInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| starrable ([Starrable](http://example.com)) | The starrable. |

---

### addVerifiableDomain

Adds a verifiable domain to an owning account.

#### Input fields

- input ([AddVerifiableDomainInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| domain ([VerifiableDomain](http://example.com)) | The verifiable domain that was added. |

---

### approveVerifiableDomain

Approve a verifiable domain for notification delivery.

#### Input fields

- input ([ApproveVerifiableDomainInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| domain ([VerifiableDomain](http://example.com)) | The verifiable domain that was approved. |

---

### archiveRepository

Marks a repository as archived.

#### Input fields

- input ([ArchiveRepositoryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The repository that was marked as archived. |

---

### cancelEnterpriseAdminInvitation

Cancels a pending invitation for an administrator to join an enterprise.

#### Input fields

- input ([CancelEnterpriseAdminInvitationInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| invitation ([EnterpriseAdministratorInvitation](http://example.com)) | The invitation that was canceled. |
| message ([String](http://example.com)) | A message confirming the result of canceling an administrator invitation. |

---

### changeUserStatus

Update your status on GitHub.

#### Input fields

- input ([ChangeUserStatusInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| status ([UserStatus](http://example.com)) | Your updated status. |

---

### clearLabelsFromLabelable

Clears all labels from a labelable object.

#### Input fields

- input ([ClearLabelsFromLabelableInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| labelable ([Labelable](http://example.com)) | The item that was unlabeled. |

---

### cloneProject

Creates a new project by cloning configuration from an existing project.

#### Input fields

- input ([CloneProjectInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| jobStatusId ([String](http://example.com)) | The id of the JobStatus for populating cloned fields. |
| project ([Project](http://example.com)) | The new cloned project. |

---

### cloneTemplateRepository

Create a new repository with the same files and directory structure as a template repository.

#### Input fields

- input ([CloneTemplateRepositoryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The new repository. |

---

### closeIssue

Close an issue.

#### Input fields

- input ([CloseIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issue ([Issue](http://example.com)) | The issue that was closed. |

---

### closePullRequest

Close a pull request.

#### Input fields

- input ([ClosePullRequestInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The pull request that was closed. |

---

### convertProjectCardNoteToIssue

Convert a project note card to one associated with a newly created issue.

#### Input fields

- input ([ConvertProjectCardNoteToIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| projectCard ([ProjectCard](http://example.com)) | The updated ProjectCard. |

---

### createBranchProtectionRule

Create a new branch protection rule

#### Input fields

- input ([CreateBranchProtectionRuleInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| branchProtectionRule ([BranchProtectionRule](http://example.com)) | The newly created BranchProtectionRule. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### createCheckRun

Create a check run.

#### Input fields

- input ([CreateCheckRunInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| checkRun ([CheckRun](http://example.com)) | The newly created check run. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### createCheckSuite

Create a check suite

#### Input fields

- input ([CreateCheckSuiteInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| checkSuite ([CheckSuite](http://example.com)) | The newly created check suite. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### createContentAttachment

Create a content attachment.

#### Input fields

- input ([CreateContentAttachmentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| contentAttachment ([ContentAttachment](http://example.com)) | The newly created content attachment. |

---

### createDeployment

Creates a new deployment event.

#### Input fields

- input ([CreateDeploymentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| autoMerged ([Boolean](http://example.com)) | True if the default branch has been auto-merged into the deployment ref. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| deployment ([Deployment](http://example.com)) | The new deployment. |

---

### createDeploymentStatus

Create a deployment status.

#### Input fields

- input ([CreateDeploymentStatusInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| deploymentStatus ([DeploymentStatus](http://example.com)) | The new deployment status. |

---

### createEnterpriseOrganization

Creates an organization as part of an enterprise account.

#### Input fields

- input ([CreateEnterpriseOrganizationInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise that owns the created organization. |
| organization ([Organization](http://example.com)) | The organization that was created. |

---

### createIpAllowListEntry

Creates a new IP allow list entry.

#### Input fields

- input ([CreateIpAllowListEntryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| ipAllowListEntry ([IpAllowListEntry](http://example.com)) | The IP allow list entry that was created. |

---

### createIssue

Creates a new issue.

#### Input fields

- input ([CreateIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issue ([Issue](http://example.com)) | The new issue. |

---

### createLabel

Creates a new label.

#### Input fields

- input ([CreateLabelInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| label ([Label](http://example.com)) | The new label. |

---

### createProject

Creates a new project.

#### Input fields

- input ([CreateProjectInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| project ([Project](http://example.com)) | The new project. |

---

### createPullRequest

Create a new pull request

#### Input fields

- input ([CreatePullRequestInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The new pull request. |

---

### createRef

Create a new Git Ref.

#### Input fields

- input ([CreateRefInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| ref ([Ref](http://example.com)) | The newly created ref. |

---

### createRepository

Create a new repository.

#### Input fields

- input ([CreateRepositoryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The new repository. |

---

### createTeamDiscussion

Creates a new team discussion.

#### Input fields

- input ([CreateTeamDiscussionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| teamDiscussion ([TeamDiscussion](http://example.com)) | The new discussion. |

---

### createTeamDiscussionComment

Creates a new team discussion comment.

#### Input fields

- input ([CreateTeamDiscussionCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| teamDiscussionComment ([TeamDiscussionComment](http://example.com)) | The new comment. |

---

### declineTopicSuggestion

Rejects a suggested topic for the repository.

#### Input fields

- input ([DeclineTopicSuggestionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| topic ([Topic](http://example.com)) | The declined topic. |

---

### deleteBranchProtectionRule

Delete a branch protection rule

#### Input fields

- input ([DeleteBranchProtectionRuleInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### deleteDeployment

Deletes a deployment.

#### Input fields

- input ([DeleteDeploymentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### deleteIpAllowListEntry

Deletes an IP allow list entry.

#### Input fields

- input ([DeleteIpAllowListEntryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| ipAllowListEntry ([IpAllowListEntry](http://example.com)) | The IP allow list entry that was deleted. |

---

### deleteIssue

Deletes an Issue object.

#### Input fields

- input ([DeleteIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The repository the issue belonged to |

---

### deleteIssueComment

Deletes an IssueComment object.

#### Input fields

- input ([DeleteIssueCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### deleteLabel

Deletes a label.

#### Input fields

- input ([DeleteLabelInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### deletePackageVersion

Delete a package version.

#### Input fields

- input ([DeletePackageVersionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| success ([Boolean](http://example.com)) | Whether or not the operation succeeded. |

---

### deleteProject

Deletes a project.

#### Input fields

- input ([DeleteProjectInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| owner ([ProjectOwner](http://example.com)) | The repository or organization the project was removed from. |

---

### deleteProjectCard

Deletes a project card.

#### Input fields

- input ([DeleteProjectCardInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| column ([ProjectColumn](http://example.com)) | The column the deleted card was in. |
| deletedCardId ([ID](http://example.com)) | The deleted card ID. |

---

### deleteProjectColumn

Deletes a project column.

#### Input fields

- input ([DeleteProjectColumnInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| deletedColumnId ([ID](http://example.com)) | The deleted column ID. |
| project ([Project](http://example.com)) | The project the deleted column was in. |

---

### deletePullRequestReview

Deletes a pull request review.

#### Input fields

- input ([DeletePullRequestReviewInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequestReview ([PullRequestReview](http://example.com)) | The deleted pull request review. |

---

### deletePullRequestReviewComment

Deletes a pull request review comment.

#### Input fields

- input ([DeletePullRequestReviewCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequestReview ([PullRequestReview](http://example.com)) | The pull request review the deleted comment belonged to. |

---

### deleteRef

Delete a Git Ref.

#### Input fields

- input ([DeleteRefInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### deleteTeamDiscussion

Deletes a team discussion.

#### Input fields

- input ([DeleteTeamDiscussionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### deleteTeamDiscussionComment

Deletes a team discussion comment.

#### Input fields

- input ([DeleteTeamDiscussionCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### deleteVerifiableDomain

Deletes a verifiable domain.

#### Input fields

- input ([DeleteVerifiableDomainInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| owner ([VerifiableDomainOwner](http://example.com)) | The owning account from which the domain was deleted. |

---

### disablePullRequestAutoMerge

Disable auto merge on the given pull request

#### Input fields

- input ([DisablePullRequestAutoMergeInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The pull request auto merge was disabled on. |

---

### dismissPullRequestReview

Dismisses an approved or rejected pull request review.

#### Input fields

- input ([DismissPullRequestReviewInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequestReview ([PullRequestReview](http://example.com)) | The dismissed pull request review. |

---

### enablePullRequestAutoMerge

Enable the default auto-merge on a pull request.

#### Input fields

- input ([EnablePullRequestAutoMergeInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The pull request auto-merge was enabled on. |

---

### followUser

Follow a user.

#### Input fields

- input ([FollowUserInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| user ([User](http://example.com)) | The user that was followed. |

---

### importProject

Creates a new project by importing columns and a list of issues/PRs.

#### Input fields

- input ([ImportProjectInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| project ([Project](http://example.com)) | The new Project! |

---

### inviteEnterpriseAdmin

Invite someone to become an administrator of the enterprise.

#### Input fields

- input ([InviteEnterpriseAdminInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| invitation ([EnterpriseAdministratorInvitation](http://example.com)) | The created enterprise administrator invitation. |

---

### linkRepositoryToProject

Creates a repository link for a project.

#### Input fields

- input ([LinkRepositoryToProjectInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| project ([Project](http://example.com)) | The linked Project. |
| repository ([Repository](http://example.com)) | The linked Repository. |

---

### lockLockable

Lock a lockable object

#### Input fields

- input ([LockLockableInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| lockedRecord ([Lockable](http://example.com)) | The item that was locked. |

---

### markFileAsViewed

Mark a pull request file as viewed

#### Input fields

- input ([MarkFileAsViewedInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The updated pull request. |

---

### markPullRequestReadyForReview

Marks a pull request ready for review.

#### Input fields

- input ([MarkPullRequestReadyForReviewInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The pull request that is ready for review. |

---

### mergeBranch

Merge a head into a branch.

#### Input fields

- input ([MergeBranchInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| mergeCommit ([Commit](http://example.com)) | The resulting merge Commit. |

---

### mergePullRequest

Merge a pull request.

#### Input fields

- input ([MergePullRequestInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The pull request that was merged. |

---

### minimizeComment

Minimizes a comment on an Issue, Commit, Pull Request, or Gist

#### Input fields

- input ([MinimizeCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| minimizedComment ([Minimizable](http://example.com)) | The comment that was minimized. |

---

### moveProjectCard

Moves a project card to another place.

#### Input fields

- input ([MoveProjectCardInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| cardEdge ([ProjectCardEdge](http://example.com)) | The new edge of the moved card. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### moveProjectColumn

Moves a project column to another place.

#### Input fields

- input ([MoveProjectColumnInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| columnEdge ([ProjectColumnEdge](http://example.com)) | The new edge of the moved column. |

---

### pinIssue

Pin an issue to a repository

#### Input fields

- input ([PinIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issue ([Issue](http://example.com)) | The issue that was pinned |

---

### regenerateEnterpriseIdentityProviderRecoveryCodes

Regenerates the identity provider recovery codes for an enterprise

#### Input fields

- input ([RegenerateEnterpriseIdentityProviderRecoveryCodesInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| identityProvider ([EnterpriseIdentityProvider](http://example.com)) | The identity provider for the enterprise. |

---

### regenerateVerifiableDomainToken

Regenerates a verifiable domain's verification token.

#### Input fields

- input ([RegenerateVerifiableDomainTokenInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| verificationToken ([String](http://example.com)) | The verification token that was generated. |

---

### removeAssigneesFromAssignable

Removes assignees from an assignable object.

#### Input fields

- input ([RemoveAssigneesFromAssignableInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| assignable ([Assignable](http://example.com)) | The item that was unassigned. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### removeEnterpriseAdmin

Removes an administrator from the enterprise.

#### Input fields

- input ([RemoveEnterpriseAdminInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| admin ([User](http://example.com)) | The user who was removed as an administrator. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The updated enterprise. |
| message ([String](http://example.com)) | A message confirming the result of removing an administrator. |
| viewer ([User](http://example.com)) | The viewer performing the mutation. |

---

### removeEnterpriseIdentityProvider

Removes the identity provider from an enterprise

#### Input fields

- input ([RemoveEnterpriseIdentityProviderInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| identityProvider ([EnterpriseIdentityProvider](http://example.com)) | The identity provider that was removed from the enterprise. |

---

### removeEnterpriseOrganization

Removes an organization from the enterprise

#### Input fields

- input ([RemoveEnterpriseOrganizationInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The updated enterprise. |
| organization ([Organization](http://example.com)) | The organization that was removed from the enterprise. |
| viewer ([User](http://example.com)) | The viewer performing the mutation. |

---

### removeEnterpriseSupportEntitlement

Removes a support entitlement from an enterprise member.

#### Input fields

- input ([RemoveEnterpriseSupportEntitlementInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| message ([String](http://example.com)) | A message confirming the result of removing the support entitlement. |

---

### removeLabelsFromLabelable

Removes labels from a Labelable object.

#### Input fields

- input ([RemoveLabelsFromLabelableInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| labelable ([Labelable](http://example.com)) | The Labelable the labels were removed from. |

---

### removeOutsideCollaborator

Removes outside collaborator from all repositories in an organization.

#### Input fields

- input ([RemoveOutsideCollaboratorInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| removedUser ([User](http://example.com)) | The user that was removed as an outside collaborator. |

---

### removeReaction

Removes a reaction from a subject.

#### Input fields

- input ([RemoveReactionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| reaction ([Reaction](http://example.com)) | The reaction object. |
| subject ([Reactable](http://example.com)) | The reactable subject. |

---

### removeStar

Removes a star from a Starrable.

#### Input fields

- input ([RemoveStarInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| starrable ([Starrable](http://example.com)) | The starrable. |

---

### reopenIssue

Reopen a issue.

#### Input fields

- input ([ReopenIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issue ([Issue](http://example.com)) | The issue that was opened. |

---

### reopenPullRequest

Reopen a pull request.

#### Input fields

- input ([ReopenPullRequestInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The pull request that was reopened. |

---

### requestReviews

Set review requests on a pull request.

#### Input fields

- input ([RequestReviewsInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The pull request that is getting requests. |
| requestedReviewersEdge ([UserEdge](http://example.com)) | The edge from the pull request to the requested reviewers. |

---

### rerequestCheckSuite

Rerequests an existing check suite.

#### Input fields

- input ([RerequestCheckSuiteInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| checkSuite ([CheckSuite](http://example.com)) | The requested check suite. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### resolveReviewThread

Marks a review thread as resolved.

#### Input fields

- input ([ResolveReviewThreadInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| thread ([PullRequestReviewThread](http://example.com)) | The thread to resolve. |

---

### setEnterpriseIdentityProvider

Creates or updates the identity provider for an enterprise.

#### Input fields

- input ([SetEnterpriseIdentityProviderInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| identityProvider ([EnterpriseIdentityProvider](http://example.com)) | The identity provider for the enterprise. |

---

### setOrganizationInteractionLimit

Set an organization level interaction limit for an organization's public repositories.

#### Input fields

- input ([SetOrganizationInteractionLimitInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| organization ([Organization](http://example.com)) | The organization that the interaction limit was set for. |

---

### setRepositoryInteractionLimit

Sets an interaction limit setting for a repository.

#### Input fields

- input ([SetRepositoryInteractionLimitInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The repository that the interaction limit was set for. |

---

### setUserInteractionLimit

Set a user level interaction limit for an user's public repositories.

#### Input fields

- input ([SetUserInteractionLimitInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| user ([User](http://example.com)) | The user that the interaction limit was set for. |

---

### submitPullRequestReview

Submits a pending pull request review.

#### Input fields

- input ([SubmitPullRequestReviewInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequestReview ([PullRequestReview](http://example.com)) | The submitted pull request review. |

---

### transferIssue

Transfer an issue to a different repository

#### Input fields

- input ([TransferIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issue ([Issue](http://example.com)) | The issue that was transferred |

---

### unarchiveRepository

Unarchives a repository.

#### Input fields

- input ([UnarchiveRepositoryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The repository that was unarchived. |

---

### unfollowUser

Unfollow a user.

#### Input fields

- input ([UnfollowUserInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| user ([User](http://example.com)) | The user that was unfollowed. |

---

### unlinkRepositoryFromProject

Deletes a repository link from a project.

#### Input fields

- input ([UnlinkRepositoryFromProjectInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| project ([Project](http://example.com)) | The linked Project. |
| repository ([Repository](http://example.com)) | The linked Repository. |

---

### unlockLockable

Unlock a lockable object

#### Input fields

- input ([UnlockLockableInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| unlockedRecord ([Lockable](http://example.com)) | The item that was unlocked. |

---

### unmarkFileAsViewed

Unmark a pull request file as viewed

#### Input fields

- input ([UnmarkFileAsViewedInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The updated pull request. |

---

### unmarkIssueAsDuplicate

Unmark an issue as a duplicate of another issue.

#### Input fields

- input ([UnmarkIssueAsDuplicateInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| duplicate ([IssueOrPullRequest](http://example.com)) | The issue or pull request that was marked as a duplicate. |

---

### unminimizeComment

Unminimizes a comment on an Issue, Commit, Pull Request, or Gist

#### Input fields

- input ([UnminimizeCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| unminimizedComment ([Minimizable](http://example.com)) | The comment that was unminimized. |

---

### unpinIssue

Unpin a pinned issue from a repository

#### Input fields

- input ([UnpinIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issue ([Issue](http://example.com)) | The issue that was unpinned |

---

### unresolveReviewThread

Marks a review thread as unresolved.

#### Input fields

- input ([UnresolveReviewThreadInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| thread ([PullRequestReviewThread](http://example.com)) | The thread to resolve. |

---

### updateBranchProtectionRule

Create a new branch protection rule

#### Input fields

- input ([UpdateBranchProtectionRuleInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| branchProtectionRule ([BranchProtectionRule](http://example.com)) | The newly created BranchProtectionRule. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### updateCheckRun

Update a check run

#### Input fields

- input ([UpdateCheckRunInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| checkRun ([CheckRun](http://example.com)) | The updated check run. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### updateCheckSuitePreferences

Modifies the settings of an existing check suite

#### Input fields

- input ([UpdateCheckSuitePreferencesInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The updated repository. |

---

### updateEnterpriseAdministratorRole

Updates the role of an enterprise administrator.

#### Input fields

- input ([UpdateEnterpriseAdministratorRoleInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| message ([String](http://example.com)) | A message confirming the result of changing the administrator's role. |

---

### updateEnterpriseAllowPrivateRepositoryForkingSetting

Sets whether private repository forks are enabled for an enterprise.

#### Input fields

- input ([UpdateEnterpriseAllowPrivateRepositoryForkingSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated allow private repository forking setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the allow private repository forking setting. |

---

### updateEnterpriseDefaultRepositoryPermissionSetting

Sets the default repository permission for organizations in an enterprise.

#### Input fields

- input ([UpdateEnterpriseDefaultRepositoryPermissionSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated default repository permission setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the default repository permission setting. |

---

### updateEnterpriseMembersCanChangeRepositoryVisibilitySetting

Sets whether organization members with admin permissions on a repository can change repository visibility.

#### Input fields

- input ([UpdateEnterpriseMembersCanChangeRepositoryVisibilitySettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can change repository visibility setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can change repository visibility setting. |

---

### updateEnterpriseMembersCanCreateRepositoriesSetting

Sets the members can create repositories setting for an enterprise.

#### Input fields

- input ([UpdateEnterpriseMembersCanCreateRepositoriesSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can create repositories setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can create repositories setting. |

---

### updateEnterpriseMembersCanDeleteIssuesSetting

Sets the members can delete issues setting for an enterprise.

#### Input fields

- input ([UpdateEnterpriseMembersCanDeleteIssuesSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can delete issues setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can delete issues setting. |

---

### updateEnterpriseMembersCanDeleteRepositoriesSetting

Sets the members can delete repositories setting for an enterprise.

#### Input fields

- input ([UpdateEnterpriseMembersCanDeleteRepositoriesSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can delete repositories setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can delete repositories setting. |

---

### updateEnterpriseMembersCanInviteCollaboratorsSetting

Sets whether members can invite collaborators are enabled for an enterprise.

#### Input fields

- input ([UpdateEnterpriseMembersCanInviteCollaboratorsSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can invite collaborators setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can invite collaborators setting. |

---

### updateEnterpriseMembersCanMakePurchasesSetting

Sets whether or not an organization admin can make purchases.

#### Input fields

- input ([UpdateEnterpriseMembersCanMakePurchasesSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can make purchases setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can make purchases setting. |

---

### updateEnterpriseMembersCanUpdateProtectedBranchesSetting

Sets the members can update protected branches setting for an enterprise.

#### Input fields

- input ([UpdateEnterpriseMembersCanUpdateProtectedBranchesSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can update protected branches setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can update protected branches setting. |

---

### updateEnterpriseMembersCanViewDependencyInsightsSetting

Sets the members can view dependency insights for an enterprise.

#### Input fields

- input ([UpdateEnterpriseMembersCanViewDependencyInsightsSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated members can view dependency insights setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the members can view dependency insights setting. |

---

### updateEnterpriseOrganizationProjectsSetting

Sets whether organization projects are enabled for an enterprise.

#### Input fields

- input ([UpdateEnterpriseOrganizationProjectsSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated organization projects setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the organization projects setting. |

---

### updateEnterpriseProfile

Updates an enterprise's profile.

#### Input fields

- input ([UpdateEnterpriseProfileInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The updated enterprise. |

---

### updateEnterpriseRepositoryProjectsSetting

Sets whether repository projects are enabled for a enterprise.

#### Input fields

- input ([UpdateEnterpriseRepositoryProjectsSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated repository projects setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the repository projects setting. |

---

### updateEnterpriseTeamDiscussionsSetting

Sets whether team discussions are enabled for an enterprise.

#### Input fields

- input ([UpdateEnterpriseTeamDiscussionsSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated team discussions setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the team discussions setting. |

---

### updateEnterpriseTwoFactorAuthenticationRequiredSetting

Sets whether two factor authentication is required for all users in an enterprise.

#### Input fields

- input ([UpdateEnterpriseTwoFactorAuthenticationRequiredSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| enterprise ([Enterprise](http://example.com)) | The enterprise with the updated two factor authentication required setting. |
| message ([String](http://example.com)) | A message confirming the result of updating the two factor authentication required setting. |

---

### updateIpAllowListEnabledSetting

Sets whether an IP allow list is enabled on an owner.

#### Input fields

- input ([UpdateIpAllowListEnabledSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| owner ([IpAllowListOwner](http://example.com)) | The IP allow list owner on which the setting was updated. |

---

### updateIpAllowListEntry

Updates an IP allow list entry.

#### Input fields

- input ([UpdateIpAllowListEntryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| ipAllowListEntry ([IpAllowListEntry](http://example.com)) | The IP allow list entry that was updated. |

---

### updateIssue

Updates an Issue.

#### Input fields

- input ([UpdateIssueInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issue ([Issue](http://example.com)) | The issue. |

---

### updateIssueComment

Updates an IssueComment object.

#### Input fields

- input ([UpdateIssueCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| issueComment ([IssueComment](http://example.com)) | The updated comment. |

---

### updateLabel

Updates an existing label.

#### Input fields

- input ([UpdateLabelInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| label ([Label](http://example.com)) | The updated label. |

---

### updateNotificationRestrictionSetting

Update the setting to restrict notifications to only verified domains available to an owner.

#### Input fields

- input ([UpdateNotificationRestrictionSettingInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| owner ([VerifiableDomainOwner](http://example.com)) | The owner on which the setting was updated. |

---

### updateProject

Updates an existing project.

#### Input fields

- input ([UpdateProjectInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| project ([Project](http://example.com)) | The updated project. |

---

### updateProjectCard

Updates an existing project card.

#### Input fields

- input ([UpdateProjectCardInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| projectCard ([ProjectCard](http://example.com)) | The updated ProjectCard. |

---

### updateProjectColumn

Updates an existing project column.

#### Input fields

- input ([UpdateProjectColumnInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| projectColumn ([ProjectColumn](http://example.com)) | The updated project column. |

---

### updatePullRequest

Update a pull request

#### Input fields

- input ([UpdatePullRequestInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](http://example.com)) | Identifies the actor who performed the event. |
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequest ([PullRequest](http://example.com)) | The updated pull request. |

---

### updatePullRequestReview

Updates the body of a pull request review.

#### Input fields

- input ([UpdatePullRequestReviewInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequestReview ([PullRequestReview](http://example.com)) | The updated pull request review. |

---

### updatePullRequestReviewComment

Updates a pull request review comment.

#### Input fields

- input ([UpdatePullRequestReviewCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| pullRequestReviewComment ([PullRequestReviewComment](http://example.com)) | The updated comment. |

---

### updateRef

Update a Git Ref.

#### Input fields

- input ([UpdateRefInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| ref ([Ref](http://example.com)) | The updated Ref. |

---

### updateRefs

Creates, updates and/or deletes multiple refs in a repository.  This mutation takes a list of `RefUpdate`s and performs these updates on the repository. All updates are performed atomically, meaning that if one of them is rejected, no other ref will be modified.  `RefUpdate.beforeOid` specifies that the given reference needs to point to the given value before performing any updates. A value of `0000000000000000000000000000000000000000` can be used to verify that the references should not exist.  `RefUpdate.afterOid` specifies the value that the given reference will point to after performing all updates. A value of `0000000000000000000000000000000000000000` can be used to delete a reference.  If `RefUpdate.force` is set to `true`, a non-fast-forward updates for the given reference will be allowed.

#### Input fields

- input ([UpdateRefsInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |

---

### updateRepository

Update information about a repository.

#### Input fields

- input ([UpdateRepositoryInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| repository ([Repository](http://example.com)) | The updated repository. |

---

### updateSubscription

Updates the state for subscribable subjects.

#### Input fields

- input ([UpdateSubscriptionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| subscribable ([Subscribable](http://example.com)) | The input subscribable entity. |

---

### updateTeamDiscussion

Updates a team discussion.

#### Input fields

- input ([UpdateTeamDiscussionInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| teamDiscussion ([TeamDiscussion](http://example.com)) | The updated discussion. |

---

### updateTeamDiscussionComment

Updates a discussion comment.

#### Input fields

- input ([UpdateTeamDiscussionCommentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| teamDiscussionComment ([TeamDiscussionComment](http://example.com)) | The updated comment. |

---

### updateTeamReviewAssignment

Updates team review assignment.

#### Input fields

- input ([UpdateTeamReviewAssignmentInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| team ([Team](http://example.com)) | The team that was modified |

---

### updateTopics

Replaces the repository's topics with the given topics.

#### Input fields

- input ([UpdateTopicsInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| invalidTopicNames ([[String!]](http://example.com)) | Names of the provided topics that are not valid. |
| repository ([Repository](http://example.com)) | The updated repository. |

---

### verifyVerifiableDomain

Verify that a verifiable domain has the expected DNS record.

#### Input fields

- input ([VerifyVerifiableDomainInput!](http://example.com))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](http://example.com)) | A unique identifier for the client performing the mutation. |
| domain ([VerifiableDomain](http://example.com)) | The verifiable domain that was verified. |

---