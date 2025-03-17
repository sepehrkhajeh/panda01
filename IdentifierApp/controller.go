package IdentifierApp

import (
	"Panda/dbconnection"
	"Panda/repositories"
	"Panda/schemas"
	"log"

	"github.com/labstack/echo/v4"
)

// هندلر ایجاد شناسه جدید
func CreateIdentifier(repo *repositories.IdentifierRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(IdentifierCreateRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		ctx := c.Request().Context()

		// دریافت Repository مربوط به دامنه‌ها
		dom, err := dbconnection.ConnectToRepo("domains")
		if err != nil {
			return c.JSON(500, map[string]string{"error": "مشکلی در سرور رخ داد"})
		}

		// بررسی مقدار برگشتی از ConnectToRepo
		if dom == nil {
			return c.JSON(500, map[string]string{"error": "Repository دامنه یافت نشد"})
		}

		// تبدیل نوع (Type Assertion)
		domainRepo, ok := dom.(*repositories.DomainRepository)
		if !ok {
			log.Println("نوع برگشتی با *DomainRepository مطابقت ندارد")
			return c.JSON(500, map[string]string{"error": "خطا در نوع Repository"})
		}

		// دریافت دامنه با شناسه
		d, err := domainRepo.GetByID(ctx, req.DomainID)
		if err != nil {
			log.Println("خطا در دریافت دامنه:", err)
			return c.JSON(500, map[string]string{"error": "خطا در دریافت اطلاعات دامنه"})
		}

		// بررسی وجود دامنه
		if d == nil {
			return c.JSON(400, map[string]string{"error": "دامنه یافت نشد"})
		}

		// ساخت شناسه کامل
		FullIdentifier := req.Name + "@" + d.Domain
		identifier, err := repo.Add(ctx, schemas.IdentifierSchema{
			Name:           req.Name,
			Pubkey:         req.Pubkey,
			DomainID:       req.DomainID,
			FullIdentifier: FullIdentifier,
		})

		if err != nil {
			log.Println("خطا در ایجاد شناسه:", err)
			return c.JSON(500, map[string]string{"error": "خطا در ایجاد شناسه"})
		}

		return c.JSON(200, identifier)
	}
}
