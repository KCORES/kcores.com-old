<?php

function loadFolderFilename($dir){
    if(empty($dir)) return false;
    $files = array();
    if($handle = opendir($dir)){
        while(false !== ($file = readdir($handle))){
            if($file === '.' || $file === '..') continue;
            $files[] = $file;
        }
        closedir($handle);
    }else{
        return false;
    }
    return $files;
}

// opendir & load file name
$markdownFiles = array(
    "../database/opensource/repo/KCORES-FlexibleLOM-Adapter/README.md",
    "../database/opensource/repo/OCP2PCIe/README.md",
);

$outputs = array(
    "../generated/KCORES-FlexibleLOM-Adapter.html",
    "../generated/OCP2PCIe.html",
);

// generate
foreach ($markdownFiles as $serial => $file) {
    print("building: {$file}\n");
    exec("node main.js {$file} {$outputs[$serial]}");
}


print("done.\n");