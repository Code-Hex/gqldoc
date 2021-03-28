# Mutations

### About mutations

The root query for implementing GraphQL mutations.

### acceptEnterpriseAdministratorInvitation

<p>Accepts a pending invitation for a user to become an administrator of an enterprise.</p>

#### Input fields

- input ([AcceptEnterpriseAdministratorInvitationInput!](input_objects.md#acceptenterpriseadministratorinvitationinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| invitation ([EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)) | <p>The invitation that was accepted.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of accepting an administrator invitation.</p> |

---

### acceptTopicSuggestion

<p>Applies a suggested topic to the repository.</p>

#### Input fields

- input ([AcceptTopicSuggestionInput!](input_objects.md#accepttopicsuggestioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| topic ([Topic](objects.md#topic)) | <p>The accepted topic.</p> |

---

### addAssigneesToAssignable

<p>Adds assignees to an assignable object.</p>

#### Input fields

- input ([AddAssigneesToAssignableInput!](input_objects.md#addassigneestoassignableinput))
 

#### Returns

| Name | Description |
|------|-------------|
| assignable ([Assignable](interfaces.md#assignable)) | <p>The item that was assigned.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### addComment

<p>Adds a comment to an Issue or Pull Request.</p>

#### Input fields

- input ([AddCommentInput!](input_objects.md#addcommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| commentEdge ([IssueCommentEdge](objects.md#issuecommentedge)) | <p>The edge from the subject&rsquo;s comment connection.</p> |
| subject ([Node](interfaces.md#node)) | <p>The subject</p> |
| timelineEdge ([IssueTimelineItemEdge](objects.md#issuetimelineitemedge)) | <p>The edge from the subject&rsquo;s timeline connection.</p> |

---

### addEnterpriseSupportEntitlement

<p>Adds a support entitlement to an enterprise member.</p>

#### Input fields

- input ([AddEnterpriseSupportEntitlementInput!](input_objects.md#addenterprisesupportentitlementinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of adding the support entitlement.</p> |

---

### addLabelsToLabelable

<p>Adds labels to a labelable object.</p>

#### Input fields

- input ([AddLabelsToLabelableInput!](input_objects.md#addlabelstolabelableinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| labelable ([Labelable](interfaces.md#labelable)) | <p>The item that was labeled.</p> |

---

### addProjectCard

<p>Adds a card to a ProjectColumn. Either <code>contentId</code> or <code>note</code> must be provided but <strong>not</strong> both.</p>

#### Input fields

- input ([AddProjectCardInput!](input_objects.md#addprojectcardinput))
 

#### Returns

| Name | Description |
|------|-------------|
| cardEdge ([ProjectCardEdge](objects.md#projectcardedge)) | <p>The edge from the ProjectColumn&rsquo;s card connection.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| projectColumn ([ProjectColumn](objects.md#projectcolumn)) | <p>The ProjectColumn</p> |

---

### addProjectColumn

<p>Adds a column to a Project.</p>

#### Input fields

- input ([AddProjectColumnInput!](input_objects.md#addprojectcolumninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| columnEdge ([ProjectColumnEdge](objects.md#projectcolumnedge)) | <p>The edge from the project&rsquo;s column connection.</p> |
| project ([Project](objects.md#project)) | <p>The project</p> |

---

### addPullRequestReview

<p>Adds a review to a Pull Request.</p>

#### Input fields

- input ([AddPullRequestReviewInput!](input_objects.md#addpullrequestreviewinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequestReview ([PullRequestReview](objects.md#pullrequestreview)) | <p>The newly created pull request review.</p> |
| reviewEdge ([PullRequestReviewEdge](objects.md#pullrequestreviewedge)) | <p>The edge from the pull request&rsquo;s review connection.</p> |

---

### addPullRequestReviewComment

<p>Adds a comment to a review.</p>

#### Input fields

- input ([AddPullRequestReviewCommentInput!](input_objects.md#addpullrequestreviewcommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| comment ([PullRequestReviewComment](objects.md#pullrequestreviewcomment)) | <p>The newly created comment.</p> |
| commentEdge ([PullRequestReviewCommentEdge](objects.md#pullrequestreviewcommentedge)) | <p>The edge from the review&rsquo;s comment connection.</p> |

---

### addPullRequestReviewThread

<p>Adds a new thread to a pending Pull Request Review.</p>

#### Input fields

- input ([AddPullRequestReviewThreadInput!](input_objects.md#addpullrequestreviewthreadinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| thread ([PullRequestReviewThread](objects.md#pullrequestreviewthread)) | <p>The newly created thread.</p> |

---

### addReaction

<p>Adds a reaction to a subject.</p>

#### Input fields

- input ([AddReactionInput!](input_objects.md#addreactioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| reaction ([Reaction](objects.md#reaction)) | <p>The reaction object.</p> |
| subject ([Reactable](interfaces.md#reactable)) | <p>The reactable subject.</p> |

---

### addStar

<p>Adds a star to a Starrable.</p>

#### Input fields

- input ([AddStarInput!](input_objects.md#addstarinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| starrable ([Starrable](interfaces.md#starrable)) | <p>The starrable.</p> |

---

### addVerifiableDomain

<p>Adds a verifiable domain to an owning account.</p>

#### Input fields

- input ([AddVerifiableDomainInput!](input_objects.md#addverifiabledomaininput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| domain ([VerifiableDomain](objects.md#verifiabledomain)) | <p>The verifiable domain that was added.</p> |

---

### approveVerifiableDomain

<p>Approve a verifiable domain for notification delivery.</p>

#### Input fields

- input ([ApproveVerifiableDomainInput!](input_objects.md#approveverifiabledomaininput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| domain ([VerifiableDomain](objects.md#verifiabledomain)) | <p>The verifiable domain that was approved.</p> |

---

### archiveRepository

<p>Marks a repository as archived.</p>

#### Input fields

- input ([ArchiveRepositoryInput!](input_objects.md#archiverepositoryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The repository that was marked as archived.</p> |

---

### cancelEnterpriseAdminInvitation

<p>Cancels a pending invitation for an administrator to join an enterprise.</p>

#### Input fields

- input ([CancelEnterpriseAdminInvitationInput!](input_objects.md#cancelenterpriseadmininvitationinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| invitation ([EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)) | <p>The invitation that was canceled.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of canceling an administrator invitation.</p> |

---

### changeUserStatus

<p>Update your status on GitHub.</p>

#### Input fields

- input ([ChangeUserStatusInput!](input_objects.md#changeuserstatusinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| status ([UserStatus](objects.md#userstatus)) | <p>Your updated status.</p> |

---

### clearLabelsFromLabelable

<p>Clears all labels from a labelable object.</p>

#### Input fields

- input ([ClearLabelsFromLabelableInput!](input_objects.md#clearlabelsfromlabelableinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| labelable ([Labelable](interfaces.md#labelable)) | <p>The item that was unlabeled.</p> |

---

### cloneProject

<p>Creates a new project by cloning configuration from an existing project.</p>

#### Input fields

- input ([CloneProjectInput!](input_objects.md#cloneprojectinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| jobStatusId ([String](scalars.md#string)) | <p>The id of the JobStatus for populating cloned fields.</p> |
| project ([Project](objects.md#project)) | <p>The new cloned project.</p> |

---

### cloneTemplateRepository

<p>Create a new repository with the same files and directory structure as a template repository.</p>

#### Input fields

- input ([CloneTemplateRepositoryInput!](input_objects.md#clonetemplaterepositoryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The new repository.</p> |

---

### closeIssue

<p>Close an issue.</p>

#### Input fields

- input ([CloseIssueInput!](input_objects.md#closeissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issue ([Issue](objects.md#issue)) | <p>The issue that was closed.</p> |

---

### closePullRequest

<p>Close a pull request.</p>

#### Input fields

- input ([ClosePullRequestInput!](input_objects.md#closepullrequestinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The pull request that was closed.</p> |

---

### convertProjectCardNoteToIssue

<p>Convert a project note card to one associated with a newly created issue.</p>

#### Input fields

- input ([ConvertProjectCardNoteToIssueInput!](input_objects.md#convertprojectcardnotetoissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| projectCard ([ProjectCard](objects.md#projectcard)) | <p>The updated ProjectCard.</p> |

---

### createBranchProtectionRule

<p>Create a new branch protection rule</p>

#### Input fields

- input ([CreateBranchProtectionRuleInput!](input_objects.md#createbranchprotectionruleinput))
 

#### Returns

| Name | Description |
|------|-------------|
| branchProtectionRule ([BranchProtectionRule](objects.md#branchprotectionrule)) | <p>The newly created BranchProtectionRule.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### createCheckRun

<p>Create a check run.</p>

#### Input fields

- input ([CreateCheckRunInput!](input_objects.md#createcheckruninput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkRun ([CheckRun](objects.md#checkrun)) | <p>The newly created check run.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### createCheckSuite

<p>Create a check suite</p>

#### Input fields

- input ([CreateCheckSuiteInput!](input_objects.md#createchecksuiteinput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkSuite ([CheckSuite](objects.md#checksuite)) | <p>The newly created check suite.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### createContentAttachment

<p>Create a content attachment.</p>

#### Input fields

- input ([CreateContentAttachmentInput!](input_objects.md#createcontentattachmentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| contentAttachment ([ContentAttachment](objects.md#contentattachment)) | <p>The newly created content attachment.</p> |

---

### createDeployment

<p>Creates a new deployment event.</p>

#### Input fields

- input ([CreateDeploymentInput!](input_objects.md#createdeploymentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| autoMerged ([Boolean](scalars.md#boolean)) | <p>True if the default branch has been auto-merged into the deployment ref.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| deployment ([Deployment](objects.md#deployment)) | <p>The new deployment.</p> |

---

### createDeploymentStatus

<p>Create a deployment status.</p>

#### Input fields

- input ([CreateDeploymentStatusInput!](input_objects.md#createdeploymentstatusinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| deploymentStatus ([DeploymentStatus](objects.md#deploymentstatus)) | <p>The new deployment status.</p> |

---

### createEnterpriseOrganization

<p>Creates an organization as part of an enterprise account.</p>

#### Input fields

- input ([CreateEnterpriseOrganizationInput!](input_objects.md#createenterpriseorganizationinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise that owns the created organization.</p> |
| organization ([Organization](objects.md#organization)) | <p>The organization that was created.</p> |

---

### createIpAllowListEntry

<p>Creates a new IP allow list entry.</p>

#### Input fields

- input ([CreateIpAllowListEntryInput!](input_objects.md#createipallowlistentryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| ipAllowListEntry ([IpAllowListEntry](objects.md#ipallowlistentry)) | <p>The IP allow list entry that was created.</p> |

---

### createIssue

<p>Creates a new issue.</p>

#### Input fields

- input ([CreateIssueInput!](input_objects.md#createissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issue ([Issue](objects.md#issue)) | <p>The new issue.</p> |

---

### createLabel

<p>Creates a new label.</p>

#### Input fields

- input ([CreateLabelInput!](input_objects.md#createlabelinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| label ([Label](objects.md#label)) | <p>The new label.</p> |

---

### createProject

<p>Creates a new project.</p>

#### Input fields

- input ([CreateProjectInput!](input_objects.md#createprojectinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| project ([Project](objects.md#project)) | <p>The new project.</p> |

---

### createPullRequest

<p>Create a new pull request</p>

#### Input fields

- input ([CreatePullRequestInput!](input_objects.md#createpullrequestinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The new pull request.</p> |

---

### createRef

<p>Create a new Git Ref.</p>

#### Input fields

- input ([CreateRefInput!](input_objects.md#createrefinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| ref ([Ref](objects.md#ref)) | <p>The newly created ref.</p> |

---

### createRepository

<p>Create a new repository.</p>

#### Input fields

- input ([CreateRepositoryInput!](input_objects.md#createrepositoryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The new repository.</p> |

---

### createTeamDiscussion

<p>Creates a new team discussion.</p>

#### Input fields

- input ([CreateTeamDiscussionInput!](input_objects.md#createteamdiscussioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| teamDiscussion ([TeamDiscussion](objects.md#teamdiscussion)) | <p>The new discussion.</p> |

---

### createTeamDiscussionComment

<p>Creates a new team discussion comment.</p>

#### Input fields

- input ([CreateTeamDiscussionCommentInput!](input_objects.md#createteamdiscussioncommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| teamDiscussionComment ([TeamDiscussionComment](objects.md#teamdiscussioncomment)) | <p>The new comment.</p> |

---

### declineTopicSuggestion

<p>Rejects a suggested topic for the repository.</p>

#### Input fields

- input ([DeclineTopicSuggestionInput!](input_objects.md#declinetopicsuggestioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| topic ([Topic](objects.md#topic)) | <p>The declined topic.</p> |

---

### deleteBranchProtectionRule

<p>Delete a branch protection rule</p>

#### Input fields

- input ([DeleteBranchProtectionRuleInput!](input_objects.md#deletebranchprotectionruleinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### deleteDeployment

<p>Deletes a deployment.</p>

#### Input fields

- input ([DeleteDeploymentInput!](input_objects.md#deletedeploymentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### deleteIpAllowListEntry

<p>Deletes an IP allow list entry.</p>

#### Input fields

- input ([DeleteIpAllowListEntryInput!](input_objects.md#deleteipallowlistentryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| ipAllowListEntry ([IpAllowListEntry](objects.md#ipallowlistentry)) | <p>The IP allow list entry that was deleted.</p> |

---

### deleteIssue

<p>Deletes an Issue object.</p>

#### Input fields

- input ([DeleteIssueInput!](input_objects.md#deleteissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The repository the issue belonged to</p> |

---

### deleteIssueComment

<p>Deletes an IssueComment object.</p>

#### Input fields

- input ([DeleteIssueCommentInput!](input_objects.md#deleteissuecommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### deleteLabel

<p>Deletes a label.</p>

#### Input fields

- input ([DeleteLabelInput!](input_objects.md#deletelabelinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### deletePackageVersion

<p>Delete a package version.</p>

#### Input fields

- input ([DeletePackageVersionInput!](input_objects.md#deletepackageversioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| success ([Boolean](scalars.md#boolean)) | <p>Whether or not the operation succeeded.</p> |

---

### deleteProject

<p>Deletes a project.</p>

#### Input fields

- input ([DeleteProjectInput!](input_objects.md#deleteprojectinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| owner ([ProjectOwner](interfaces.md#projectowner)) | <p>The repository or organization the project was removed from.</p> |

---

### deleteProjectCard

<p>Deletes a project card.</p>

#### Input fields

- input ([DeleteProjectCardInput!](input_objects.md#deleteprojectcardinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| column ([ProjectColumn](objects.md#projectcolumn)) | <p>The column the deleted card was in.</p> |
| deletedCardId ([ID](scalars.md#id)) | <p>The deleted card ID.</p> |

---

### deleteProjectColumn

<p>Deletes a project column.</p>

#### Input fields

- input ([DeleteProjectColumnInput!](input_objects.md#deleteprojectcolumninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| deletedColumnId ([ID](scalars.md#id)) | <p>The deleted column ID.</p> |
| project ([Project](objects.md#project)) | <p>The project the deleted column was in.</p> |

---

### deletePullRequestReview

<p>Deletes a pull request review.</p>

#### Input fields

- input ([DeletePullRequestReviewInput!](input_objects.md#deletepullrequestreviewinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequestReview ([PullRequestReview](objects.md#pullrequestreview)) | <p>The deleted pull request review.</p> |

---

### deletePullRequestReviewComment

<p>Deletes a pull request review comment.</p>

#### Input fields

- input ([DeletePullRequestReviewCommentInput!](input_objects.md#deletepullrequestreviewcommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequestReview ([PullRequestReview](objects.md#pullrequestreview)) | <p>The pull request review the deleted comment belonged to.</p> |

---

### deleteRef

<p>Delete a Git Ref.</p>

#### Input fields

- input ([DeleteRefInput!](input_objects.md#deleterefinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### deleteTeamDiscussion

<p>Deletes a team discussion.</p>

#### Input fields

- input ([DeleteTeamDiscussionInput!](input_objects.md#deleteteamdiscussioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### deleteTeamDiscussionComment

<p>Deletes a team discussion comment.</p>

#### Input fields

- input ([DeleteTeamDiscussionCommentInput!](input_objects.md#deleteteamdiscussioncommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### deleteVerifiableDomain

<p>Deletes a verifiable domain.</p>

#### Input fields

- input ([DeleteVerifiableDomainInput!](input_objects.md#deleteverifiabledomaininput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| owner ([VerifiableDomainOwner](unions.md#verifiabledomainowner)) | <p>The owning account from which the domain was deleted.</p> |

---

### disablePullRequestAutoMerge

<p>Disable auto merge on the given pull request</p>

#### Input fields

- input ([DisablePullRequestAutoMergeInput!](input_objects.md#disablepullrequestautomergeinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The pull request auto merge was disabled on.</p> |

---

### dismissPullRequestReview

<p>Dismisses an approved or rejected pull request review.</p>

#### Input fields

- input ([DismissPullRequestReviewInput!](input_objects.md#dismisspullrequestreviewinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequestReview ([PullRequestReview](objects.md#pullrequestreview)) | <p>The dismissed pull request review.</p> |

---

### enablePullRequestAutoMerge

<p>Enable the default auto-merge on a pull request.</p>

#### Input fields

- input ([EnablePullRequestAutoMergeInput!](input_objects.md#enablepullrequestautomergeinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The pull request auto-merge was enabled on.</p> |

---

### followUser

<p>Follow a user.</p>

#### Input fields

- input ([FollowUserInput!](input_objects.md#followuserinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| user ([User](objects.md#user)) | <p>The user that was followed.</p> |

---

### importProject

<p>Creates a new project by importing columns and a list of issues/PRs.</p>

#### Input fields

- input ([ImportProjectInput!](input_objects.md#importprojectinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| project ([Project](objects.md#project)) | <p>The new Project!</p> |

---

### inviteEnterpriseAdmin

<p>Invite someone to become an administrator of the enterprise.</p>

#### Input fields

- input ([InviteEnterpriseAdminInput!](input_objects.md#inviteenterpriseadmininput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| invitation ([EnterpriseAdministratorInvitation](objects.md#enterpriseadministratorinvitation)) | <p>The created enterprise administrator invitation.</p> |

---

### linkRepositoryToProject

<p>Creates a repository link for a project.</p>

#### Input fields

- input ([LinkRepositoryToProjectInput!](input_objects.md#linkrepositorytoprojectinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| project ([Project](objects.md#project)) | <p>The linked Project.</p> |
| repository ([Repository](objects.md#repository)) | <p>The linked Repository.</p> |

---

### lockLockable

<p>Lock a lockable object</p>

#### Input fields

- input ([LockLockableInput!](input_objects.md#locklockableinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| lockedRecord ([Lockable](interfaces.md#lockable)) | <p>The item that was locked.</p> |

---

### markFileAsViewed

<p>Mark a pull request file as viewed</p>

#### Input fields

- input ([MarkFileAsViewedInput!](input_objects.md#markfileasviewedinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The updated pull request.</p> |

---

### markPullRequestReadyForReview

<p>Marks a pull request ready for review.</p>

#### Input fields

- input ([MarkPullRequestReadyForReviewInput!](input_objects.md#markpullrequestreadyforreviewinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The pull request that is ready for review.</p> |

---

### mergeBranch

<p>Merge a head into a branch.</p>

#### Input fields

- input ([MergeBranchInput!](input_objects.md#mergebranchinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| mergeCommit ([Commit](objects.md#commit)) | <p>The resulting merge Commit.</p> |

---

### mergePullRequest

<p>Merge a pull request.</p>

#### Input fields

- input ([MergePullRequestInput!](input_objects.md#mergepullrequestinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The pull request that was merged.</p> |

---

### minimizeComment

<p>Minimizes a comment on an Issue, Commit, Pull Request, or Gist</p>

#### Input fields

- input ([MinimizeCommentInput!](input_objects.md#minimizecommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| minimizedComment ([Minimizable](interfaces.md#minimizable)) | <p>The comment that was minimized.</p> |

---

### moveProjectCard

<p>Moves a project card to another place.</p>

#### Input fields

- input ([MoveProjectCardInput!](input_objects.md#moveprojectcardinput))
 

#### Returns

| Name | Description |
|------|-------------|
| cardEdge ([ProjectCardEdge](objects.md#projectcardedge)) | <p>The new edge of the moved card.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### moveProjectColumn

<p>Moves a project column to another place.</p>

#### Input fields

- input ([MoveProjectColumnInput!](input_objects.md#moveprojectcolumninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| columnEdge ([ProjectColumnEdge](objects.md#projectcolumnedge)) | <p>The new edge of the moved column.</p> |

---

### pinIssue

<p>Pin an issue to a repository</p>

#### Input fields

- input ([PinIssueInput!](input_objects.md#pinissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issue ([Issue](objects.md#issue)) | <p>The issue that was pinned</p> |

---

### regenerateEnterpriseIdentityProviderRecoveryCodes

<p>Regenerates the identity provider recovery codes for an enterprise</p>

#### Input fields

- input ([RegenerateEnterpriseIdentityProviderRecoveryCodesInput!](input_objects.md#regenerateenterpriseidentityproviderrecoverycodesinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| identityProvider ([EnterpriseIdentityProvider](objects.md#enterpriseidentityprovider)) | <p>The identity provider for the enterprise.</p> |

---

### regenerateVerifiableDomainToken

<p>Regenerates a verifiable domain&rsquo;s verification token.</p>

#### Input fields

- input ([RegenerateVerifiableDomainTokenInput!](input_objects.md#regenerateverifiabledomaintokeninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| verificationToken ([String](scalars.md#string)) | <p>The verification token that was generated.</p> |

---

### removeAssigneesFromAssignable

<p>Removes assignees from an assignable object.</p>

#### Input fields

- input ([RemoveAssigneesFromAssignableInput!](input_objects.md#removeassigneesfromassignableinput))
 

#### Returns

| Name | Description |
|------|-------------|
| assignable ([Assignable](interfaces.md#assignable)) | <p>The item that was unassigned.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### removeEnterpriseAdmin

<p>Removes an administrator from the enterprise.</p>

#### Input fields

- input ([RemoveEnterpriseAdminInput!](input_objects.md#removeenterpriseadmininput))
 

#### Returns

| Name | Description |
|------|-------------|
| admin ([User](objects.md#user)) | <p>The user who was removed as an administrator.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The updated enterprise.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of removing an administrator.</p> |
| viewer ([User](objects.md#user)) | <p>The viewer performing the mutation.</p> |

---

### removeEnterpriseIdentityProvider

<p>Removes the identity provider from an enterprise</p>

#### Input fields

- input ([RemoveEnterpriseIdentityProviderInput!](input_objects.md#removeenterpriseidentityproviderinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| identityProvider ([EnterpriseIdentityProvider](objects.md#enterpriseidentityprovider)) | <p>The identity provider that was removed from the enterprise.</p> |

---

### removeEnterpriseOrganization

<p>Removes an organization from the enterprise</p>

#### Input fields

- input ([RemoveEnterpriseOrganizationInput!](input_objects.md#removeenterpriseorganizationinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The updated enterprise.</p> |
| organization ([Organization](objects.md#organization)) | <p>The organization that was removed from the enterprise.</p> |
| viewer ([User](objects.md#user)) | <p>The viewer performing the mutation.</p> |

---

### removeEnterpriseSupportEntitlement

<p>Removes a support entitlement from an enterprise member.</p>

#### Input fields

- input ([RemoveEnterpriseSupportEntitlementInput!](input_objects.md#removeenterprisesupportentitlementinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of removing the support entitlement.</p> |

---

### removeLabelsFromLabelable

<p>Removes labels from a Labelable object.</p>

#### Input fields

- input ([RemoveLabelsFromLabelableInput!](input_objects.md#removelabelsfromlabelableinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| labelable ([Labelable](interfaces.md#labelable)) | <p>The Labelable the labels were removed from.</p> |

---

### removeOutsideCollaborator

<p>Removes outside collaborator from all repositories in an organization.</p>

#### Input fields

- input ([RemoveOutsideCollaboratorInput!](input_objects.md#removeoutsidecollaboratorinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| removedUser ([User](objects.md#user)) | <p>The user that was removed as an outside collaborator.</p> |

---

### removeReaction

<p>Removes a reaction from a subject.</p>

#### Input fields

- input ([RemoveReactionInput!](input_objects.md#removereactioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| reaction ([Reaction](objects.md#reaction)) | <p>The reaction object.</p> |
| subject ([Reactable](interfaces.md#reactable)) | <p>The reactable subject.</p> |

---

### removeStar

<p>Removes a star from a Starrable.</p>

#### Input fields

- input ([RemoveStarInput!](input_objects.md#removestarinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| starrable ([Starrable](interfaces.md#starrable)) | <p>The starrable.</p> |

---

### reopenIssue

<p>Reopen a issue.</p>

#### Input fields

- input ([ReopenIssueInput!](input_objects.md#reopenissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issue ([Issue](objects.md#issue)) | <p>The issue that was opened.</p> |

---

### reopenPullRequest

<p>Reopen a pull request.</p>

#### Input fields

- input ([ReopenPullRequestInput!](input_objects.md#reopenpullrequestinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The pull request that was reopened.</p> |

---

### requestReviews

<p>Set review requests on a pull request.</p>

#### Input fields

- input ([RequestReviewsInput!](input_objects.md#requestreviewsinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The pull request that is getting requests.</p> |
| requestedReviewersEdge ([UserEdge](objects.md#useredge)) | <p>The edge from the pull request to the requested reviewers.</p> |

---

### rerequestCheckSuite

<p>Rerequests an existing check suite.</p>

#### Input fields

- input ([RerequestCheckSuiteInput!](input_objects.md#rerequestchecksuiteinput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkSuite ([CheckSuite](objects.md#checksuite)) | <p>The requested check suite.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### resolveReviewThread

<p>Marks a review thread as resolved.</p>

#### Input fields

- input ([ResolveReviewThreadInput!](input_objects.md#resolvereviewthreadinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| thread ([PullRequestReviewThread](objects.md#pullrequestreviewthread)) | <p>The thread to resolve.</p> |

---

### setEnterpriseIdentityProvider

<p>Creates or updates the identity provider for an enterprise.</p>

#### Input fields

- input ([SetEnterpriseIdentityProviderInput!](input_objects.md#setenterpriseidentityproviderinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| identityProvider ([EnterpriseIdentityProvider](objects.md#enterpriseidentityprovider)) | <p>The identity provider for the enterprise.</p> |

---

### setOrganizationInteractionLimit

<p>Set an organization level interaction limit for an organization&rsquo;s public repositories.</p>

#### Input fields

- input ([SetOrganizationInteractionLimitInput!](input_objects.md#setorganizationinteractionlimitinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| organization ([Organization](objects.md#organization)) | <p>The organization that the interaction limit was set for.</p> |

---

### setRepositoryInteractionLimit

<p>Sets an interaction limit setting for a repository.</p>

#### Input fields

- input ([SetRepositoryInteractionLimitInput!](input_objects.md#setrepositoryinteractionlimitinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The repository that the interaction limit was set for.</p> |

---

### setUserInteractionLimit

<p>Set a user level interaction limit for an user&rsquo;s public repositories.</p>

#### Input fields

- input ([SetUserInteractionLimitInput!](input_objects.md#setuserinteractionlimitinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| user ([User](objects.md#user)) | <p>The user that the interaction limit was set for.</p> |

---

### submitPullRequestReview

<p>Submits a pending pull request review.</p>

#### Input fields

- input ([SubmitPullRequestReviewInput!](input_objects.md#submitpullrequestreviewinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequestReview ([PullRequestReview](objects.md#pullrequestreview)) | <p>The submitted pull request review.</p> |

---

### transferIssue

<p>Transfer an issue to a different repository</p>

#### Input fields

- input ([TransferIssueInput!](input_objects.md#transferissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issue ([Issue](objects.md#issue)) | <p>The issue that was transferred</p> |

---

### unarchiveRepository

<p>Unarchives a repository.</p>

#### Input fields

- input ([UnarchiveRepositoryInput!](input_objects.md#unarchiverepositoryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The repository that was unarchived.</p> |

---

### unfollowUser

<p>Unfollow a user.</p>

#### Input fields

- input ([UnfollowUserInput!](input_objects.md#unfollowuserinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| user ([User](objects.md#user)) | <p>The user that was unfollowed.</p> |

---

### unlinkRepositoryFromProject

<p>Deletes a repository link from a project.</p>

#### Input fields

- input ([UnlinkRepositoryFromProjectInput!](input_objects.md#unlinkrepositoryfromprojectinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| project ([Project](objects.md#project)) | <p>The linked Project.</p> |
| repository ([Repository](objects.md#repository)) | <p>The linked Repository.</p> |

---

### unlockLockable

<p>Unlock a lockable object</p>

#### Input fields

- input ([UnlockLockableInput!](input_objects.md#unlocklockableinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| unlockedRecord ([Lockable](interfaces.md#lockable)) | <p>The item that was unlocked.</p> |

---

### unmarkFileAsViewed

<p>Unmark a pull request file as viewed</p>

#### Input fields

- input ([UnmarkFileAsViewedInput!](input_objects.md#unmarkfileasviewedinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The updated pull request.</p> |

---

### unmarkIssueAsDuplicate

<p>Unmark an issue as a duplicate of another issue.</p>

#### Input fields

- input ([UnmarkIssueAsDuplicateInput!](input_objects.md#unmarkissueasduplicateinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| duplicate ([IssueOrPullRequest](unions.md#issueorpullrequest)) | <p>The issue or pull request that was marked as a duplicate.</p> |

---

### unminimizeComment

<p>Unminimizes a comment on an Issue, Commit, Pull Request, or Gist</p>

#### Input fields

- input ([UnminimizeCommentInput!](input_objects.md#unminimizecommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| unminimizedComment ([Minimizable](interfaces.md#minimizable)) | <p>The comment that was unminimized.</p> |

---

### unpinIssue

<p>Unpin a pinned issue from a repository</p>

#### Input fields

- input ([UnpinIssueInput!](input_objects.md#unpinissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issue ([Issue](objects.md#issue)) | <p>The issue that was unpinned</p> |

---

### unresolveReviewThread

<p>Marks a review thread as unresolved.</p>

#### Input fields

- input ([UnresolveReviewThreadInput!](input_objects.md#unresolvereviewthreadinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| thread ([PullRequestReviewThread](objects.md#pullrequestreviewthread)) | <p>The thread to resolve.</p> |

---

### updateBranchProtectionRule

<p>Create a new branch protection rule</p>

#### Input fields

- input ([UpdateBranchProtectionRuleInput!](input_objects.md#updatebranchprotectionruleinput))
 

#### Returns

| Name | Description |
|------|-------------|
| branchProtectionRule ([BranchProtectionRule](objects.md#branchprotectionrule)) | <p>The newly created BranchProtectionRule.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### updateCheckRun

<p>Update a check run</p>

#### Input fields

- input ([UpdateCheckRunInput!](input_objects.md#updatecheckruninput))
 

#### Returns

| Name | Description |
|------|-------------|
| checkRun ([CheckRun](objects.md#checkrun)) | <p>The updated check run.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### updateCheckSuitePreferences

<p>Modifies the settings of an existing check suite</p>

#### Input fields

- input ([UpdateCheckSuitePreferencesInput!](input_objects.md#updatechecksuitepreferencesinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The updated repository.</p> |

---

### updateEnterpriseAdministratorRole

<p>Updates the role of an enterprise administrator.</p>

#### Input fields

- input ([UpdateEnterpriseAdministratorRoleInput!](input_objects.md#updateenterpriseadministratorroleinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of changing the administrator&rsquo;s role.</p> |

---

### updateEnterpriseAllowPrivateRepositoryForkingSetting

<p>Sets whether private repository forks are enabled for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseAllowPrivateRepositoryForkingSettingInput!](input_objects.md#updateenterpriseallowprivaterepositoryforkingsettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated allow private repository forking setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the allow private repository forking setting.</p> |

---

### updateEnterpriseDefaultRepositoryPermissionSetting

<p>Sets the default repository permission for organizations in an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseDefaultRepositoryPermissionSettingInput!](input_objects.md#updateenterprisedefaultrepositorypermissionsettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated default repository permission setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the default repository permission setting.</p> |

---

### updateEnterpriseMembersCanChangeRepositoryVisibilitySetting

<p>Sets whether organization members with admin permissions on a repository can change repository visibility.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanChangeRepositoryVisibilitySettingInput!](input_objects.md#updateenterprisememberscanchangerepositoryvisibilitysettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can change repository visibility setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can change repository visibility setting.</p> |

---

### updateEnterpriseMembersCanCreateRepositoriesSetting

<p>Sets the members can create repositories setting for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanCreateRepositoriesSettingInput!](input_objects.md#updateenterprisememberscancreaterepositoriessettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can create repositories setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can create repositories setting.</p> |

---

### updateEnterpriseMembersCanDeleteIssuesSetting

<p>Sets the members can delete issues setting for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanDeleteIssuesSettingInput!](input_objects.md#updateenterprisememberscandeleteissuessettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can delete issues setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can delete issues setting.</p> |

---

### updateEnterpriseMembersCanDeleteRepositoriesSetting

<p>Sets the members can delete repositories setting for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanDeleteRepositoriesSettingInput!](input_objects.md#updateenterprisememberscandeleterepositoriessettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can delete repositories setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can delete repositories setting.</p> |

---

### updateEnterpriseMembersCanInviteCollaboratorsSetting

<p>Sets whether members can invite collaborators are enabled for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanInviteCollaboratorsSettingInput!](input_objects.md#updateenterprisememberscaninvitecollaboratorssettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can invite collaborators setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can invite collaborators setting.</p> |

---

### updateEnterpriseMembersCanMakePurchasesSetting

<p>Sets whether or not an organization admin can make purchases.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanMakePurchasesSettingInput!](input_objects.md#updateenterprisememberscanmakepurchasessettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can make purchases setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can make purchases setting.</p> |

---

### updateEnterpriseMembersCanUpdateProtectedBranchesSetting

<p>Sets the members can update protected branches setting for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanUpdateProtectedBranchesSettingInput!](input_objects.md#updateenterprisememberscanupdateprotectedbranchessettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can update protected branches setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can update protected branches setting.</p> |

---

### updateEnterpriseMembersCanViewDependencyInsightsSetting

<p>Sets the members can view dependency insights for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseMembersCanViewDependencyInsightsSettingInput!](input_objects.md#updateenterprisememberscanviewdependencyinsightssettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated members can view dependency insights setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the members can view dependency insights setting.</p> |

---

### updateEnterpriseOrganizationProjectsSetting

<p>Sets whether organization projects are enabled for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseOrganizationProjectsSettingInput!](input_objects.md#updateenterpriseorganizationprojectssettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated organization projects setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the organization projects setting.</p> |

---

### updateEnterpriseProfile

<p>Updates an enterprise&rsquo;s profile.</p>

#### Input fields

- input ([UpdateEnterpriseProfileInput!](input_objects.md#updateenterpriseprofileinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The updated enterprise.</p> |

---

### updateEnterpriseRepositoryProjectsSetting

<p>Sets whether repository projects are enabled for a enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseRepositoryProjectsSettingInput!](input_objects.md#updateenterpriserepositoryprojectssettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated repository projects setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the repository projects setting.</p> |

---

### updateEnterpriseTeamDiscussionsSetting

<p>Sets whether team discussions are enabled for an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseTeamDiscussionsSettingInput!](input_objects.md#updateenterpriseteamdiscussionssettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated team discussions setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the team discussions setting.</p> |

---

### updateEnterpriseTwoFactorAuthenticationRequiredSetting

<p>Sets whether two factor authentication is required for all users in an enterprise.</p>

#### Input fields

- input ([UpdateEnterpriseTwoFactorAuthenticationRequiredSettingInput!](input_objects.md#updateenterprisetwofactorauthenticationrequiredsettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| enterprise ([Enterprise](objects.md#enterprise)) | <p>The enterprise with the updated two factor authentication required setting.</p> |
| message ([String](scalars.md#string)) | <p>A message confirming the result of updating the two factor authentication required setting.</p> |

---

### updateIpAllowListEnabledSetting

<p>Sets whether an IP allow list is enabled on an owner.</p>

#### Input fields

- input ([UpdateIpAllowListEnabledSettingInput!](input_objects.md#updateipallowlistenabledsettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| owner ([IpAllowListOwner](unions.md#ipallowlistowner)) | <p>The IP allow list owner on which the setting was updated.</p> |

---

### updateIpAllowListEntry

<p>Updates an IP allow list entry.</p>

#### Input fields

- input ([UpdateIpAllowListEntryInput!](input_objects.md#updateipallowlistentryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| ipAllowListEntry ([IpAllowListEntry](objects.md#ipallowlistentry)) | <p>The IP allow list entry that was updated.</p> |

---

### updateIssue

<p>Updates an Issue.</p>

#### Input fields

- input ([UpdateIssueInput!](input_objects.md#updateissueinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issue ([Issue](objects.md#issue)) | <p>The issue.</p> |

---

### updateIssueComment

<p>Updates an IssueComment object.</p>

#### Input fields

- input ([UpdateIssueCommentInput!](input_objects.md#updateissuecommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| issueComment ([IssueComment](objects.md#issuecomment)) | <p>The updated comment.</p> |

---

### updateLabel

<p>Updates an existing label.</p>

#### Input fields

- input ([UpdateLabelInput!](input_objects.md#updatelabelinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| label ([Label](objects.md#label)) | <p>The updated label.</p> |

---

### updateNotificationRestrictionSetting

<p>Update the setting to restrict notifications to only verified domains available to an owner.</p>

#### Input fields

- input ([UpdateNotificationRestrictionSettingInput!](input_objects.md#updatenotificationrestrictionsettinginput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| owner ([VerifiableDomainOwner](unions.md#verifiabledomainowner)) | <p>The owner on which the setting was updated.</p> |

---

### updateProject

<p>Updates an existing project.</p>

#### Input fields

- input ([UpdateProjectInput!](input_objects.md#updateprojectinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| project ([Project](objects.md#project)) | <p>The updated project.</p> |

---

### updateProjectCard

<p>Updates an existing project card.</p>

#### Input fields

- input ([UpdateProjectCardInput!](input_objects.md#updateprojectcardinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| projectCard ([ProjectCard](objects.md#projectcard)) | <p>The updated ProjectCard.</p> |

---

### updateProjectColumn

<p>Updates an existing project column.</p>

#### Input fields

- input ([UpdateProjectColumnInput!](input_objects.md#updateprojectcolumninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| projectColumn ([ProjectColumn](objects.md#projectcolumn)) | <p>The updated project column.</p> |

---

### updatePullRequest

<p>Update a pull request</p>

#### Input fields

- input ([UpdatePullRequestInput!](input_objects.md#updatepullrequestinput))
 

#### Returns

| Name | Description |
|------|-------------|
| actor ([Actor](interfaces.md#actor)) | <p>Identifies the actor who performed the event.</p> |
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequest ([PullRequest](objects.md#pullrequest)) | <p>The updated pull request.</p> |

---

### updatePullRequestReview

<p>Updates the body of a pull request review.</p>

#### Input fields

- input ([UpdatePullRequestReviewInput!](input_objects.md#updatepullrequestreviewinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequestReview ([PullRequestReview](objects.md#pullrequestreview)) | <p>The updated pull request review.</p> |

---

### updatePullRequestReviewComment

<p>Updates a pull request review comment.</p>

#### Input fields

- input ([UpdatePullRequestReviewCommentInput!](input_objects.md#updatepullrequestreviewcommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| pullRequestReviewComment ([PullRequestReviewComment](objects.md#pullrequestreviewcomment)) | <p>The updated comment.</p> |

---

### updateRef

<p>Update a Git Ref.</p>

#### Input fields

- input ([UpdateRefInput!](input_objects.md#updaterefinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| ref ([Ref](objects.md#ref)) | <p>The updated Ref.</p> |

---

### updateRefs

<p>Creates, updates and/or deletes multiple refs in a repository.</p>

<p>This mutation takes a list of <code>RefUpdate</code>s and performs these updates
on the repository. All updates are performed atomically, meaning that
if one of them is rejected, no other ref will be modified.</p>

<p><code>RefUpdate.beforeOid</code> specifies that the given reference needs to point
to the given value before performing any updates. A value of
<code>0000000000000000000000000000000000000000</code> can be used to verify that
the references should not exist.</p>

<p><code>RefUpdate.afterOid</code> specifies the value that the given reference
will point to after performing all updates. A value of
<code>0000000000000000000000000000000000000000</code> can be used to delete a
reference.</p>

<p>If <code>RefUpdate.force</code> is set to <code>true</code>, a non-fast-forward updates
for the given reference will be allowed.</p>

#### Input fields

- input ([UpdateRefsInput!](input_objects.md#updaterefsinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |

---

### updateRepository

<p>Update information about a repository.</p>

#### Input fields

- input ([UpdateRepositoryInput!](input_objects.md#updaterepositoryinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| repository ([Repository](objects.md#repository)) | <p>The updated repository.</p> |

---

### updateSubscription

<p>Updates the state for subscribable subjects.</p>

#### Input fields

- input ([UpdateSubscriptionInput!](input_objects.md#updatesubscriptioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| subscribable ([Subscribable](interfaces.md#subscribable)) | <p>The input subscribable entity.</p> |

---

### updateTeamDiscussion

<p>Updates a team discussion.</p>

#### Input fields

- input ([UpdateTeamDiscussionInput!](input_objects.md#updateteamdiscussioninput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| teamDiscussion ([TeamDiscussion](objects.md#teamdiscussion)) | <p>The updated discussion.</p> |

---

### updateTeamDiscussionComment

<p>Updates a discussion comment.</p>

#### Input fields

- input ([UpdateTeamDiscussionCommentInput!](input_objects.md#updateteamdiscussioncommentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| teamDiscussionComment ([TeamDiscussionComment](objects.md#teamdiscussioncomment)) | <p>The updated comment.</p> |

---

### updateTeamReviewAssignment

<p>Updates team review assignment.</p>

#### Input fields

- input ([UpdateTeamReviewAssignmentInput!](input_objects.md#updateteamreviewassignmentinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| team ([Team](objects.md#team)) | <p>The team that was modified</p> |

---

### updateTopics

<p>Replaces the repository&rsquo;s topics with the given topics.</p>

#### Input fields

- input ([UpdateTopicsInput!](input_objects.md#updatetopicsinput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| invalidTopicNames ([[String!]](scalars.md#string)) | <p>Names of the provided topics that are not valid.</p> |
| repository ([Repository](objects.md#repository)) | <p>The updated repository.</p> |

---

### verifyVerifiableDomain

<p>Verify that a verifiable domain has the expected DNS record.</p>

#### Input fields

- input ([VerifyVerifiableDomainInput!](input_objects.md#verifyverifiabledomaininput))
 

#### Returns

| Name | Description |
|------|-------------|
| clientMutationId ([String](scalars.md#string)) | <p>A unique identifier for the client performing the mutation.</p> |
| domain ([VerifiableDomain](objects.md#verifiabledomain)) | <p>The verifiable domain that was verified.</p> |

---