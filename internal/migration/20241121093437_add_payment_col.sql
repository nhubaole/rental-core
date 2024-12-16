-- +goose Up
-- +goose StatementBegin
-- Xóa khóa ngoại hiện tại từ "bill_id" đến bảng "billing"
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_bill_id_fkey";

-- Cho phép "bill_id" nhận giá trị NULL
ALTER TABLE "payments" ALTER COLUMN "bill_id" DROP NOT NULL;

-- Thêm khóa ngoại mới từ "return_request_id" đến bảng "return_requests"
ALTER TABLE "payments" 
ADD COLUMN "return_request_id" INTEGER NULL,
ADD CONSTRAINT "payments_return_request_id_fkey" FOREIGN KEY ("return_request_id") REFERENCES "return_requests" ("id") ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Xóa khóa ngoại từ "return_request_id" đến bảng "return_requests"
ALTER TABLE "payments" DROP CONSTRAINT IF EXISTS "payments_return_request_id_fkey";

-- Xóa cột "return_request_id" khỏi bảng "payments"
ALTER TABLE "payments" DROP COLUMN IF EXISTS "return_request_id";

-- Thêm lại ràng buộc khóa ngoại cho "bill_id" đến bảng "billing"
ALTER TABLE "payments" 
ALTER COLUMN "bill_id" SET NOT NULL,
ADD CONSTRAINT "payments_bill_id_fkey" FOREIGN KEY ("bill_id") REFERENCES "billing" ("id") ON DELETE CASCADE;

-- +goose StatementEnd
