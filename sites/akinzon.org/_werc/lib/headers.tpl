<!DOCTYPE HTML>
<html>
<head>
    <title>%($pageTitle%)</title>
    <link rel="stylesheet" href="/pub/style/style.css" type="text/css" media="screen, handheld" title="default">
<link rel="shortcut icon" href="/pub/style/aki.ico" type="image/vnd.microsoft.icon">
    <meta charset="UTF-8">
% h = `{get_lib_file headers.inc}
% if(: ~ $#h 0)
%   cat $h
    %($"extraHeaders%)
</head>
<body>

