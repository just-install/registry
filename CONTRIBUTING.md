# Contributing

Adding a new entry to just-install's registry is easy but we have to set the bar of what goes into
it.

Not everything in the registry currently follows these guidelines. We also make exceptions to these
rules from time to time. We have to draw a line somewhere and we hope that you trust us to use our
discretion to handle what should and should not go in the main registry.

We are (slowly) adding features that should allow you to run your custom registries so that you are
not dependent on our rules that may go against your needs.

## Guidelines

- **Stable versions**: We only add stable versions that the developer deems stable. Exceptions are
  apps that are only available in pre-release form and that are sufficiently popular.
- **Must have an installer**: Portable applications and stuff distributed via Zip, Tar, or 7zip
  archives will not be accepted. 7zip in particular is especially hard to support. We expect the
  installer to create a Start menu entry to launch the application.
- **Must have a known installer format**: Bespoke installers or distribution formats (e.g. 7zip
  archives, weird one-of-a-kind setup tools) will not be accepted unless they are immensely popular,
  in which case we probably already support it. The Windows installer ecosystem is already very,
  very fragmented. Please refrain from adding even more options to the mix.
- **Must be fully silent AND unattended**: While the point above ensures that we are able to carry
  out a silent installation in the vast majority of cases, some software still requires manual
  intervention, for example because they launch a GUI application as part of a post-install step.
  This makes automated deployments almost impossible. As such we cannot accept this kind of
  software.
- **Should install for all users on the system**: just-install runs as `Administrator` and we expect
  installers to automatically install the software for all users on the system. If an installer
  doesn't, we should pass the appropriate command line switches for a system-wide installation,
  unless it's absolutely impossible to do so. Some software (e.g. Visual Studio Code) provides
  "System" and "User" installers as separate downloads. You should always choose the "System"
  installer.
- **Popular**: Apps added to the registry should be reasonably popular. We won't accept entries that
  seem too obscure or unmaintained unless there's strong demand for it.

## Acceptable and Unacceptable Content

Almost everything under the sun is allowed, with the following **exceptions**:

- **Anything from Oracle**: It's a legal minefield. Sorry.
- **Bundles**: Some applications bundle other software with their installer. For example, some
  popular utility apps bundle the Chrome browser. Such software will be rejected unless its
  installer provides a way to exclude installation of bundled software, either automatically for
  silent installations, or via command line switches. An exception is made for software bundling
  necessary run-time libraries required to work properly (e.g. Visual C++ Runtime, .NET Framework,
  DirectX, etc).
- **Games**: Most games nowadays are distributed through game management platforms such as Steam,
  GoG, and similar. We provide means to install their clients but not the games themselves.
- **Node/Python/Ruby packages, plug-ins, browser extensions, toolbars**: Packages available from
  application-level package managers (e.g. `npm`, `pip`, `gem`) should be installed by the
  appropriate tool, not through just-install. Plug-ins or browser extensions are similarly not
  accepted. Many applications provide their own extension marketplace or plug-in manager and we
  don't interfere with them.
- **Scripts**: PowerShell, VBS, and batch scripts are not accepted.
