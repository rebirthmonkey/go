# Commit 规范

## 作用

一个好的 Commit Message  至关重要：

- 能够清晰地展示每个 commit 的变更内容，方便快速浏览变更历史，比如可以直接略过文档类型或者格式化类型的代码变更。
- 可以基于这些 Commit Message 进行过滤查找，比如只查找某个版本新增的功能：git log --oneline --grep  "^feat|^fix|^perf"。
- 可以基于规范化的 Commit Message 生成 Change Log。
- 可以依据某些类型的 Commit Message 触发构建或发布流程，比如当 type 为 feat、fix 时才触发 CI 流程。
- 确定语义化版本的版本号：比如 fix 类型可以映射为 PATCH 版本，feat 类型可以映射为 MINOR 版本。带有 BREAKING CHANGE 的 commit，可以映射为 MAJOR 版本。

总之，一个好的 Commit Message 规范可以使 Commit Message 的可读性更好，并且可以实现自动化。

## Angular 规范

在众多 commit 规范中，Angular 规范在功能上最能够满足开发者 commit 需求，在格式上清晰易读，目前也是用得最多的。Angular 规范其实是一种语义化的提交规范（Semantic Commit Messages），所谓语义化的提交规范包含以下内容：

- Commit  Message 是语义化的：Commit Message 会被归为一个有意义的类型，用来说明本次 commit 的类型。
- Commit  Message 是规范化的：Commit Message 遵循预先定义好的规范，比如 Commit Message  格式固定、都属于某个类型，这些规范不仅可被开发者识别也可以被工具识别。

在 Angular 规范中，Commit Message 包含三个部分，分别是 Header、Body 和 Footer，格式如下：

```
<type>[optional scope]:<description> 
// 空行
[optional body]
// 空行
[optional footer(s)]
```

其中，Header 是必需的，Body 和 Footer 可以省略。在以上规范中，<scope> 必须用括号 () 括起来，<type> [<scope>] 后必须紧跟冒号 ，冒号后必须紧跟空格，2 个空行也是必需的。

在实际开发中，为了使 Commit Message 在 GitHub 上更加易读，往往会限制每行 message 的长度。根据需要，可以限制为 50/72/100 个字符，这里将长度限制在 72 个字符以内。以下是一个符合 Angular 规范的 Commit  Message：

```
fix($compile): couple of unit tests for IE9
# Please enter the Commit Message for your changes. Lines starting
# with '#' will be  ignored, and an empty message aborts the commit.
# On branch master
# Changes to be committed:
# ...

Older IEs serialize html uppercased, but  IE9 does not...
Would be better to expect case insensitive, unfortunately jasmine does
not allow to user regexps for throw expectations.

Closes  #392Breaks foo.bar api, foo.baz should be used instead
```

### Header

Header  部分只有一行，包括三个字段：type（必选）、scope（可选）和 subject（必选）：

#### type

用来说明 commit 的类型。它们主要可以归为 Development 和 Production 共两类：1/ Development：这类修改一般是项目管理类的变更，不会影响生产环境的代码，比如 CI 流程、构建方式等的修改。遇到这类修改，通常也意味着可以免测发布。2/ Production：这类修改会影响最终生产环境的代码。所以对于这种改动，一定要慎重，并在提交前做好充分的测试。这里列出了 Angular 规范中的常见 type 和它们所属的类别，在提交 Commit Message 的时候，一定要注意区分它的类别。例如，在做 Code Review 时，如果遇到 Production 类型的代码，一定要认真  Review，因为这种类型，会影响到现网用户的使用和现网应用的功能。

![img](figures/89c618a7415c0c38b09d86d7f882a427.png)

如果变更了应用代码，比如某个 Go 函数代码，那这次修改属于代码类。在代码类中，有 4 种具有明确变更意图的类型：feat、fix、perf 和  style；如果代码变更不属于这 4 类，那就全都归为 refactor 类，也就是优化代码。如果变更了非应用代码，例如更改了文档，那它属于非代码类。在非代码类中，有 3 种具有明确变更意图的类型：test、ci、docs；如果非代码变更不属于这 3 类，那就全部归入到 chore 类。Angular 的  Commit Message 规范提供了大部分的 type，在实际开发中，可以使用部分 type，或扩展添加自己的 type。但无论选择哪种方式，一定要保证一个项目中的 type 类型一致。

<img src="figures/3509bd169ce285f59fbcfa6ebea75aa7.png" alt="img" style="zoom:33%;" />

#### scope

scope 是用来说明 commit 的影响范围的，必须是名词。不同项目会有不同的 scope。在项目初期，可以设置一些粒度比较大的 scope，比如可以按组件名或者功能来设置 scope。后续，如果项目有变动或有新功能，可以再用追加的方式添加新的 scope。在大部分情况下，scope 主要是根据组件名和功能来设置的。例如，支持 apiserver、authzserver、user 这些 scope。另外，scope 不适合设置的太具体。太具体的话，一方面会导致项目有太多的  scope，难以维护。另一方面，开发者也难以确定 commit 属于哪个具体的 scope，导致错放 scope，反而会使 scope  失去了分类的意义。

当然，在指定 scope 时，也需要遵循预先规划的 scope，所以要将 scope 文档化，放在类似 docs/devel 目录下，具体可参考 [scope.md](31_scope.md) 文档。

#### subject

subject 是 commit 的简短描述，必须以动词开头、使用现在时。比如，可以用 change，却不能用  changed 或 changes，而且这个动词的第一个字母必须是小写。通过这个动词，可以明确地知道 commit  所执行的操作。此外，subject 的结尾不能加英文句号。

### Body

Header 对 commit  做了高度概括，可以方便查看 Commit Message。Body 部分是对本次 commit 的更详细描述，是可选的。Body 部分可以分成多行，而且格式也比较自由。不过，和 Header 里的一样，它也要以动词开头，使用现在时。此外，它还必须要包括修改的动机，以及和跟上一版本相比的改动点。例如：

```
The body is mandatory for all commits except for those of scope "docs".  
When the body is required it must be at least 20 characters long.
```

### Footer

Footer 部分不是必选的，可以根据需要来选择，主要用来说明本次 commit 导致的后果。在实际应用中，Footer 通常用来说明不兼容的改动和关闭的 Issue 列表，格式如下：

```
BREAKING CHANGE: <breaking change summary>
// 空行
<breaking change description + migration instructions>
// 空行
// 空行
Fixes #<issue number>
```



接下来，我给你详细说明下这两种情况：

- 不兼容的改动：如果当前代码跟上一个版本不兼容，需要在 Footer 部分，以 BREAKING CHANG: 开头，后面跟上不兼容改动的摘要。Footer 的其他部分需要说明变动的描述、变动的理由和迁移方法，例如：

```

BREAKING CHANGE: isolate scope bindings definition has changed and
    the inject option for the directive controller injection was removed.

    To migrate the code follow the example below:

    Before:

    scope: {
      myAttr: 'attribute',
    }

    After:

    scope: {
      myAttr: '@',
    }
    The removed `inject` wasn't generaly useful for directives so there should be no code using it.
```

- 关闭的 Issue 列表：关闭的 Bug 需要在 Footer 部分新建一行，并以 Closes 开头列出，例如：Closes #123。如果关闭了多个 Issue，可以这样列出：Closes #123, #432, #886。例如:

```
 Change pause version value to a constant for image
    
    Closes #1137
```

### Revert Commit

除了 Header、Body 和 Footer 这 3 个部分，Commit Message 还有一种特殊情况：如果当前 commit 还原了先前的  commit，则应以 revert: 开头，后跟还原的 commit 的 Header。而且，在 Body 中必须写成 This reverts commit，其中 hash 是要还原的 commit 的 SHA 标识。例如：

```
revert: feat(iam-apiserver): add 'Host' option

This reverts commit 079360c7cfc830ea8a6e13f4c8b8114febc9b48a.
```

为了更好地遵循 Angular  规范，建议你在提交代码时养成不用 git commit -m，即不用 -m 选项的习惯，而是直接用 git commit 或 git  commit -a 进入交互界面编辑 Commit Message。这样可以更好地格式化 Commit Message。但是

## 提交操作

除了 Commit  Message 规范之外，在代码提交时，还需要关注 3 个重点内容：提交频率、合并提交和 Commit Message 修改。

### 提交频率

在实际项目开发中，如果是个人项目，随意 commit 可能影响不大，但如果是多人开发的项目，随意 commit 不仅会让 Commit Message 变得难以理解，还会让其他研发同事觉得你不专业。因此，要规定 commit 的提交频率。主要可以分成两种情况：

- 只要对项目进行了一个重要的修改，一通过测试就立即 commit。比如修复完一个 bug、开发完一个小功能，或开发完一个完整的功能，测试通过后就提交。
- 规定一个时间，定期提交：这我建议代码下班前固定提交一次，并且要确保本地未提交的代码，延期不超过 1 天。这样，如果本地代码丢失，可以尽可能减少丢失的代码量。

### 合并提交

按照上面 2 种方式提交代码，可能会觉得代码 commit 比较多，看起来比较随意。或者说，希望等开发完一个完整的功能之后，放在一个 commit 中一起提交。这时可以在最后合并代码或者提交 Pull Request 前，执行 git rebase -i 合并之前的所有 commit。合并提交就是将多个 commit 合并为一个 commit 提交。建议把新的 commit 合并到主干时，只保留 2~3 个 commit 记录。

在 Git 中，主要使用 git rebase 命令来合并。git  rebase 也是日后开发需要经常使用的一个命令，所以一定要掌握好它的使用方法。git rebase 的最大作用是它可以重写历史。通常会通过 git rebase -i <commit ID> 使用 git rebase 命令，-i 参数表示交互（interactive），该命令会进入到一个交互界面中，其实就是 Vim 编辑器。在该界面中，可以对里面的 commit 做一些操作，交互界面如图所示：

<img src="figures/c63a8682c03862802e5eacf1641b86ac.png" alt="img" style="zoom:50%;" />

这个交互界面会首先列出给定之前（不包括，越下面越新）的所有 commit，每个 commit 前面有一个操作命令，默认是 pick。可以选择不同的 commit，并修改 commit 前面的命令，来对该 commit 执行不同的变更操作。git rebase 支持的变更操作如下：

![img](figures/5f5a79a5d2bde029d4de9d98026ef3f2.png)

squash 和 fixup 可以用来合并 commit。例如用 squash 来合并，只需要把要合并的 commit 前面的动词，改成 squash（或者 s）即可。

```
pick 07c5abd Introduce OpenPGP and teach basic usage
s de9b1eb Fix PostChecker::Post#urls
s 3e7ee36 Hey kids, stop all the highlighting
pick fa20af3 git interactive rebase, squash, amend
```

rebase 后，第 2 行和第 3 行的 commit 都会合并到第 1 行的 commit。这个时候，提交的信息会同时包含这三个 commit 的提交信息：

```
# This is a combination of 3 commits.
# The first commit's message is:
Introduce OpenPGP and teach basic usage

# This is the 2ndCommit Message:
Fix PostChecker::Post#urls

# This is the 3rdCommit Message:
Hey kids, stop all the highlighting
```

如果将第 3 行的 squash 命令改成 fixup 命令：

```
pick 07c5abd Introduce OpenPGP and teach basic usage
s de9b1eb Fix PostChecker::Post#urls
f 3e7ee36 Hey kids, stop all the highlighting
pick fa20af3 git interactive rebase, squash, amend
```

rebase 后，还是会生成两个 commit，第 2 行和第 3 行的 commit，都合并到第 1 行的 commit。但是，新的提交信息里面，第 3 行 commit 的提交信息会被注释掉：

```

# This is a combination of 3 commits.
# The first commit's message is:
Introduce OpenPGP and teach basic usage

# This is the 2ndCommit Message:
Fix PostChecker::Post#urls

# This is the 3rdCommit Message:
# Hey kids, stop all the highlighting
```

在使用 git rebase 进行操作的时候，还需要注意以下几点：

- 删除某个 commit 行，则该 commit 会丢失掉。
- 删除所有的 commit 行，则 rebase 会被终止掉。
- 可以对 commits 进行排序，git 会从上到下进行合并。

## 修改 Commit Message

即使有了 Commit Message 规范，但仍然可能会遇到提交的 Commit Message 不符合规范的情况，这个时候就需要修改之前某次  commit 的 Commit Message。有两种修改方法，分别对应两种不同情况：

### git commit  --amend

git commit --amend：修改最近一次 commit 的 message。有时，刚提交完一个 commit，但是发现 commit 的描述不符合规范或需要纠正，这时可以通过 git  commit --amend 命令来修改刚刚提交 commit 的 Commit Message。具体修改步骤如下：

- 查看当前分支的日志记录：

```
$ git log –oneline
418bd4 docs(docs): append test line 'update$i' to README.md
89651d4 docs(doc): add README.md
```

- 更新最近一次提交的 Commit Message：在当前 Git 仓库下执行命令：git commit --amend，在交互界面中修改最近一次的 Commit Message
- 查看最近一次的 Commit Message 是否被更新

```
$ git log --oneline
55892fa docs(docs): append test line 'update1' to README.md
89651d4 docs(doc): add README.md
```

### git rebase -i

如果想修改的 Commit Message 不是最近一次的 Commit Message，可以通过 git rebase -i <父 commit ID>命令来修改。

- 查看当前分支的日志记录:

```
$ git log --oneline
1d6289f docs(docs): append test line 'update3' to README.md
a38f808 docs(docs): append test line 'update$i' to README.md
55892fa docs(docs): append test line 'update1' to README.md
89651d4 docs(doc): add README.md
```

- 修改倒数第 3 次提交 commit 的  message。在 Git 仓库下直接执行命令 git rebase -i 55892fa，然后会进入一个交互界面。在交互界面中，修改最近一次的 Commit Message。这里使用 reword 或者 r，保留倒数第 3 次的变更信息，但是修改其 message。修改完成后执行:wq 保存，还会跳转到该 Commit Message 的交互页面。
- 查看倒数第 3 次 commit 的 message 是否被更新：

```
$ git log --oneline
7157e9e docs(docs): append test line 'update3' to README.md
5a26aa2 docs(docs): append test line 'update2' to README.md
55892fa docs(docs): append test line 'update1' to README.md
89651d4 docs(doc): add README.md
```

这里有两点需要你注意：

- Commit  Message 是 commit 数据结构中的一个属性，如果 Commit Message 有变更，则 commit ID 一定会变，git  commit --amend 只会变更最近一次的 commit ID，但是 git rebase -i 会变更父 commit ID  之后所有提交的 commit ID。
- 如果当前分支有未 commit 的代码，需要先执行 git stash  将工作状态进行暂存，当修改完成后再执行 git stash pop 恢复之前的工作状态。