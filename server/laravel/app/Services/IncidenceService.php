<?php
namespace App\Services;

use App\Models\Incidence;
use App\Models\IncidencePhoto;

class IncidenceService
{
    protected Incidence $incidence;
    // protected IncidencePhoto $incidencePhoto;

    public function __construct(Incidence $incidence, IncidencePhoto $incidencePhoto)
    {
        $this->incidence = $incidence;
        $this->incidencePhoto = $incidencePhoto;
    }

    public function syncPhotos(Incidence $incidence, array $photos): void
    {

        foreach ($photos as $photo) {
            // $tagsIds[] = $this->tag->firstOrCreate(['name' => $tag])->id;
        }
    }
}
