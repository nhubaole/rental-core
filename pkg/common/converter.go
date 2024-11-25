package common

import (
	"mime/multipart"
	"reflect"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func MapStruct(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)
	dstVal := reflect.ValueOf(dst).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldType := srcType.Field(i)
		dstField := dstVal.FieldByName(srcFieldType.Name)

		// If the destination struct has a corresponding field and it's settable
		if dstField.IsValid() && dstField.CanSet() {
			switch srcField.Kind() {
			case reflect.String:
				dstField.SetString(srcField.String())

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				dstField.SetInt(srcField.Int())

			case reflect.Float32, reflect.Float64:
				dstField.SetFloat(srcField.Float())

			case reflect.Bool:
				dstField.SetBool(srcField.Bool())

			case reflect.Ptr:
				if !srcField.IsNil() {
					if srcField.Type().Elem().Name() == "FileHeader" {
						fileHeader := srcField.Interface().(*multipart.FileHeader)
						if fileHeader != nil {
							fileName := fileHeader.Filename            // Extract filename
							fileNamePtr := &fileName                   // Create a *string
							dstField.Set(reflect.ValueOf(fileNamePtr)) // Set as *string
						} else {
							// Handle nil case
							dstField.Set(reflect.Zero(dstField.Type()))
						}
					} else {
						dstField.Set(srcField)

					}
				}

			case reflect.Slice:
				// Check if it's a slice of strings
				if srcField.Type().Elem().Kind() == reflect.String {
					dstField.Set(srcField)
				}
				// Handle file uploads (single or multiple)
				// Handle file uploads (slice of *multipart.FileHeader)
				if srcField.Type().Elem().Kind() == reflect.Ptr && srcField.Type().Elem().Elem().Name() == "FileHeader" {
					// Extract filenames without saving the files
					var fileNames []string
					for j := 0; j < srcField.Len(); j++ {
						fileHeader := srcField.Index(j).Interface().(*multipart.FileHeader)
						fileNames = append(fileNames, fileHeader.Filename) // Just collect filenames
					}
					// Set the filenames to the destination field
					dstField.Set(reflect.ValueOf(fileNames))
				}

			}

		}
	}
	return nil
}

func Int64ToPgTimestamptz(unixTimestamp int64, isMilliseconds bool) pgtype.Timestamptz {
    var ts pgtype.Timestamptz
    var t time.Time

    if isMilliseconds {
        t = time.Unix(0, unixTimestamp*int64(time.Millisecond))
    } else {
        t = time.Unix(unixTimestamp, 0)
    }

    ts.Time = t          // Set the time value
    ts.Valid = true      // Set the status to valid
    return ts
}


