import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Field, FieldLabel } from "@/components/ui/field";


function AircraftSelectionField({ onValueChange }: { onValueChange: (value: string) => void }) {
  // I think this would be better if we hit an API to get list of aircraft rather than dropdown.
  // It is because we'd need deployment for new aircraft. If we put it in DB instead it will be more "safer"
  // However, as the requirement doesn't mention anything about aircraft validation, I think we can just hard code this
  return (
    <Field>
      <FieldLabel htmlFor="aircraft-type">Aircraft Type</FieldLabel>
      <Select defaultValue="" required onValueChange={onValueChange}>
        <SelectTrigger id="aircraft-type">
          <SelectValue placeholder="ATR" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="ATR">ATR</SelectItem>
          <SelectItem value="Airbus 320">Airbus 320</SelectItem>
          <SelectItem value="Boeing 737 Max">Boeing 737 Max</SelectItem>
        </SelectContent>
      </Select>
    </Field>
  );
}

export default AircraftSelectionField;