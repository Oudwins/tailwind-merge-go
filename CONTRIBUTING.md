# Contributing

Contributions are welcome and will be fully credited.

Please take a moment to review this document before creating an issue or pull request. It is based on the [original tailwind-merge](https://github.com/dcastil/tailwind-merge/blob/v2.2.2/.github/CONTRIBUTING.md) contributing guidelines.

## Etiquette

This project is open source, and as such, the maintainers give their free time to build and maintain the source code held within. They make the code freely available in the hope that it will be of use to other developers. It would be extremely unfair for them to suffer abuse or anger for their hard work.

Please be considerate towards maintainers when raising issues or presenting pull requests. Let's show the world that developers are civilized and selfless people.

It's the duty of the maintainer to ensure that all submissions to the project are of sufficient quality to benefit the project. Many developers have different skillsets, strengths, and weaknesses. Respect the maintainer's decision, and do not be upset or abusive if your submission is not used.

See [CODE OF CONDUCT](/CODE_OF_CONDUCT.md) for more info.

## Viability

When requesting or submitting new features, first consider whether it might be useful to others. Open source projects are used by many developers, who may have entirely different needs to your own. Think about whether or not your feature is likely to be used by other users of the project.

## Project Overview

The relevant file structure for the project is:

```bash
./pkg/twmerge/ # Stores the code for the twmerge-go library
./pkg/cache/ # Stores the code for the default cache, an LRU cache
./cmd/main/ # for manual testing purposes
```

In some spots the code might be a little bit difficult to follow due to the heavy usage of [closures](https://developer.mozilla.org/es/docs/Web/JavaScript/Closures). So let me try to explain simply what is going on.

1. The main function that does the merging is found in `./pkgs/twmerge/merge-classlist.go`. It generates a sort of unique collision key for each class and removes any colliding classes.
2. For it to be able to correctly identify if a class collides with another it uses a custom (and very long) hashmap (see `./pkg/twmerge/default-config`) with a bunch of huristics about different tailwind classes.

Its obviously a little more complex than that. But I think that you can probably figure it out much quicker than I did with this heads up.

## Creating an issue

**Before filing an issue:**

- Attempt to replicate the problem, to ensure that it wasn't a coincidental incident.
- Check to make sure your feature suggestion isn't already present within the project.
- Check the pull requests tab to ensure that the bug doesn't have a fix in progress.
- Check the pull requests tab to ensure that the feature isn't already in progress.

## Submitting a PR

#### **Before submitting a pull request:**

- Check the codebase to ensure that your feature doesn't already exist.
- Check the pull requests to ensure that another person hasn't already submitted the feature or fix.

3. How tu submit issues, features, PRs...

#### **PR Requirements**

- **Document any change in behaviour** - Make sure the `README.md` and any other relevant documentation are kept up-to-date.

- **One pull request per feature** - If you want to do more than one thing, send multiple pull requests.

- **Send coherent history** - Make sure each individual commit in your pull request is meaningful. If you had to make multiple intermediate commits while developing, please [squash them](https://www.git-scm.com/book/en/v2/Git-Tools-Rewriting-History#Changing-Multiple-Commit-Messages) before submitting.

- **Use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for your commit messages** - This helps make reviewing PRs easier.

## Development

You will need go 1.21.5 or later installed. tailwind-merge-go has no dependencies so no need to install anything.
Finally, I recommend running tests after every change, you can do so by using this command:

```bash
make test
```
