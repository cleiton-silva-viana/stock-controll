docker run --rm -v C:\Users\Inara\Documents\stock-controll:/app -w /app golangci/golangci-lint:v1.61.0 golangci-lint run -v
level=info msg="golangci-lint has version 1.61.0 built with go1.23.1 from a1d6c560 on 2024-09-09T17:44:42Z"
level=info msg="[config_reader] Config search paths: [./ /app / /root]"
level=info msg="[config_reader] Used config file .golangci.yml"
level=info msg="[lintersdb] Active 23 linters: [asasalint bodyclose cyclop dupl errcheck exhaustive funlen gci goconst gocritic godox gosec gosimple ineffassign maintidx misspell mnd prealloc revive unconvert unparam usestdlibvars whitespace]"
level=info msg="[loader] Go packages loading at mode 575 (name|deps|exports_file|files|imports|types_sizes|compiled_files) took 26.34914624s"
level=info msg="[runner/filename_unadjuster] Pre-built 0 adjustments in 111.499192ms"
level=info msg="[linters_context/goanalysis] analyzers took 23.503916049s with top 10 stages: buildir: 13.697156829s, the_only_name: 2.655971313s, exhaustive: 1.432517804s, inspect: 840.443037ms, dupl: 770.554299ms, gci: 760.486371ms, misspell: 536.923519ms, unparam: 509.35462ms, isgenerated: 429.75005ms, gocritic: 346.073703ms"
internal/domain/factory/credential.go:4:19: could not import stock-controll/internal/domain/entity/credential (-: # stock-controll/internal/domain/entity/credential
internal/domain/entity/credential/credentials.go:48:14: uuid.IsValid undefined (type string has no field or method IsValid)) (typecheck) 
        credentialEntity "stock-controll/internal/domain/entity/credential"
                         ^
internal/domain/factory/user.go:5:9: could not import stock-controll/internal/domain/entity/user (-: # stock-controll/internal/domain/entity/user
internal/domain/entity/user/user.go:136:59: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:155:12: undefined: validate.IsFutureDate
internal/domain/entity/user/user.go:155:34: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:156:56: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:157:55: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:177:12: undefined: validate.IsValueInRange) (typecheck)
        entity "stock-controll/internal/domain/entity/user"
               ^
internal/domain/factory/contact.go:14:47: too many arguments in call to contactEntity.NewContact
        have (string, string, string)
        want (string, string) (typecheck)
        return contactEntity.NewContact(uuid, email, phone)
                                                     ^
internal/domain/entity/promotion/promotion.go:21:2: could not import stock-controll/internal/domain/entity/product (-: # stock-controll/internal/domain/entity/product
internal/domain/entity/product/product.go:199:25: tag.Tag is not a type
internal/domain/entity/product/product.go:258:14: uuid.IsValid undefined (type string has no field or method IsValid)
internal/domain/entity/product/product.go:270:14: uuid.IsValid undefined (type string has no field or method IsValid)
internal/domain/entity/product/product.go:282:14: uuid.IsValid undefined (type string has no field or method IsValid)) (typecheck)       
        "stock-controll/internal/domain/entity/product"
        ^
internal/domain/entity/cart/cart.go:74:2: could not import stock-controll/internal/domain/entity/product (-: # stock-controll/internal/domain/entity/product
internal/domain/entity/product/product.go:199:25: tag.Tag is not a type
internal/domain/entity/product/product.go:258:14: uuid.IsValid undefined (type string has no field or method IsValid)
internal/domain/entity/product/product.go:270:14: uuid.IsValid undefined (type string has no field or method IsValid)
internal/domain/entity/product/product.go:282:14: uuid.IsValid undefined (type string has no field or method IsValid)) (typecheck)       
        "stock-controll/internal/domain/entity/product"
        ^
internal/domain/entity/cart/cart.go:204:25: productCart.UUID undefined (type *ProductCart has no field or method UUID) (typecheck)       
        c.products[productCart.UUID()] = *productCart
                               ^
internal/domain/entity/image/banner.go:1: : # stock-controll/internal/domain/entity/image
internal/domain/entity/image/image.go:111:128: missing return
internal/domain/entity/image/image.go:117:63: missing return (typecheck)
package image
internal/domain/entity/report/discard/discard.go:51:4: expected selector or type assertion, found ')' (typecheck)
                        )
                        ^
internal/domain/entity/report/discard/discard.go:59:97: missing ',' before newline in argument list (typecheck)
                                validation.IsInRange(MinQuantityForDiscard, MaxQuantityForDiscard, validation.ErrUnknown)))
                                                                                                                           ^
internal/domain/entity/report/discard/discard.go:61:2: expected operand, found 'if' (typecheck)
        if discardError.HasError() {
        ^
internal/domain/entity/report/discard/discard.go:62:3: expected operand, found 'return' (typecheck)
                return nil, discardError
                ^
internal/domain/entity/report/discard/discard.go:63:2: expected operand, found '}' (typecheck)
        }
        ^
internal/domain/entity/report/discard/discard.go:65:2: missing ',' in argument list (typecheck)
        return &discard{
        ^
internal/domain/entity/report/discard/discard.go:73:9: missing ',' before newline in argument list (typecheck)
        }, nil
              ^
internal/domain/entity/report/discard/discard.go:74:1: expected operand, found '}' (typecheck)
}
^
internal/domain/entity/report/discard/discard.go:77:2: missing ',' in argument list (typecheck)
        return d.uuid
        ^
internal/domain/entity/report/discard/discard.go:78:1: expected operand, found '}' (typecheck)
}
^
internal/domain/entity/report/discard/discard.go:81:2: missing ',' in argument list (typecheck)
        return d.checkerUUID
        ^
internal/domain/entity/report/order/order.go:37:120: undefined: validationError.IValidationError (typecheck)
func NewOrder(buyerUUID, supplierUUID string, expectedDelivery time.Time, products []Product) (*Order, validationError.IValidationError) 
{
                                                                                                                       ^
internal/domain/entity/report/order/order.go:38:35: undefined: validationError.NewValidationError (typecheck)
        var orderError = validationError.NewValidationError("order")
                                         ^
internal/domain/entity/report/order/order.go:50:36: orderInstance.ValidateUUID undefined (type *Order has no field or method ValidateUUID) (typecheck)
                AddValidationError(orderInstance.ValidateUUID(buyerUUID)).
                                                 ^
internal/domain/entity/report/order/order.go:51:36: orderInstance.ValidateUUID undefined (type *Order has no field or method ValidateUUID) (typecheck)
                AddValidationError(orderInstance.ValidateUUID(supplierUUID)).
                                                 ^
level=info msg="[runner/max_same_issues] 1/4 issues with text \"could not import stock-controll/internal/domain/entity/product (-: # stock-controll/internal/domain/entity/product\\ninternal/domain/entity/product/product.go:199:25: tag.Tag is not a type\\ninternal/domain/entity/product/product.go:258:14: uuid.IsValid undefined (type string has no field or method IsValid)\\ninternal/domain/entity/product/product.go:270:14: uuid.IsValid undefined (type string has no field or method IsValid)\\ninternal/domain/entity/product/product.go:282:14: uuid.IsValid undefined (type string has no field or method IsValid))\" were hidden, use --max-same-issues"
level=info msg="[runner/max_from_linter] 11/61 issues from linter typecheck were hidden, use --max-issues-per-linter"
level=info msg="[runner] Issues before processing: 4331, after processing: 50"
level=info msg="[runner] Processors filtering stat (in/out): nolint: 4113/4113, uniq_by_line: 4113/62, max_from_linter: 61/50, sort_results: 50/50, cgo: 4331/4331, invalid_issue: 4331/4113, skip_dirs: 4113/4113, autogenerated_exclude: 4113/4113, exclude: 4113/4113, max_same_issues: 62/61, severity-rules: 50/50, fixer: 50/50, source_code: 50/50, path_shortener: 50/50, path_prefixer: 50/50, filename_unadjuster: 4331/4331, identifier_marker: 4113/4113, exclude-rules: 4113/4113, max_per_file_from_linter: 62/62, path_prettifier: 4113/4113, skip_files: 4113/4113, diff: 62/62"
level=info msg="[runner] processing took 527.012121ms with stages: path_prettifier: 259.618823ms, identifier_marker: 195.071281ms, source_code: 68.30786ms, filename_unadjuster: 717.751µs, cgo: 704.279µs, invalid_issue: 550.159µs, uniq_by_line: 462.71µs, exclude-rules: 395.488µs, nolint: 348.708µs, skip_dirs: 322.637µs, autogenerated_exclude: 314.892µs, max_same_issues: 132.267µs, path_shortener: 39.197µs, max_from_linter: 19.157µs, max_per_file_from_linter: 4.158µs, fixer: 771ns, exclude: 571ns, diff: 451ns, skip_files: 370ns, sort_results: 360ns, path_prefixer: 121ns, severity-rules: 110ns"
level=info msg="[runner] linters took 6.424390851s with stages: goanalysis_metalinter: 5.897230699s"
internal/domain/entity/report/sale/sale.go:65:158: undefined: validationError.IValidationError (typecheck)
func NewSale(clientUUID, sellerUUID string, products []Product, payment PaymentMethod, discountApplyed Discount, status salesStatus) (*Sale, validationError.IValidationError) {

                    ^
internal/domain/entity/report/sale/sale.go:129:16: field and method with the same name PaymentMethod (typecheck)
func (s *Sale) PaymentMethod() string {
               ^
internal/domain/entity/report/sale/sale.go:59:2: other declaration of PaymentMethod (typecheck)
        PaymentMethod PaymentMethod
        ^
internal/domain/entity/report/sale/sale.go:66:34: undefined: validationError.NewValidationError (typecheck)
        var saleError = validationError.NewValidationError("sale")
                                        ^
internal/domain/entity/report/sale/sale.go:130:16: cannot convert s.PaymentMethod (value of type func() string) to type string (typecheck)
        return string(s.PaymentMethod)
                      ^
internal/application/feature/user/abstractions_test.go:16:2: unknown field Name in struct literal of type dto.CreateUserRequestDTO (typecheck)
        Name:      fake.Person().FirstName(),
        ^
internal/application/feature/user/abstractions_test.go:17:2: unknown field Gender in struct literal of type dto.CreateUserRequestDTO (typecheck)
        Gender:    "male",
        ^
internal/application/feature/user/abstractions_test.go:18:13: cannot use "1989-05-25" (untyped string constant) as time.Time value in struct literal (typecheck)
        BirthDate: "1989-05-25",
                   ^
internal/application/feature/user/abstractions_test.go:25:45: userRequestDTO.Name undefined (type *dto.CreateUserRequestDTO has no field 
or method Name) (typecheck)
var user, _ = entity.NewUser(userRequestDTO.Name, userRequestDTO.Gender, userRequestDTO.BirthDate)
                                            ^
internal/application/feature/user/create_user.go:77:3: unknown field UID in struct literal of type dto.CreateUserResponseDTO (typecheck) 
                UID:    userUID,
                ^
internal/application/feature/user/create_user.go:78:19: userDTO.Name undefined (type dto.UserDTO has no field or method Name) (typecheck)                Name:   userDTO.Name,
                                ^
internal/application/feature/user/create_user.go:79:3: unknown field Gender in struct literal of type dto.CreateUserResponseDTO (typecheck)
                Gender: userDTO.Gender,
                ^
internal/application/feature/user/create_user.go:142:18: DTO.Name undefined (type dto.CreateUserRequestDTO has no field or method Name) (typecheck)
                Name:      DTO.Name,
                               ^
internal/application/feature/user/create_user.go:143:3: unknown field Gender in struct literal of type dto.UserDTO (typecheck)
                Gender:    DTO.Gender,
                ^
internal/application/feature/user/create_user_test.go:45:9: dto.Name undefined (type *dto.CreateUserRequestDTO has no field or method Name) (typecheck)
                                dto.Name = "invalid.name_for_user@"
                                    ^
internal/application/feature/user/update_user_test.go:17:3: unknown field Name in struct literal of type dto.UserDTO (typecheck)
                Name:      fake.Person().FirstName(),
                ^
internal/application/feature/user/update_user_test.go:18:3: unknown field Gender in struct literal of type dto.UserDTO (typecheck)       
                Gender:    "Male",
                ^
internal/application/feature/user/update_user_test.go:19:14: cannot use "1999-05-01" (untyped string constant) as time.Time value in struct literal (typecheck)
                BirthDate: "1999-05-01",
                           ^
internal/domain/entity/brand/brand.go:1: : # stock-controll/internal/domain/entity/brand [stock-controll/internal/domain/entity/brand.test]
internal/domain/entity/brand/brand.go:115:14: uuid.IsValid undefined (type string has no field or method IsValid) (typecheck)
package brand
internal/domain/entity/commpany/supplier.go:9:2: could not import stock-controll/internal/domain/entity/user (-: # stock-controll/internal/domain/entity/user
internal/domain/entity/user/user.go:136:59: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:155:12: undefined: validate.IsFutureDate
internal/domain/entity/user/user.go:155:34: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:156:56: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:157:55: undefined: validate.ErrUnknown
internal/domain/entity/user/user.go:177:12: undefined: validate.IsValueInRange) (typecheck)
        "stock-controll/internal/domain/entity/user"
        ^
internal/domain/entity/commpany/supplier.go:18:146: undefined: validationError.IValidationError (typecheck)
func NewSupplier(name, cnpj, billingEmail, billingPhone, purchaseEmail, purchasePhone string, addr address.IAddress) (*Supplier, validationError.IValidationError) {

        ^
internal/domain/entity/commpany/supplier.go:19:39: undefined: validationError.NewValidationError (typecheck)
        var supplierErrors = validationError.NewValidationError("supplier")
                                             ^
internal/domain/entity/commpany/supplier.go:29:39: supplierInstance.SetName undefined (type Supplier has no field or method SetName) (typecheck)
                AddValidationError(supplierInstance.SetName(name)).
                                                    ^
internal/domain/entity/commpany/supplier.go:30:39: supplierInstance.SetCNPJ undefined (type Supplier has no field or method SetCNPJ) (typecheck)
                AddValidationError(supplierInstance.SetCNPJ(cnpj)).
                                                    ^
internal/domain/entity/commpany/supplier.go:31:39: supplierInstance.SetBillingContact undefined (type Supplier has no field or method SetBillingContact) (typecheck)
                AddValidationError(supplierInstance.SetBillingContact(billingEmail, billingPhone)).
                                                    ^
internal/domain/entity/commpany/supplier.go:32:39: supplierInstance.SetPurchaseContact undefined (type Supplier has no field or method SetPurchaseContact) (typecheck)
                AddValidationError(supplierInstance.SetPurchaseContact(purchaseEmail, purchasePhone)).
                                                    ^
internal/domain/entity/commpany/supplier.go:33:39: supplierInstance.SetAddress undefined (type Supplier has no field or method SetAddress) (typecheck)
                AddValidationError(supplierInstance.SetAddress(addr))
                                                    ^
internal/domain/entity/contact/contact.go:1: : # stock-controll/internal/domain/entity/contact [stock-controll/internal/domain/entity/contact.test]
internal/domain/entity/contact/contact_test.go:56:19: cannot use uuid.IsValid(contact.UUID()) (value of type *validate.FieldError) as bool value in argument to assert.True
internal/domain/entity/contact/contact_test.go:56:32: not enough arguments in call to uuid.IsValid
        have (string)
        want (string, string) (typecheck)
package contact
level=info msg="File cache stats: 41 entries of total size 96.6KiB"
level=info msg="Memory: 331 samples, avg is 82.3MB, max is 456.2MB"
level=info msg="Execution took 32.905886798s"