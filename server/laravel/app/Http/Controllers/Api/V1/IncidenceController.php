<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use App\Services\IncidenceService;
use App\Http\Requests\IncidenceRequest;
use Illuminate\Http\Request;

use App\Repositories\IncidenceRepository;
class IncidenceController extends Controller
{

    public function __construct(IncidenceRepository $IncidenceRepository, IncidenceService $incidenceService)
    {
        $this->incidenceService = $incidenceService;
        $this->incidenceRepository = $IncidenceRepository;
    }
    public function store(Request $request) {
        $request->validate([
            'file' => 'required',
            'file.*' => 'mimes:jpeg,jpg,png|max:2048'
          ]);
        //   foreach ($request->file('file') as $file) { // try here ]);
        //     $name = $file->getClientOriginalName();
        //     $file->move(public_path().'/uploads/', $name);
        // }
        $createIncidence = $this->incidenceRepository->createIncidence();
        $this->incidenceService->syncPhotos($createIncidence, $request->file('file'));
        return response()->json(['status' =>'success', 'data' => $createIncidence], 201);

    }

}
