<?php

namespace App\Http\Controllers\Api\V1;

use App\Http\Controllers\Controller;
use App\Services\IncidenceService;
use App\Http\Requests\IncidenceRequest;
use Illuminate\Http\Request;
use App\Models\Incidence;
use Illuminate\Support\Facades\Gate;
use App\Repositories\IncidenceRepository;
use App\Traits\ErrorTrait;
use App\Traits\UserTrait;
use Illuminate\Auth\Middleware\Authorize;

class IncidenceController extends Controller
{
    use ErrorTrait;
    use UserTrait;
    public function __construct(IncidenceRepository $IncidenceRepository, IncidenceService $incidenceService)
    {
        $this->incidenceService = $incidenceService;
        $this->incidenceRepository = $IncidenceRepository;
    }
    public function list() {
        return response()->json(['status' =>'success', 'data' => $this->incidenceRepository->list()], 201);
    }
    public function close() {
        return response()->json(['status' =>'success', 'data' => $this->incidenceRepository->update()], 201);
    }
    public function store(Request $request) {
        $request->validate([
            'file' => 'required',
            'file.*' => 'mimes:jpeg,jpg,png|max:2048'
          ]);
        $createIncidence = $this->incidenceRepository->createIncidence();
        $this->incidenceService->syncPhotos($createIncidence, $request->file('file'));
        return response()->json(['status' =>'success', 'data' => 'succes created'], 201);
    }
    public function delete() {
        try {
            if (self::checkId(Incidence::firstWhere('id',request()->id)->owner)) {
                return response()->json(['status' =>'success', 'data' => Incidence::firstWhere('id',request()->id)->delete()], 201);
            } else {
                return response()->json(['status' =>'success',false, 'data' => "Not allowed"], 201);
            }
        } catch (\Throwable $th) {
            self::throwError();
        }
    }
}
