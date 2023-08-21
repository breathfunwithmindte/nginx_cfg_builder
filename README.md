# README.md for nginx_cfg_builder

---

## Table of Contents
1. [Introduction](#introduction)
2. [Installation and Usage](#installation-and-usage)
3. [Configuration Structure](#configuration-structure)
   - [Single File Configuration](#single-file-configuration)
   - [Folder Configuration](#folder-configuration)
4. [Example Configurations](#example-configurations)
5. [About the Author](#about-the-author)
6. [Contributing and Feedback](#contributing-and-feedback)
7. [License](#license)

---

## Introduction

`nginx_cfg_builder` is a robust utility crafted to transform user-friendly JSON configurations into efficient Nginx configuration files. Leveraging the capabilities of Go, this tool promises consistent, error-free, and optimized Nginx configurations.

---

## Installation and Usage

To utilize `nginx_cfg_builder`, ensure Go is installed on your machine and then clone the repository. The primary function resides in `main.go`, which reads configurations from the `/configs` directory and outputs the Nginx configurations in the `/build` directory.

---

## Configuration Structure

You can structure your JSON configurations in two main ways:

### Single File Configuration

This method enables you to define all configurations in one JSON file. The structure comprises:

- **name**: Identifier for the configuration.
- **output**: Designation of the output configuration file.
- **version**: Version identifier.
- **items**: An array of configuration items.

Each item within the `items` array can be a `comment` or `server`.

Example:

```json
{
  "kind": "comment",
  "builder": "v1",
  "docs": ["Your comment here"]
}
```

### Folder Configuration

This approach allows you to spread your configurations across multiple JSON files inside a folder. The root folder must have an `index.json` and an `items/` sub-directory containing individual JSON configurations.

- **index.json**: Contains metadata about the configuration collection.
- **items/ folder**: Contains individual JSON configuration files.

---

## Example Configurations

With the `nginx_cfg_builder`, you have the flexibility to structure your configurations either in a single file or spread across folders.

### Single File Configuration

A typical configuration in a single file might look like:

\```
{
  "name": "perf_nginxbuilder",
  ...
}
\```

### Folder Configuration

For the folder-based approach, your directory structure might resemble:

\```
/folder_case/
|-- index.json
|-- items/
    |-- config1.json
    |-- config2.json
\```

`index.json`:

\```
{
  "name": "perfect-evolution",
  ...
}
\```

Each JSON file in the `items/` directory would contain configurations similar to those you'd find in the single-file approach.

---

## About the Author

Mike Karypidis is the mastermind behind `nginx_cfg_builder`. With an in-depth grasp of Go and Nginx configurations, Mike developed this tool to simplify and streamline the Nginx configuration generation process.

---

## Contributing and Feedback

We appreciate your feedback and contributions to `nginx_cfg_builder`. If you come across any issues or have suggestions, please create an issue on our GitHub repository.

For those interested in contributing further, please ensure your pull requests conform to the project's guidelines and coding standards.

---

## License

For detailed licensing information, please refer to the LICENSE file included in the repository.
