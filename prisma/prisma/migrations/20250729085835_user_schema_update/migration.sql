-- AlterTable
ALTER TABLE "User" ADD COLUMN     "company_name" TEXT,
ADD COLUMN     "cv_json" JSONB,
ADD COLUMN     "description" TEXT,
ADD COLUMN     "location" TEXT,
ADD COLUMN     "professional_title" TEXT;
