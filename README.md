# Bifrost CLI Tool

Bifrost is a simple CLI tool for uploading files to cloud storage services. It provides a simple and easy-to-use interface for uploading files to cloud storage services, such as Amazon S3 and Google Cloud Storage.

## Installation

To install Bifrost, run the following command:

```bash
go get github.com/showbaba/bifrost-cli
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


## Contributing
If you want to contribute to Bifrost CLI, please submit a pull request on GitHub.

## License
Bifrost is released under the MIT License. See LICENSE for details.