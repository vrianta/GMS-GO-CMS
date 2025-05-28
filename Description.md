# GMS-GO-CMS Project Description

GMS-GO-CMS is a modern, developer-first Content Management System (CMS) framework built in Golang, inspired by the flexibility and extensibility of WordPress but designed for high performance, static typing, and developer productivity.

## Key Features

- **MVC Architecture:** Clean separation of concerns for scalable and maintainable code.
- **PHP-Style Templating:** Write templates using familiar `<?= $var ?>` syntax, automatically transpiled to Go's `html/template` format.
- **Dynamic Content Types:** Define custom content types (like posts, pages, etc.) without writing SQL; the system generates and migrates the database schema automatically.
- **Automatic Code Generation:** When new content types are created, Go models, controllers, and routes are generated and integrated into the application.
- **Integrated Build Tooling:** Lightweight CLI tools automate rebuilding, scaffolding, and running the app.
- **Performance:** Built on Go’s HTTP engine for speed and concurrency.
- **Frontend Flexibility:** Fully customizable frontend, with optional React integration.
- **Open Source Core:** Licensed under AGPL for individuals, with a commercial license available for organizations.

## Goals

- Deliver a fast, extensible CMS framework for Go developers.
- Enable dynamic configuration and content modeling while preserving static typing.
- Provide a backend-first alternative to WordPress, optimized for performance and developer experience.
- Lower the barrier for custom CMS development in Go.

## Tools & Concepts

- Custom transpiler for PHP-like templates
- Dynamic file/code generation and recompilation
- Automated database migrations
- CLI for scaffolding and build automation
- Dual licensing model (AGPL/commercial)

GMS-GO-CMS aims to empower developers and teams to build robust, high-performance content platforms with the flexibility of dynamic systems and the safety of static typing.