"use server";

import { env } from "@/env.mjs";

const GenerateVoucher = async (crewName: string,
                               crewId: string,
                               flightNumber: string,
                               date: string,
                               aircraftType: string) => {
  const backendUrl = env.BACKEND_URL;
  const response = await fetch(`${backendUrl}/api/generate`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      name: crewName,
      id: crewId,
      flightNumber: flightNumber,
      date: date,
      aircraft: aircraftType
    })
  });

  if (!response.ok) {
    const errorData = await response.json();
    return {
      success: false,
      error: errorData.error || "Unknown error",
      status: response.status
    };
  }

  return { success: true, data: await response.json() };
};

const CheckVoucher = async (flightNumber: string, date: string) => {
  const backendUrl = env.BACKEND_URL;
  const response = await fetch(`${backendUrl}/api/check`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      flightNumber: flightNumber,
      date: date
    })
  });

  if (!response.ok) {
    const errorData = await response.json();
    return {
      success: false,
      error: errorData.error || "Unknown error",
      status: response.status
    };
  }

  return { success: true, data: await response.json() };
};

export { GenerateVoucher, CheckVoucher };