# Bifrost CLI

A simple CLI tool for [Bifrost](https://github.com/opensaucerer/bifrost/). Biforst is a package for uploading files to cloud storage services. It provides a simple and easy-to-use interface for uploading files to cloud storage services, such as Amazon S3 and Google Cloud Storage.

## Installation

To install Bifrost, run the following command:

```bash
go install github.com/showbaba/bifrost-cli@latest
```


## Usage
To use Bifrost CLI, run the following command:

```bash
bifrost-cli upload -d -r
```

You will be prompted to enter the following information:

    - Cloud provider name (e.g. s3, gcs)
    - Default bucket name
    - Access key
    - Secret key
    - Region name
    - Zone name
    - File path
    - File name

After providing the required details, the CLI tool will start uploading your file to the cloud storage service.

## Command Options
The following command options are available:

    * --enable-debug or -d: Enables debug mode.
    * --public-read or -r: Makes uploaded files public.

## Running the Binary on Any OS

If you don't want to install Bifrost CLI via Go, you can download the binary file for your operating system from the releases page on GitHub.

To use the binary file, follow these steps:

1. Download the appropriate binary file for your operating system from the [releases page on GitHub](https://github.com/ShowBaba/bifrost-cli/tags).

2. Open a terminal or command prompt window.

3. Navigate to the directory where you downloaded the binary file.

4. Run the following command:

On Windows:
```bash
bifrost-cli.exe upload -d -r
```

On Linux or macOS:
```bash
./bifrost-cli upload -d -r
```

You will be prompted to enter the required information as described above.

## Contributing
If you want to contribute to Bifrost CLI, please submit a pull request on GitHub.

## License
Bifrost CLI is released under the MIT License. See LICENSE for details.