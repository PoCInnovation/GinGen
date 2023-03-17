# GinGen

Nowadays, in order to make developers' tasks easier, numerous tools of all kinds have been invented. This is particularly true for anything related to code documentation, as it is a crucial part for ensuring that every team member can quickly understand how each part works.

Regarding APIs, a documentation system called Swagger Editor has been implemented to facilitate the understanding of each endpoint and their different bodies and responses. Unfortunately, in order to generate this documentation, Swagger requires a JSON document with different fields to correctly transcribe the information. However, many people would find it simpler and more natural to simply document each function directly in the code, so that it can be transcribed into Swagger documentation.

This is where the idea for **GinGen** came about, an algorithm that allows for a specific documentation present in a **GO file code** to be transcribed **into JSON** that can be used in **Swagger**.

## How does it work?

This project focuses on two important parts of an API:
- **Endpoints**, which are the different triggers that are called when a specific URL is executed.
- **Handlers**, which are the methods called following the call of an endpoint.

Therefore, there are two types of documentation to look for.

For now, we are parsing only one file at a time, but we plan to delve deeper by allowing the program to be launched at the root of an API in order to retrieve each part of the documentation.

## Getting Started

### Installation

Clone the repository and simply execute:
```make```

to build the GinGen exacutable

### Quickstart

To run the program, two arguments are required:
- the file to parse
- the json file in which to write the result.

```./GinGen -i main.go -o doc.json```
```./GinGen --input main.go --output doc.json```

### Usage

[Explain how to use this project]

## Get involved

You're invited to join this project ! Check out the [contributing guide](./CONTRIBUTING.md).

If you're interested in how the project is organized at a higher level, please contact the current project manager.

## Our PoC team :heart:

Developers
| [<img src="https://github.com/VidsSkids.png?size=85" width=85><br><sub>Victor Massonnat</sub>](https://github.com/VidsSkids) | [<img src="https://github.com/TdeBoynes.png?size=85" width=85><br><sub>Timothée De Boynes</sub>](https://github.com/TdeBoynes)
| :---: | :---: |

Manager
| [<img src="https://github.com/RezaRahemtola.png?size=85" width=85><br><sub>Reza Rahemtola</sub>](https://github.com/RezaRahemtola)
| :---: |

<h2 align=center>
Organization
</h2>

<p align='center'>
    <a href="https://www.linkedin.com/company/pocinnovation/mycompany/">
        <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white">
    </a>
    <a href="https://www.instagram.com/pocinnovation/">
        <img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white">
    </a>
    <a href="https://twitter.com/PoCInnovation">
        <img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white">
    </a>
    <a href="https://discord.com/invite/Yqq2ADGDS7">
        <img src="https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white">
    </a>
</p>
<p align=center>
    <a href="https://www.poc-innovation.fr/">
        <img src="https://img.shields.io/badge/WebSite-1a2b6d?style=for-the-badge&logo=GitHub Sponsors&logoColor=white">
    </a>
</p>

> :rocket: Don't hesitate to follow us on our different networks, and put a star 🌟 on `PoC's` repositories

> Made with :heart: by PoC
