import Button from "@/components/Button";
const databases = ["users", "mflix", "system"];

export default function Dashboard() {
  return (
    <>
      <button
        data-drawer-target="default-sidebar"
        data-drawer-toggle="default-sidebar"
        aria-controls="default-sidebar"
        type="button"
        className="inline-flex items-center p-2 mt-2 ml-3 text-sm text-gray-500 rounded-lg sm:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
      >
        <span className="sr-only">Open sidebar</span>
        <svg
          className="w-6 h-6"
          aria-hidden="true"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            clip-rule="evenodd"
            fill-rule="evenodd"
            d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"
          ></path>
        </svg>
      </button>

      <aside
        id="default-sidebar"
        className="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
        aria-label="Sidebar"
      >
        <div className="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
          <ul className="space-y-2 font-medium">
            {databases.map((database) => {
              return (
                <li key={database}>
                  <a
                    href="#"
                    className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                  >
                    <svg
                      className="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
                      aria-hidden="true"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      viewBox="0 0 22 21"
                    >
                      <path d="M16.975 11H10V4.025a1 1 0 0 0-1.066-.998 8.5 8.5 0 1 0 9.039 9.039.999.999 0 0 0-1-1.066h.002Z" />
                      <path d="M12.5 0c-.157 0-.311.01-.565.027A1 1 0 0 0 11 1.02V10h8.975a1 1 0 0 0 1-.935c.013-.188.028-.374.028-.565A8.51 8.51 0 0 0 12.5 0Z" />
                    </svg>
                    <span className="ml-3">{database}</span>
                  </a>
                </li>
              );
            })}
          </ul>
        </div>
      </aside>

      <div className="p-4 sm:ml-64">
        <div className="p-4 border-2 border-gray-200 border-dashed rounded-lg color-gray flex flex-col justify-center content-center items-center ">
          <h2>Profile your database</h2>
          <p>
            It seems your database not indexed yet, click start button to enable
            profiling
          </p>
          <Button
            isLoading={false}
            type="submit"
            className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Start Profiling
          </Button>
        </div>
        <div className="container mx-auto bg-gray-50 min-h-screen p-8 antialiased">
   <div>
      <div className="bg-gray-100 mx-auto border-gray-500 border rounded-sm text-gray-700 mb-0.5 h-30">
         <div className="flex p-3 border-l-8 border-yellow-600">
            <div className="space-y-1 border-r-2 pr-3">
               <div className="text-sm leading-5 font-semibold"><span className="text-xs leading-4 font-normal text-gray-500"> Release #</span> LTC08762304</div>
               <div className="text-sm leading-5 font-semibold"><span className="text-xs leading-4 font-normal text-gray-500 pr"> BOL #</span> 10937</div>
               <div className="text-sm leading-5 font-semibold">JUN 14. 9:30 PM</div>
            </div>
            <div className="flex-1">
               <div className="ml-3 space-y-1 border-r-2 pr-3">
                  <div className="text-base leading-6 font-normal">KROGER MEMPHIS</div>
                  <div className="text-sm leading-4 font-normal"><span className="text-xs leading-4 font-normal text-gray-500"> Carrier</span> PAPER TRANSPORT INC.</div>
                  <div className="text-sm leading-4 font-normal"><span className="text-xs leading-4 font-normal text-gray-500"> Destination</span> WestRock Jacksonville - 9469 Eastport Rd, Jacksonville, FL 32218</div>
               </div>
            </div>
            <div className="border-r-2 pr-3">
               <div >
                  <div className="ml-3 my-3 border-gray-200 border-2 bg-gray-300 p-1">
                     <div className="uppercase text-xs leading-4 font-medium">Trailer</div>
                     <div className="text-center text-sm leading-4 font-semibold text-gray-800">89732</div>
                  </div>
               </div>
            </div>
            <div>
               <div className="ml-3 my-5 bg-yellow-600 p-1 w-20">
                  <div className="uppercase text-xs leading-4 font-semibold text-center text-yellow-100">Loaded</div>
               </div>
            </div>
            <div>
               <button className="text-gray-100 rounded-sm my-5 ml-2 focus:outline-none bg-gray-500">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
               </button>
            </div>
         </div>
      </div>

      <div className="bg-gray-100 mx-auto border-gray-500 border rounded-sm  text-gray-700 mb-0.5">
         <div className="flex p-3  border-l-8 border-green-600">
            <div className="space-y-1 border-r-2 pr-3">
               <div className="text-sm leading-5 font-semibold"><span className="text-xs leading-4 font-normal text-gray-500"> Release #</span> LTC08762304</div>
               <div className="text-sm leading-5 font-semibold"><span className="text-xs leading-4 font-normal text-gray-500 pr"> BOL #</span> 10937</div>
               <div className="text-sm leading-5 font-semibold">JUN 14. 9:30 PM</div>
            </div>
            <div className="flex-1">
               <div className="ml-3 space-y-1 border-r-2 pr-3">
                  <div className="text-base leading-6 font-normal">KROGER MEMPHIS</div>
                  <div className="text-sm leading-4 font-normal"><span className="text-xs leading-4 font-normal text-gray-500"> Carrier</span> PAPER TRANSPORT INC.</div>
                  <div className="text-sm leading-4 font-normal"><span className="text-xs leading-4 font-normal text-gray-500"> Destination</span> WestRock Jacksonville - 9469 Eastport Rd, Jacksonville, FL 32218</div>
               </div>
            </div>
            <div className="border-r-2 pr-3">
               <div >
                  <div className="ml-3 my-3 border-gray-200 border-2 bg-gray-300 p-1">
                     <div className="uppercase text-xs leading-4 font-medium">Trailer</div>
                     <div className="text-center text-sm leading-4 font-semibold text-gray-800">89732</div>
                  </div>
               </div>
            </div>
            <div>
               <div className="ml-3 my-5 bg-green-600 p-1 w-20">
                  <div className="uppercase text-xs leading-4 font-semibold text-center text-green-100">Picked UP</div>
               </div>
            </div>
            <div>
               <button className="text-gray-100 rounded-sm my-5 ml-2 focus:outline-none bg-gray-500">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
               </button>
            </div>
         </div>
      </div>

      <div className="bg-gray-100 mx-auto border-gray-500 border rounded-sm  text-gray-700 mb-0.5">
         <div className="flex p-3  border-l-8 border-red-600">
            <div className="space-y-1 border-r-2 pr-3">
               <div className="text-sm leading-5 font-semibold"><span className="text-xs leading-4 font-normal text-gray-500"> Release #</span> LTC08762304</div>
               <div className="text-sm leading-5 font-semibold"><span className="text-xs leading-4 font-normal text-gray-500 pr"> BOL #</span> 10937</div>
               <div className="text-sm leading-5 font-semibold">JUN 14. 9:30 PM</div>
            </div>
            <div className="flex-1">
               <div className="ml-3 space-y-1 border-r-2 pr-3">
                  <div className="text-base leading-6 font-normal">KROGER MEMPHIS</div>
                  <div className="text-sm leading-4 font-normal"><span className="text-xs leading-4 font-normal text-gray-500"> Carrier</span> PAPER TRANSPORT INC.</div>
                  <div className="text-sm leading-4 font-normal"><span className="text-xs leading-4 font-normal text-gray-500"> Destination</span> WestRock Jacksonville - 9469 Eastport Rd, Jacksonville, FL 32218</div>
               </div>
            </div>
            <div className="border-r-2 pr-3">
               <div >
                  <div className="ml-3 my-3 border-gray-200 border-2 bg-gray-300 p-1">
                     <div className="uppercase text-xs leading-4 font-medium">Trailer</div>
                     <div className="text-center text-sm leading-4 font-semibold text-gray-800">89732</div>
                  </div>
               </div>
            </div>
            <div>
               <div className="ml-3 my-5 bg-red-600 p-1 w-20">
                  <div className="uppercase text-xs leading-4 font-semibold text-center text-red-100">Canceled</div>
               </div>
            </div>
            <div>
               <button className="text-gray-100 rounded-sm my-5 ml-2 focus:outline-none bg-gray-500">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                  </svg>
               </button>
            </div>
         </div>
      </div>
   </div>
</div>
      </div>
    </>
  );
}
