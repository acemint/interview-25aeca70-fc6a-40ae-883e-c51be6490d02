"use client";
import { Button } from "@/components/ui/button";
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
  FieldLegend,
  FieldSeparator,
  FieldSet
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import AircraftSelectionField from "@/app/aircraft-selection-field";
import { FlightDatePicker } from "@/app/flight-date-picker";
import { useState } from "react";
import { toast } from "sonner";
import { CheckVoucher, GenerateVoucher } from "@/app/api-endpoint";

function VoucherGeneratorForm() {
  const [crewName, setCrewName] = useState("");
  const [crewId, setCrewId] = useState("");
  const [flightNumber, setFlightNumber] = useState("");
  const [date, setDate] = useState("");
  const [aircraftType, setAircraftType] = useState("");
  const [seats, setSeats] = useState<undefined | string[]>(undefined);

  const [flightNumberResponse, setFlightNumberResponse] = useState("");
  const [dateResponse, setDateResponse] = useState("");


  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    try {
      const checkResponse = await check();
      
      if (checkResponse.exists) {
        toast.info("Voucher has already been generated for this date");
        return;
      }
      const generateResponse = await generate();
      setFlightNumberResponse(flightNumber);
      setDateResponse(date);
      setSeats(generateResponse.seats);
      toast.success("Successfully generated vouchers");
    } catch (error) {
      console.error("Submission failed caused by: ", error);
    }
  };

  const check = async () => {
    const response = await CheckVoucher(flightNumber, date);
    if (!response.success) {
      toast.error(`Error: ${response.status} caused by ${JSON.stringify(response.error)}`);
      throw new Error(response.error);
    }
    return response.data;
  };

  const generate = async () => {
    const response = await GenerateVoucher(crewName, crewId, flightNumber, date, aircraftType);
    if (!response.success) {
      toast.error(`Error: ${response.status} caused by ${JSON.stringify(response.error)}`);
      throw new Error(response.error);
    }
    return response.data;
  };

  return (
    <div className="w-full max-w-md">
      <form onSubmit={handleSubmit}>
        <FieldGroup>
          <FieldSet>
            <FieldLegend>Voucher Generator</FieldLegend>
            <FieldDescription></FieldDescription>
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="crew-name">Crew Name</FieldLabel>
                <Input id="crew-name" placeholder="Steven Kristian" required
                       onChange={(e) => setCrewName(e.target.value)} />
              </Field>
              <Field>
                <FieldLabel htmlFor="crew-id">Crew ID</FieldLabel>
                <Input id="crew-id" placeholder="10001" required
                       onChange={(e) => setCrewId(e.target.value)} />
              </Field>
              <Field>
                <FieldLabel htmlFor="flight-number">Flight Number</FieldLabel>
                <Input id="flight-number" placeholder="ID102" required
                       onChange={(e) => setFlightNumber(e.target.value)} />
              </Field>
              <AircraftSelectionField onValueChange={setAircraftType} />
              <Field>
                <FlightDatePicker onValueChange={setDate} />
              </Field>

              <Field orientation="horizontal">
                <Button type="submit">Generate Vouchers</Button>
              </Field>
            </FieldGroup>
          </FieldSet>
          <FieldSeparator />
          {seats && <FieldSet>
            <FieldLegend>Seats</FieldLegend>
            <FieldDescription>
              The seats that is generated
            </FieldDescription>
            <div className="flex flex-col gap-2 text-sm">
              <p>Flight Number: {flightNumberResponse}</p>
              <p>Date: {dateResponse}</p>
              <p>Seats:
                {seats.map((seat, index) => (
                  <span key={index} className="mr-2">{seat}</span>
                ))}
              </p>
            </div>
          </FieldSet>}
        </FieldGroup>
      </form>
    </div>
  );
}

export default VoucherGeneratorForm;