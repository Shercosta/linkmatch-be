# Quick Start Guide

---

## Setting Up Your Development Environment

Follow these essential steps to get the project running on your local machine:

1.  **Create Your Database:** Before anything else, create a new database specifically for this project.
2.  **Root Environment Variables (`.env`):**
    - In the **root directory** of the project, create a new file named `.env`.
    - **Reference:** Use the provided `.env.local` file as a template.
3.  **Prisma Environment Variables (`prisma/.env`):**
    - Navigate into the `prisma/` folder.
    - Create another `.env` file here.
    - **Crucially:** Refer to `prisma/.env.local` and ensure you input your **actual database credentials**.
4.  **Initial Setup Script:**
    - Once steps 1-3 are finished, execute the setup script from the project's root directory:
      ```bash
      ./setup.sh
      ```
    - This script prepares everything necessary for the first run.
5.  **Running the Project:**
    - To start the application normally:
      ```bash
      go run main.go
      ```
    - For debugging with `dlv`: Simply press `F5` in your editor.

---
