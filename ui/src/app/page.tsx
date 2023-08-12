"use client";

import Image from "next/image";
import { useForm } from "react-hook-form";
import Button from "@/components/Button";
import { useState } from "react";
import dayjs from "dayjs";

type FormData = {
  connectionUrl: string;
};

const connectionsHistory = [
  {
    host: "localhost",
    lastConnect: new Date()
  },
  {
    host: "192.168.1.1",
    lastConnect: new Date()
  }
]

export default function Home() {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setError,
    watch,
  } = useForm<FormData>();
  const connectionUrl = watch("connectionUrl");
  const [isLoading, setLoading] = useState(false);
  const onSubmit = handleSubmit(() => {
    if (connectionUrl) {
      try {
        new URL(connectionUrl);
      } catch (e: any) {
        setError("connectionUrl", {
          message: e.message,
        });

        return;
      }
    } else {
      setError("connectionUrl", {
        message: "url required",
      });

      return;
    }

    setLoading(true);
    setTimeout(() => {
      setLoading(false);
    }, 2000);
  });

  return (
    <>
      <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-sm">
          <img
            className="mx-auto h-10 w-auto"
            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
            alt="Your Company"
          />
          <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
            Connect to your database
          </h2>
        </div>

        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <form className="space-y-6" onSubmit={onSubmit}>
            <div>
              <div className="mt-2">
                <input
                  placeholder="mongodb://user:password@localhost:27017/mflix"
                  className="block w-full rounded-md border-0 py-1.5 px-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  {...register("connectionUrl")}
                />
              </div>
              {errors.connectionUrl && (
                <p className="text-red-500 text-xs italic">
                  connection url required !
                </p>
              )}
            </div>
            <div>
              <Button
                isLoading={isLoading}
                type="submit"
                className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >
                Connect
              </Button>
            </div>
          </form>
        </div>

        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <h4 className="my-2 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">Recent connection</h4>
          <div className="w-full text-sm font-medium text-gray-900 bg-white border border-gray-200 rounded-lg">
            {
              connectionsHistory.map((c) => (
                <button type="button" key={c.host} aria-current="true" className="flex w-full px-4 py-2 border-b border-gray-200 rounded-t-lg cursor-pointer hover:bg-gray-100 hover:text-indigo-700">
                    <div className="flex-1 font-bold text-left">{c.host}</div>
                    <div className="flx-1 text-right text-xs">{dayjs(c.lastConnect).format("MMM D, YYYY, HH:mm")}</div>
                </button>
              ))
            }
          </div>
        </div>
      </div>
    </>
  );
}
