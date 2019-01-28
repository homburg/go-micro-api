package scripts

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/byblix/go-micro/models"
	"github.com/byblix/go-micro/services"
)

// WithdrawalsToCSV asdasd
func WithdrawalsToCSV(dbRef string) {
	ref, ctx := services.GetDBRefAndCTX(os.Getenv("ENV") + "/" + dbRef)
	result, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalf("Error initializing db %s\n", err)
	}

	csvWriter := services.CreateCSVWriterFile("withdrawals_not_completed_" + os.Getenv("ENV") + ".csv")

	for idx, r := range result {
		// TODO: Change this according to generic model
		var value models.Withdrawals
		if err := r.Unmarshal(&value); err != nil {
			log.Fatalln("Error unmarshaling result:", err)
		}

		// Check if the request is not completed
		// if !value.RequestCompleted {
		writeWithdrawalsToCSV(csvWriter, idx, value)
		// }
	}
	defer fmt.Println("Done writing CSV!")
}

// WriteWithdrawalsToCSV Does everything inside the loop above
func writeWithdrawalsToCSV(w *csv.Writer, index int, val models.Withdrawals) {
	fmt.Println("The userID: ", val.RequestUserID)
	profile := services.GetFBDataByChild("profiles", val.RequestUserID)
	var record []string
	record = append(record, strconv.FormatInt(int64(index), 10))
	record = append(record, services.ParseUnixAsDate(val.RequestDate))
	record = append(record, profile.DisplayName)
	record = append(record, strconv.FormatInt(val.RequestAmount, 10))
	record = append(record, services.ParseUnixAsDate(val.RequestCompletedDate))
	record = append(record, "Udbetalt? : "+strconv.FormatBool(val.RequestCompleted))
	err := w.Write(record)
	if err != nil {
		log.Printf("Error writing result: %s", err)
	}
	fmt.Printf("\n----\nConverting key: %v with profileName: %v\n", index, profile.DisplayName)
	w.Flush()
}

// ProfilesToCSV initiates data and loops the write process
func ProfilesToCSV(dbRef string) {
	ref, ctx := services.GetDBRefAndCTX(os.Getenv("ENV") + "/" + dbRef)
	result, err := ref.OrderByKey().GetOrdered(ctx)
	if err != nil {
		log.Fatalf("Error initializing db %s\n", err)
	}

	csvWriter := services.CreateCSVWriterFile("profile_withdrawable_" + os.Getenv("ENV") + ".csv")
	rowNames := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	services.WriteColumnHeaders(rowNames, csvWriter)

	index := int64(1)
	for _, r := range result {
		var profile models.Profile
		if err := r.Unmarshal(&profile); err != nil {
			log.Printf("Error unmarshaling result: %s", err)
		}

		// TODO: Calc the total amount
		// TODO: Set the CSV header column name
		if profile.SalesAmount > 0 {
			writeProfilesToCSV(csvWriter, index, profile)
			index++
		}
	}
	defer fmt.Println("Done writing CSV!")
}

// WriteProfilesToCSV Does everything inside the loop above
func writeProfilesToCSV(csvWriter *csv.Writer, index int64, profile models.Profile) {
	alreadyCashedAmount := (profile.SalesAmount - profile.WithdrawableAmount)
	var record []string
	record = append(record, strconv.FormatInt(index, 10))
	record = append(record, profile.DisplayName)
	record = append(record, profile.FirstName+" "+profile.LastName)
	record = append(record, profile.Email)
	record = append(record, strconv.FormatBool(profile.IsProfessional))
	record = append(record, services.ParseUnixAsDate(profile.CreateDate))
	record = append(record, strconv.FormatInt(profile.SalesQuantity, 10))
	record = append(record, strconv.FormatInt(profile.WithdrawableAmount, 10))
	record = append(record, strconv.FormatInt(alreadyCashedAmount, 10))
	record = append(record, strconv.FormatInt(profile.SalesAmount, 10))
	err := csvWriter.Write(record)
	if err != nil {
		log.Printf("Error writing result: %s", err)
	}
	fmt.Printf("\n----\nConverting key: %v with profileName: %v\n", index, profile.DisplayName)
	csvWriter.Flush()
}
