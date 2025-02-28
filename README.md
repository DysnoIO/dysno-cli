# Dysno CLI

The **Dysno CLI** is a command-line tool for mining Dysno nuggets, which can be redeemed for mining rewards. This CLI provides commands to manage your Dysno mining account, configure settings, test your connection, and start mining

---

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [Install from Precompiled Executables](#install-from-precompiled-executables)
  - [Install from Source](#install-from-source)
- [Usage](#usage)
  - [Commands](#commands)
    - [`account`](#account)
    - [`test`](#test)
    - [`mine`](#mine)
    - [`config`](#config)
- [Contributions](#contributions)
- [License](#license)
- [Acknowledgments](#acknowledgments)

---

## Features

- Manage Dysno mining account(s).
- Test mining configuration and connection.
- Start mining Dysno nuggets.
- Configure CLI settings.

---

## Installation

### Install from Precompiled Executables

1. Go to the [Releases Page](https://github.com/dysnoio/dysno-cli/releases) and download the appropriate executable for your operating system.

2. Extract the downloaded file if necessary.

3. Create a dedicated folder for Dysno CLI, for example:

   ```bash
   mkdir -p $HOME/dysno-cli
   ```

4. Move the executable to this folder:

   ```bash
   mv dysno $HOME/dysno-cli/
   ```

5. Add the folder to your system PATH. For example, on Linux or macOS:

   ```bash
   echo 'export PATH=$HOME/dysno-cli:$PATH' >> $HOME/.bashrc
   source $HOME/.bashrc
   ```

On Windows, update the PATH variable in the system environment variables.

### Install from Source

1. Ensure you have Go installed. If not, download and install Go from [https://go.dev/dl/](https://go.dev/dl/).

2. Clone the repository:

   ```bash
   git clone https://github.com/your-username/dysno-cli.git
   cd dysno-cli
   ```

3. Build the CLI from source:

   ```bash
   go build -o dysno
   ```

4. Create a dedicated folder for Dysno CLI:

   ```bash
   mkdir -p $HOME/dysno-cli
   ```

5. Move the `dysno` binary to this folder:

   ```bash
   mv dysno $HOME/dysno-cli/
   ```

6. Add the folder to your system PATH. For example:

   ```bash
   echo 'export PATH=$HOME/dysno-cli:$PATH' >> $HOME/.bashrc
   source $HOME/.bashrc
   ```

---

## Usage

After installation, you can access the CLI by running the `dysno` command in your terminal.

### Commands

#### `account`

Manage your Dysno account(s). This command has the following subcommands:

- **add**: Add a new account.

  ```bash
  dysno account add
  ```

- **set**: Set the active account.

  ```bash
  dysno account set
  ```

- **delete**: Delete an existing account.

  ```bash
  dysno account delete
  ```

- **list**: List all accounts.

  ```bash
  dysno account list
  ```

#### `test`

Test your mining setup and connection.

```bash
dysno test
```

#### `mine`

Start mining Dysno nuggets.

```bash
dysno mine
```

#### `config`

Configure the Dysno CLI settings. This command has the following subcommands:

- **rpc**: Configure the RPC endpoints.

  ```bash
  dysno config rpc
  ```

- **ca**: Set the contract address.

  ```bash
  dysno config ca
  ```

- **gaslimit**: Configure the gas limit.

  ```bash
  dysno config gaslimit
  ```

- **chain**: Set the chain ID.

  ```bash
  dysno config chain
  ```

- **show**: Display the current configuration.

  ```bash
  dysno config show
  ```

---

## Contributions

Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Thank you for using Dysno CLI!
