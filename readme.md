#   Hello Github
一些常用的Github操作指令吧
##  初始化仓库
git init
git目录初始化

git add .
添加所有文件至暂存区

git commit -m "first commit"
创建第一次提交
PS：git commit -am ""等同于git add . 然后 -m
进行提交代码操作

git branch -M main
创建一个主分支

git remote add origin git@github.com:WeiYunfeng/testfile.git
指明连接的远程仓库

git push -u origin main
推送本地仓库主分支内容到远程仓库

##  新建分支和分支切换
git branch develop
创建develop分支

git checkout develop
切换到develop分支

当然直接使用git checkout -b develop
也会直接创建分支并自动切换到develop分支视图

一般建议在其他分支开发
git commit -am "develop ???"
然后push分支的本地修改到远程或者再合并到main push最好

##  commit重命名
git commit --amend 来修改合并提交的提示信息时，通常建议保留原有的合并信息，但根据需要进行适当的自定义，以提供更多的上下文信息。合并提交的提示信息通常应该包含有关合并的信息，以便其他开发者能够理解这个合并的目的。

##  撤销commit操作

git commit --reset HEAD~ 的命令会取消掉最新的提交，并将更改保留在工作目录中，以便您可以重新提交或修改这些更改。这个命令会将您的当前分支的 HEAD 指针移动到前一个提交（HEAD~ 表示前一个提交）并保留原始更改。
请注意，这个命令仅仅影响您的本地仓库，不会影响远程仓库。如果您已经将不希望提交的更改提交到了远程仓库，您需要额外的步骤来从远程仓库撤销这些提交。
如果您执行了 git commit --reset HEAD~，您可以在工作目录中看到未提交的更改，然后您可以对这些更改进行修改或重新提交。例如，您可以运行 git status 来查看未提交的更改，并使用 git add 和 git commit 来重新提交它们。

git commit --reset HEAD~ --soft 命令与之前的命令非常相似，但有一个重要区别。使用 --soft 参数时，它将取消最新的提交，并将更改保留在工作目录中，但不会修改索引（暂存区）。这意味着您可以重新审查并修改已取消的提交，然后将其重新提交。
这个命令在以下情况可能会很有用：
您希望修改之前的提交信息。
您意识到您忘记了一些文件，想要将它们包含在之前的提交中。
使用 --soft 参数时，运行命令后，您的工作目录中的更改将保持不变，并且取消的提交将成为未提交状态。您可以编辑提交信息，然后使用 git commit --amend 来重新提交这些更改，或者使用 git add 和 git commit 来将新更改添加到这个提交中。

**请小心使用这个命令，特别是在团队环境中，因为它会修改提交历史。如果其他人已经基于您的提交进行了工作，可能会引起混淆。**

##  分支合并
git checkout main
切换到主分支

git merge develop main
进行合并

**注意：
在 Git 中，当您执行合并操作后，Git 会自动生成一个合并提交，该提交包含有关合并的信息。因此，通常不建议直接使用 git commit -am "merge develop main" 这样的命令进行合并提交的手动添加。
Git 自动生成的合并提交消息会包含有关合并的详细信息，例如哪些分支被合并，合并的提交点等。手动添加合并提交消息可能会导致信息不一致。
如果您需要添加一些自定义信息，建议在合并后使用 git log 查看合并提交的哈希值，然后使用 git commit --amend 来修改最新的提交信息。这样可以确保合并提交信息是准确的且包含有关合并的完整信息。**

##  git pull 和 git fetch的对比
git pull 和 git fetch 是用于从远程仓库获取更新的两个不同的 Git 命令，它们之间有一些关键区别：

### git pull：

git pull 是一个组合命令，它包括了 git fetch 和 git merge（或 git rebase）两个步骤。
当您运行 git pull 时，Git 首先会从远程仓库获取最新的提交和分支信息（使用 git fetch），然后将这些更新合并（或重新基于）到当前分支。
例如，如果您在本地分支上运行 git pull origin main，Git 会从远程仓库 origin 的 main 分支获取最新的提交，并将它们合并到当前本地分支。

使用 git pull 可以方便地获取远程更改并立即合并到本地分支，但它可能会导致自动合并（merge）或自动重新基于（rebase），这可能会产生冲突需要解决。

### git fetch：

git fetch 仅用于从远程仓库获取最新的提交和分支信息，但它不会自动合并这些更改到当前分支。
运行 git fetch 后，远程仓库的最新更改将被下载到本地，但您需要手动决定如何处理这些更改，例如合并或重新基于。
例如，如果您运行 git fetch origin，Git 会将远程仓库 origin 的最新提交信息下载到本地，但不会将这些更改自动合并到当前分支。您可以使用 git merge 或 git rebase 命令来手动将这些更改合并到当前分支。

总的来说，git pull 是一个更方便的命令，因为它将获取和合并两个步骤结合在一起。但是，如果您想要更细粒度地控制获取和合并操作，或者想要先查看远程更改再决定如何处理它们，那么使用 git fetch 是一个更灵活的选择。

