// const.go
package bass

/*
#include "bass.h"
*/
import "C"

const (
	// BASSVERSION as defined in bass.h:41
	BASSVERSION = 0x204
	// BASSVERSIONTEXT as defined in bass.h:42
	BASSVERSIONTEXT = "2.4"
	// BASS_OK as defined in bass.h:61
	OK = Error(C.BASS_OK)
	// BASS_ERROR_MEM as defined in bass.h:62
	ERROR_MEM = Error(C.BASS_ERROR_MEM)
	// BASS_ERROR_FILEOPEN as defined in bass.h:63
	ERROR_FILEOPEN = Error(C.BASS_ERROR_FILEOPEN)
	// BASS_ERROR_DRIVER as defined in bass.h:64
	ERROR_DRIVER = Error(C.BASS_ERROR_DRIVER)
	// BASS_ERROR_BUFLOST as defined in bass.h:65
	ERROR_BUFLOST = Error(C.BASS_ERROR_BUFLOST)
	// BASS_ERROR_HANDLE as defined in bass.h:66
	ERROR_HANDLE = Error(C.BASS_ERROR_HANDLE)
	// BASS_ERROR_FORMAT as defined in bass.h:67
	ERROR_FORMAT = Error(C.BASS_ERROR_FORMAT)
	// BASS_ERROR_POSITION as defined in bass.h:68
	ERROR_POSITION = Error(C.BASS_ERROR_POSITION)
	// BASS_ERROR_INIT as defined in bass.h:69
	ERROR_INIT = Error(C.BASS_ERROR_INIT)
	// BASS_ERROR_START as defined in bass.h:70
	ERROR_START = Error(C.BASS_ERROR_START)
	// BASS_ERROR_SSL as defined in bass.h:71
	ERROR_SSL = Error(C.BASS_ERROR_SSL)
	// BASS_ERROR_ALREADY as defined in bass.h:72
	ERROR_ALREADY = Error(C.BASS_ERROR_ALREADY)
	// BASS_ERROR_NOCHAN as defined in bass.h:73
	ERROR_NOTAUDIO=Error(C.BASS_ERROR_NOTAUDIO)
	ERROR_NOCHAN = Error(C.BASS_ERROR_NOCHAN)
	// BASS_ERROR_ILLTYPE as defined in bass.h:74)
	ERROR_ILLTYPE = Error(C.BASS_ERROR_ILLTYPE)
	// BASS_ERROR_ILLPARAM as defined in bass.h:75
	ERROR_ILLPARAM = Error(C.BASS_ERROR_ILLPARAM)
	// BASS_ERROR_NO3D as defined in bass.h:76
	ERROR_NO3D = Error(C.BASS_ERROR_NO3D)
	// BASS_ERROR_NOEAX as defined in bass.h:77
	ERROR_NOEAX = Error(C.BASS_ERROR_NOEAX)
	// BASS_ERROR_DEVICE as defined in bass.h:78
	ERROR_DEVICE = Error(C.BASS_ERROR_DEVICE)
	// ERROR_NOPLAY as defined in bass.h:79
	ERROR_NOPLAY = Error(C.BASS_ERROR_NOPLAY)
	// BASS_ERROR_FREQ as defined in bass.h:80
	ERROR_FREQ = Error(C.BASS_ERROR_FREQ)
	// BASS_ERROR_NOTFILE as defined in bass.h:81
	ERROR_NOTFILE = Error(C.BASS_ERROR_NOTFILE)
	// BASS_ERROR_NOHW as defined in bass.h:82
	ERROR_NOHW = Error(C.BASS_ERROR_NOHW)
	// BASS_ERROR_EMPTY as defined in bass.h:83
	ERROR_EMPTY = Error(C.BASS_ERROR_EMPTY)
	// BASS_ERROR_NONET as defined in bass.h:84
	ERROR_NONET = Error(C.BASS_ERROR_NONET)
	// BASS_ERROR_CREATE as defined in bass.h:85
	ERROR_CREATE = Error(C.BASS_ERROR_CREATE)
	// BASS_ERROR_NOFX as defined in bass.h:86
	ERROR_NOFX = Error(C.BASS_ERROR_NOFX)
	// BASS_ERROR_NOTAVAIL as defined in bass.h:87
	ERROR_NOTAVAIL = Error(C.BASS_ERROR_NOTAVAIL)
	// BASS_ERROR_DECODE as defined in bass.h:88
	ERROR_DECODE = Error(C.BASS_ERROR_DECODE)
	// BASS_ERROR_DX as defined in bass.h:89
	ERROR_DX = Error(C.BASS_ERROR_DX)
	// BASS_ERROR_TIMEOUT as defined in bass.h:90
	ERROR_TIMEOUT = Error(C.BASS_ERROR_TIMEOUT)
	// BASS_ERROR_FILEFORM as defined in bass.h:91
	ERROR_FILEFORM = Error(C.BASS_ERROR_FILEFORM)
	// BASS_ERROR_SPEAKER as defined in bass.h:92
	ERROR_SPEAKER = Error(C.BASS_ERROR_SPEAKER)
	// BASS_ERROR_VERSION as defined in bass.h:93
	ERROR_VERSION = Error(C.BASS_ERROR_VERSION)
	// BASS_ERROR_CODEC as defined in bass.h:94
	ERROR_CODEC = Error(C.BASS_ERROR_CODEC)
	// BASS_ERROR_ENDED as defined in bass.h:95
	ERROR_ENDED = Error(C.BASS_ERROR_ENDED)
	// BASS_ERROR_BUSY as defined in bass.h:96
	ERROR_BUSY = Error(C.BASS_ERROR_BUSY)
	ERROR_UNSTREAMABLE = Error(C.BASS_ERROR_UNSTREAMABLE)
	// BASS_ERROR_UNKNOWN as defined in bass.h:97
	ERROR_UNKNOWN = Error(C.BASS_ERROR_UNKNOWN)
	// BASS_CONFIG_BUFFER as defined in bass.h:100
	CONFIG_BUFFER = 0
	// BASS_CONFIG_UPDATEPERIOD as defined in bass.h:101
	CONFIG_UPDATEPERIOD = 1
	// BASS_CONFIG_GVOL_SAMPLE as defined in bass.h:102
	CONFIG_GVOL_SAMPLE = 4
	// BASS_CONFIG_GVOL_STREAM as defined in bass.h:103
	CONFIG_GVOL_STREAM = 5
	// BASS_CONFIG_GVOL_MUSIC as defined in bass.h:104
	CONFIG_GVOL_MUSIC = 6
	// BASS_CONFIG_CURVE_VOL as defined in bass.h:105
	CONFIG_CURVE_VOL = 7
	// BASS_CONFIG_CURVE_PAN as defined in bass.h:106
	CONFIG_CURVE_PAN = 8
	// BASS_CONFIG_FLOATDSP as defined in bass.h:107
	CONFIG_FLOATDSP = 9
	// BASS_CONFIG_3DALGORITHM as defined in bass.h:108
	CONFIG_3DALGORITHM = 10
	// BASS_CONFIG_NET_TIMEOUT as defined in bass.h:109
	CONFIG_NET_TIMEOUT = 11
	// BASS_CONFIG_NET_BUFFER as defined in bass.h:110
	CONFIG_NET_BUFFER = 12
	// BASS_CONFIG_PAUSE_NOPLAY as defined in bass.h:111
	CONFIG_PAUSE_NOPLAY = 13
	// CONFIG_NET_PREBUF as defined in bass.h:112
	CONFIG_NET_PREBUF = 15
	// BASS_CONFIG_NET_PASSIVE as defined in bass.h:113
	CONFIG_NET_PASSIVE = 18
	// BASS_CONFIG_REBASS_BUFFER as defined in bass.h:114
	CONFIG_REBASS_BUFFER = 19
	// BASS_CONFIG_NET_PLAYLIST as defined in bass.h:115
	CONFIG_NET_PLAYLIST = 21
	// BASS_CONFIG_MUSIBASS_VIRTUAL as defined in bass.h:116
	CONFIG_MUSIBASS_VIRTUAL = 22
	// BASS_CONFIG_VERIFY as defined in bass.h:117
	CONFIG_VERIFY = 23
	// BASS_CONFIG_UPDATETHREADS as defined in bass.h:118
	CONFIG_UPDATETHREADS = 24
	// BASS_CONFIG_DEV_BUFFER as defined in bass.h:119
	CONFIG_DEV_BUFFER = 27
	// BASS_CONFIG_VISTA_TRUEPOS as defined in bass.h:120
	CONFIG_VISTA_TRUEPOS = 30
	// BASS_CONFIG_IOS_MIXAUDIO as defined in bass.h:121
	CONFIG_IOS_MIXAUDIO = 34
	// BASS_CONFIG_DEV_DEFAULT as defined in bass.h:122
	CONFIG_DEV_DEFAULT = 36
	// BASS_CONFIG_NET_READTIMEOUT as defined in bass.h:123
	CONFIG_NET_READTIMEOUT = 37
	// BASS_CONFIG_VISTA_SPEAKERS as defined in bass.h:124
	CONFIG_VISTA_SPEAKERS = 38
	// BASS_CONFIG_IOS_SPEAKER as defined in bass.h:125
	CONFIG_IOS_SPEAKER = 39
	// BASS_CONFIG_MF_DISABLE as defined in bass.h:126
	CONFIG_MF_DISABLE = 40
	// BASS_CONFIG_HANDLES as defined in bass.h:127
	CONFIG_HANDLES = 41
	// BASS_CONFIG_UNICODE as defined in bass.h:128
	CONFIG_UNICODE = 42
	// BASS_CONFIG_SRC as defined in bass.h:129
	CONFIG_SRC = 43
	// BASS_CONFIG_SRBASS_SAMPLE as defined in bass.h:130
	CONFIG_SRBASS_SAMPLE = 44
	// BASS_CONFIG_ASYNCFILE_BUFFER as defined in bass.h:131
	CONFIG_ASYNCFILE_BUFFER = 45
	// BASS_CONFIG_OGG_PRESCAN as defined in bass.h:132
	CONFIG_OGG_PRESCAN = 47
	// BASS_CONFIG_MF_VIDEO as defined in bass.h:133
	CONFIG_MF_VIDEO = 48
	// BASS_CONFIG_AIRPLAY as defined in bass.h:134
	CONFIG_AIRPLAY = 49
	// BASS_CONFIG_DEV_NONSTOP as defined in bass.h:135
	CONFIG_DEV_NONSTOP = 50
	// BASS_CONFIG_IOS_NOCATEGORY as defined in bass.h:136
	CONFIG_IOS_NOCATEGORY = 51
	// BASS_CONFIG_VERIFY_NET as defined in bass.h:137
	CONFIG_VERIFY_NET = 52
	// BASS_CONFIG_DEV_PERIOD as defined in bass.h:138
	CONFIG_DEV_PERIOD = 53
	// BASS_CONFIG_FLOAT as defined in bass.h:139
	CONFIG_FLOAT = 54
	// BASS_CONFIG_NET_SEEK as defined in bass.h:140
	CONFIG_NET_SEEK = 56
	// BASS_CONFIG_NET_AGENT as defined in bass.h:143
	CONFIG_NET_AGENT = 16
	// BASS_CONFIG_NET_PROXY as defined in bass.h:144
	CONFIG_NET_PROXY = 17
	// BASS_CONFIG_IOS_NOTIFY as defined in bass.h:145
	CONFIG_IOS_NOTIFY = 46
	// DEVICE_8BITS as defined in bass.h:148
	DEVICE_8BITS = 1
	// DEVICE_MONO as defined in bass.h:149
	DEVICE_MONO = 2
	// DEVICE_3D as defined in bass.h:150
	DEVICE_3D = 4
	// DEVICE_16BITS as defined in bass.h:151
	DEVICE_16BITS = 8
	// DEVICE_LATENCY as defined in bass.h:152
	DEVICE_LATENCY = 0x100
	// DEVICE_CPSPEAKERS as defined in bass.h:153
	DEVICE_CPSPEAKERS = 0x400
	// DEVICE_SPEAKERS as defined in bass.h:154
	DEVICE_SPEAKERS = 0x800
	// DEVICE_NOSPEAKER as defined in bass.h:155
	DEVICE_NOSPEAKER = 0x1000
	// DEVICE_DMIX as defined in bass.h:156
	DEVICE_DMIX = 0x2000
	// DEVICE_FREQ as defined in bass.h:157
	DEVICE_FREQ = 0x4000
	// DEVICE_STEREO as defined in bass.h:158
	DEVICE_STEREO = 0x8000
	// BASS_OBJECT_DS as defined in bass.h:161
	BASS_OBJECT_DS = 1
	// BASS_OBJECT_DS3DL as defined in bass.h:162
	BASS_OBJECT_DS3DL = 2
	// DEVICE_ENABLED as defined in bass.h:177
	DEVICE_ENABLED = 1
	// DEVICE_DEFAULT as defined in bass.h:178
	DEVICE_DEFAULT = 2
	// DEVICE_INIT as defined in bass.h:179
	DEVICE_INIT = 4
	// DEVICE_TYPE_MASK as defined in bass.h:181
	DEVICE_TYPE_MASK = 0xff000000
	// DEVICE_TYPE_NETWORK as defined in bass.h:182
	DEVICE_TYPE_NETWORK = 0x01000000
	// DEVICE_TYPE_SPEAKERS as defined in bass.h:183
	DEVICE_TYPE_SPEAKERS = 0x02000000
	// DEVICE_TYPE_LINE as defined in bass.h:184
	DEVICE_TYPE_LINE = 0x03000000
	// DEVICE_TYPE_HEADPHONES as defined in bass.h:185
	DEVICE_TYPE_HEADPHONES = 0x04000000
	// DEVICE_TYPE_MICROPHONE as defined in bass.h:186
	DEVICE_TYPE_MICROPHONE = 0x05000000
	// DEVICE_TYPE_HEADSET as defined in bass.h:187
	DEVICE_TYPE_HEADSET = 0x06000000
	// DEVICE_TYPE_HANDSET as defined in bass.h:188
	DEVICE_TYPE_HANDSET = 0x07000000
	// DEVICE_TYPE_DIGITAL as defined in bass.h:189
	DEVICE_TYPE_DIGITAL = 0x08000000
	// DEVICE_TYPE_SPDIF as defined in bass.h:190
	DEVICE_TYPE_SPDIF = 0x09000000
	// DEVICE_TYPE_HDMI as defined in bass.h:191
	DEVICE_TYPE_HDMI = 0x0a000000
	// DEVICE_TYPE_DISPLAYPORT as defined in bass.h:192
	DEVICE_TYPE_DISPLAYPORT = 0x40000000
	// DEVICES_AIRPLAY as defined in bass.h:195
	DEVICES_AIRPLAY = 0x1000000
	// DSCAPS_CONTINUOUSRATE as defined in bass.h:215
	DSCAPS_CONTINUOUSRATE = 0x00000010
	// DSCAPS_EMULDRIVER as defined in bass.h:216
	DSCAPS_EMULDRIVER = 0x00000020
	// DSCAPS_CERTIFIED as defined in bass.h:217
	DSCAPS_CERTIFIED = 0x00000040
	// DSCAPS_SECONDARYMONO as defined in bass.h:218
	DSCAPS_SECONDARYMONO = 0x00000100
	// DSCAPS_SECONDARYSTEREO as defined in bass.h:219
	DSCAPS_SECONDARYSTEREO = 0x00000200
	// DSCAPS_SECONDARY8BIT as defined in bass.h:220
	DSCAPS_SECONDARY8BIT = 0x00000400
	// DSCAPS_SECONDARY16BIT as defined in bass.h:221
	DSCAPS_SECONDARY16BIT = 0x00000800
	// WAVE_FORMAT_1M08 as defined in bass.h:238
	WAVE_FORMAT_1M08 = 0x00000001
	// WAVE_FORMAT_1S08 as defined in bass.h:239
	WAVE_FORMAT_1S08 = 0x00000002
	// WAVE_FORMAT_1M16 as defined in bass.h:240
	WAVE_FORMAT_1M16 = 0x00000004
	// WAVE_FORMAT_1S16 as defined in bass.h:241
	WAVE_FORMAT_1S16 = 0x00000008
	// WAVE_FORMAT_2M08 as defined in bass.h:242
	WAVE_FORMAT_2M08 = 0x00000010
	// WAVE_FORMAT_2S08 as defined in bass.h:243
	WAVE_FORMAT_2S08 = 0x00000020
	// WAVE_FORMAT_2M16 as defined in bass.h:244
	WAVE_FORMAT_2M16 = 0x00000040
	// WAVE_FORMAT_2S16 as defined in bass.h:245
	WAVE_FORMAT_2S16 = 0x00000080
	// WAVE_FORMAT_4M08 as defined in bass.h:246
	WAVE_FORMAT_4M08 = 0x00000100
	// WAVE_FORMAT_4S08 as defined in bass.h:247
	WAVE_FORMAT_4S08 = 0x00000200
	// WAVE_FORMAT_4M16 as defined in bass.h:248
	WAVE_FORMAT_4M16 = 0x00000400
	// WAVE_FORMAT_4S16 as defined in bass.h:249
	WAVE_FORMAT_4S16 = 0x00000800
	// BASS_SAMPLE_8BITS as defined in bass.h:273
	SAMPLE_8BITS = 1
	// BASS_SAMPLE_FLOAT as defined in bass.h:274
	SAMPLE_FLOAT = 256
	// BASS_SAMPLE_MONO as defined in bass.h:275
	SAMPLE_MONO = 2


	// BASS_SAMPLE_3D as defined in bass.h:277
	SAMPLE_3D = 8
	// BASS_SAMPLE_SOFTWARE as defined in bass.h:278
	SAMPLE_SOFTWARE = 16
	// BASS_SAMPLE_MUTEMAX as defined in bass.h:279
	SAMPLE_MUTEMAX = 32
	// BASS_SAMPLE_VAM as defined in bass.h:280
	SAMPLE_VAM = 64
	// BASS_SAMPLE_FX as defined in bass.h:281
	SAMPLE_FX = 128
	// BASS_SAMPLE_OVER_POS as defined in bass.h:283
	// BASS_STREAM_PRESCAN as defined in bass.h:286
	STREAM_PRESCAN = 0x20000
	// BASS_MP3_SETPOS as defined in bass.h:287
	MP3_SETPOS = C.BASS_MP3_SETPOS
	// BASS_STREAM_AUTOFREE as defined in bass.h:288
	STREAM_AUTOFREE = 0x40000
	// BASS_STREAM_RESTRATE as defined in bass.h:289
	STREAM_RESTRATE = 0x80000
	// BASS_STREAM_BLOCK as defined in bass.h:290
	STREAM_BLOCK = 0x100000
	// BASS_STREAM_DECODE as defined in bass.h:291
	STREAM_DECODE = 0x200000
	// BASS_STREAM_STATUS as defined in bass.h:292
	STREAM_STATUS = 0x800000
	// BASS_SPEAKER_FRONT as defined in bass.h:318
	BASS_SPEAKER_FRONT = 0x1000000
	// BASS_SPEAKER_REAR as defined in bass.h:319
	BASS_SPEAKER_REAR = 0x2000000
	// BASS_SPEAKER_CENLFE as defined in bass.h:320
	BASS_SPEAKER_CENLFE = 0x3000000
	// BASS_SPEAKER_REAR2 as defined in bass.h:321
	BASS_SPEAKER_REAR2 = 0x4000000
	// BASS_SPEAKER_LEFT as defined in bass.h:323
	BASS_SPEAKER_LEFT = 0x10000000
	// BASS_SPEAKER_RIGHT as defined in bass.h:324
	BASS_SPEAKER_RIGHT = 0x20000000
	// BASS_SPEAKER_FRONTLEFT as defined in bass.h:325
	BASS_SPEAKER_FRONTLEFT = BASS_SPEAKER_FRONT | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_FRONTRIGHT as defined in bass.h:326
	BASS_SPEAKER_FRONTRIGHT = BASS_SPEAKER_FRONT | BASS_SPEAKER_RIGHT
	// BASS_SPEAKER_REARLEFT as defined in bass.h:327
	BASS_SPEAKER_REARLEFT = BASS_SPEAKER_REAR | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_REARRIGHT as defined in bass.h:328
	BASS_SPEAKER_REARRIGHT = BASS_SPEAKER_REAR | BASS_SPEAKER_RIGHT
	// BASS_SPEAKER_CENTER as defined in bass.h:329
	BASS_SPEAKER_CENTER = BASS_SPEAKER_CENLFE | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_LFE as defined in bass.h:330
	BASS_SPEAKER_LFE = BASS_SPEAKER_CENLFE | BASS_SPEAKER_RIGHT
	// BASS_SPEAKER_REAR2LEFT as defined in bass.h:331
	BASS_SPEAKER_REAR2LEFT = BASS_SPEAKER_REAR2 | BASS_SPEAKER_LEFT
	// BASS_SPEAKER_REAR2RIGHT as defined in bass.h:332
	BASS_SPEAKER_REAR2RIGHT = BASS_SPEAKER_REAR2 | BASS_SPEAKER_RIGHT
	// BASS_ASYNCFILE as defined in bass.h:334
	BASS_ASYNCFILE = 0x40000000
	// BASS_UNICODE as defined in bass.h:335
	BASS_UNICODE = 0x80000000
	// BASS_RECORD_PAUSE as defined in bass.h:337
	BASS_RECORD_PAUSE = 0x8000
	// BASS_RECORD_ECHOCANCEL as defined in bass.h:338
	BASS_RECORD_ECHOCANCEL = 0x2000
	// BASS_RECORD_AGC as defined in bass.h:339
	BASS_RECORD_AGC = 0x4000
	// BASS_VAM_HARDWARE as defined in bass.h:342
	BASS_VAM_HARDWARE = 1
	// BASS_VAM_SOFTWARE as defined in bass.h:343
	BASS_VAM_SOFTWARE = 2
	// BASS_VAM_TERM_TIME as defined in bass.h:344
	BASS_VAM_TERM_TIME = 4
	// BASS_VAM_TERM_DIST as defined in bass.h:345
	BASS_VAM_TERM_DIST = 8
	// BASS_VAM_TERM_PRIO as defined in bass.h:346
	BASS_VAM_TERM_PRIO = 16
	// BASS_CTYPE_SAMPLE as defined in bass.h:361
	BASS_CTYPE_SAMPLE = 1
	// BASS_CTYPE_RECORD as defined in bass.h:362
	BASS_CTYPE_RECORD = 2
	// BASS_CTYPE_STREAM as defined in bass.h:363
	BASS_CTYPE_STREAM = 0x10000
	// BASS_CTYPE_STREAM_OGG as defined in bass.h:364
	BASS_CTYPE_STREAM_OGG = 0x10002
	// BASS_CTYPE_STREAM_MP1 as defined in bass.h:365
	BASS_CTYPE_STREAM_MP1 = 0x10003
	// BASS_CTYPE_STREAM_MP2 as defined in bass.h:366
	BASS_CTYPE_STREAM_MP2 = 0x10004
	// BASS_CTYPE_STREAM_MP3 as defined in bass.h:367
	BASS_CTYPE_STREAM_MP3 = 0x10005
	// BASS_CTYPE_STREAM_AIFF as defined in bass.h:368
	BASS_CTYPE_STREAM_AIFF = 0x10006
	// BASS_CTYPE_STREAM_CA as defined in bass.h:369
	BASS_CTYPE_STREAM_CA = 0x10007
	// BASS_CTYPE_STREAM_MF as defined in bass.h:370
	BASS_CTYPE_STREAM_MF = 0x10008
	// BASS_CTYPE_STREAM_WAV as defined in bass.h:371
	BASS_CTYPE_STREAM_WAV = 0x40000
	// BASS_CTYPE_STREAM_WAV_PCM as defined in bass.h:372
	BASS_CTYPE_STREAM_WAV_PCM = 0x50001
	// BASS_CTYPE_STREAM_WAV_FLOAT as defined in bass.h:373
	BASS_CTYPE_STREAM_WAV_FLOAT = 0x50003
	// BASS_CTYPE_MUSIBASS_MOD as defined in bass.h:374
	BASS_CTYPE_MUSIBASS_MOD = 0x20000
	// BASS_CTYPE_MUSIBASS_MTM as defined in bass.h:375
	BASS_CTYPE_MUSIBASS_MTM = 0x20001
	// BASS_CTYPE_MUSIBASS_S3M as defined in bass.h:376
	BASS_CTYPE_MUSIBASS_S3M = 0x20002
	// BASS_CTYPE_MUSIBASS_XM as defined in bass.h:377
	BASS_CTYPE_MUSIBASS_XM = 0x20003
	// BASS_CTYPE_MUSIBASS_IT as defined in bass.h:378
	BASS_CTYPE_MUSIBASS_IT = 0x20004
	// BASS_CTYPE_MUSIBASS_MO3 as defined in bass.h:379
	BASS_CTYPE_MUSIBASS_MO3 = 0x00100
	// BASS_3DMODE_NORMAL as defined in bass.h:410
	BASS_3DMODE_NORMAL = 0
	// BASS_3DMODE_RELATIVE as defined in bass.h:411
	BASS_3DMODE_RELATIVE = 1
	// BASS_3DMODE_OFF as defined in bass.h:412
	BASS_3DMODE_OFF = 2
	// BASS_3DALG_DEFAULT as defined in bass.h:415
	BASS_3DALG_DEFAULT = 0
	// BASS_3DALG_OFF as defined in bass.h:416
	BASS_3DALG_OFF = 1
	// BASS_3DALG_FULL as defined in bass.h:417
	BASS_3DALG_FULL = 2
	// BASS_3DALG_LIGHT as defined in bass.h:418
	BASS_3DALG_LIGHT = 3
	// BASS_STREAMPROBASS_END as defined in bass.h:491
	BASS_STREAMPROBASS_END = 0x80000000
	// STREAMFILE_NOBUFFER as defined in bass.h:498
	STREAMFILE_NOBUFFER = 0
	// STREAMFILE_BUFFER as defined in bass.h:499
	STREAMFILE_BUFFER = 1
	// STREAMFILE_BUFFERPUSH as defined in bass.h:500
	STREAMFILE_BUFFERPUSH = 2
	// BASS_FILEDATA_END as defined in bass.h:516
	BASS_FILEDATA_END = 0
	// BASS_FILEPOS_CURRENT as defined in bass.h:519
	BASS_FILEPOS_CURRENT = 0
	// BASS_FILEPOS_DECODE as defined in bass.h:520
	BASS_FILEPOS_DECODE = BASS_FILEPOS_CURRENT
	// BASS_FILEPOS_DOWNLOAD as defined in bass.h:521
	BASS_FILEPOS_DOWNLOAD = 1
	// BASS_FILEPOS_END as defined in bass.h:522
	BASS_FILEPOS_END = 2
	// BASS_FILEPOS_START as defined in bass.h:523
	BASS_FILEPOS_START = 3
	// BASS_FILEPOS_CONNECTED as defined in bass.h:524
	BASS_FILEPOS_CONNECTED = 4
	// BASS_FILEPOS_BUFFER as defined in bass.h:525
	BASS_FILEPOS_BUFFER = 5
	// BASS_FILEPOS_SOCKET as defined in bass.h:526
	BASS_FILEPOS_SOCKET = 6
	// BASS_FILEPOS_ASYNCBUF as defined in bass.h:527
	BASS_FILEPOS_ASYNCBUF = 7
	// BASS_FILEPOS_SIZE as defined in bass.h:528
	BASS_FILEPOS_SIZE = 8
	// BASS_SYNBASS_POS as defined in bass.h:537
	BASS_SYNBASS_POS = 0
	// BASS_SYNBASS_END as defined in bass.h:538
	BASS_SYNBASS_END = 2
	// BASS_SYNBASS_META as defined in bass.h:539
	BASS_SYNBASS_META = 4
	// BASS_SYNBASS_SLIDE as defined in bass.h:540
	BASS_SYNBASS_SLIDE = 5
	// BASS_SYNBASS_STALL as defined in bass.h:541
	BASS_SYNBASS_STALL = 6
	// BASS_SYNBASS_DOWNLOAD as defined in bass.h:542
	BASS_SYNBASS_DOWNLOAD = 7
	// BASS_SYNBASS_FREE as defined in bass.h:543
	BASS_SYNBASS_FREE = 8
	// BASS_SYNBASS_SETPOS as defined in bass.h:544
	BASS_SYNBASS_SETPOS = 11
	// BASS_SYNBASS_MUSICPOS as defined in bass.h:545
	BASS_SYNBASS_MUSICPOS = 10
	// BASS_SYNBASS_MUSICINST as defined in bass.h:546
	BASS_SYNBASS_MUSICINST = 1
	// BASS_SYNBASS_MUSICFX as defined in bass.h:547
	BASS_SYNBASS_MUSICFX = 3
	// BASS_SYNBASS_OGG_CHANGE as defined in bass.h:548
	BASS_SYNBASS_OGG_CHANGE = 12
	// BASS_SYNBASS_MIXTIME as defined in bass.h:549
	BASS_SYNBASS_MIXTIME = 0x40000000
	// BASS_SYNBASS_ONETIME as defined in bass.h:550
	BASS_SYNBASS_ONETIME = 0x80000000
	// BASS_ACTIVE_STOPPED as defined in bass.h:581
	ACTIVE_STOPPED = 0
	// BASS_ACTIVE_PLAYING as defined in bass.h:582
	ACTIVE_PLAYING = 1
	// BASS_ACTIVE_STALLED as defined in bass.h:583
	ACTIVE_STALLED = 2
	// BASS_ACTIVE_PAUSED as defined in bass.h:584
	ACTIVE_PAUSED = 3
	// BASS_ATTRIB_FREQ as defined in bass.h:587
	ATTRIB_FREQ = 1
	// BASS_ATTRIB_VOL as defined in bass.h:588
	ATTRIB_VOL = 2
	// BASS_ATTRIB_PAN as defined in bass.h:589
	ATTRIB_PAN = 3
	// BASS_ATTRIB_EAXMIX as defined in bass.h:590
	ATTRIB_EAXMIX = 4
	// BASS_ATTRIB_NOBUFFER as defined in bass.h:591
	ATTRIB_NOBUFFER = 5
	// BASS_ATTRIB_VBR as defined in bass.h:592
	ATTRIB_VBR = 6
	// BASS_ATTRIB_CPU as defined in bass.h:593
	ATTRIB_CPU = 7
	// BASS_ATTRIB_SRC as defined in bass.h:594
	ATTRIB_SRC = 8
	// BASS_ATTRIB_NET_RESUME as defined in bass.h:595
	ATTRIB_NET_RESUME = 9
	// BASS_ATTRIB_SCANINFO as defined in bass.h:596
	ATTRIB_SCANINFO = 10
	// BASS_ATTRIB_NORAMP as defined in bass.h:597
	ATTRIB_NORAMP = 11
	// BASS_ATTRIB_BITRATE as defined in bass.h:598
	ATTRIB_BITRATE = 12
	// BASS_ATTRIB_MUSIBASS_AMPLIFY as defined in bass.h:599
	ATTRIB_MUSIBASS_AMPLIFY = 0x100
	// BASS_ATTRIB_MUSIBASS_PANSEP as defined in bass.h:600
	ATTRIB_MUSIBASS_PANSEP = 0x101
	// BASS_ATTRIB_MUSIBASS_PSCALER as defined in bass.h:601
	ATTRIB_MUSIBASS_PSCALER = 0x102
	// BASS_ATTRIB_MUSIBASS_BPM as defined in bass.h:602
	ATTRIB_MUSIBASS_BPM = 0x103
	// BASS_ATTRIB_MUSIBASS_SPEED as defined in bass.h:603
	ATTRIB_MUSIBASS_SPEED = 0x104
	// BASS_ATTRIB_MUSIBASS_VOL_GLOBAL as defined in bass.h:604
	ATTRIB_MUSIBASS_VOL_GLOBAL = 0x105
	// BASS_ATTRIB_MUSIBASS_ACTIVE as defined in bass.h:605
	ATTRIB_MUSIBASS_ACTIVE = 0x106
	// BASS_ATTRIB_MUSIBASS_VOL_CHAN as defined in bass.h:606
	ATTRIB_MUSIBASS_VOL_CHAN = 0x200
	// BASS_ATTRIB_MUSIBASS_VOL_INST as defined in bass.h:607
	ATTRIB_MUSIBASS_VOL_INST = 0x300
	// BASS_DATA_AVAILABLE as defined in bass.h:610
	BASS_DATA_AVAILABLE = 0
	// BASS_DATA_FIXED as defined in bass.h:611
	BASS_DATA_FIXED = 0x20000000
	// BASS_DATA_FLOAT as defined in bass.h:612
	BASS_DATA_FLOAT = 0x40000000
	// BASS_DATA_FFT256 as defined in bass.h:613
	BASS_DATA_FFT256 = 0x80000000
	// BASS_DATA_FFT512 as defined in bass.h:614
	BASS_DATA_FFT512 = 0x80000001
	// BASS_DATA_FFT1024 as defined in bass.h:615
	BASS_DATA_FFT1024 = 0x80000002
	// BASS_DATA_FFT2048 as defined in bass.h:616
	BASS_DATA_FFT2048 = 0x80000003
	// BASS_DATA_FFT4096 as defined in bass.h:617
	BASS_DATA_FFT4096 = 0x80000004
	// BASS_DATA_FFT8192 as defined in bass.h:618
	BASS_DATA_FFT8192 = 0x80000005
	// BASS_DATA_FFT16384 as defined in bass.h:619
	BASS_DATA_FFT16384 = 0x80000006
	// BASS_DATA_FFT32768 as defined in bass.h:620
	BASS_DATA_FFT32768 = 0x80000007
	// BASS_DATA_FFT_INDIVIDUAL as defined in bass.h:621
	BASS_DATA_FFT_INDIVIDUAL = 0x10
	// BASS_DATA_FFT_NOWINDOW as defined in bass.h:622
	BASS_DATA_FFT_NOWINDOW = 0x20
	// BASS_DATA_FFT_REMOVEDC as defined in bass.h:623
	BASS_DATA_FFT_REMOVEDC = 0x40
	// BASS_DATA_FFT_COMPLEX as defined in bass.h:624
	BASS_DATA_FFT_COMPLEX = 0x80
	// BASS_LEVEL_MONO as defined in bass.h:627
	BASS_LEVEL_MONO = 1
	// BASS_LEVEL_STEREO as defined in bass.h:628
	BASS_LEVEL_STEREO = 2
	// BASS_LEVEL_RMS as defined in bass.h:629
	BASS_LEVEL_RMS = 4
	// BASS_TAG_ID3 as defined in bass.h:632
	BASS_TAG_ID3 = 0
	// BASS_TAG_ID3V2 as defined in bass.h:633
	BASS_TAG_ID3V2 = 1
	// BASS_TAG_OGG as defined in bass.h:634
	BASS_TAG_OGG = 2
	// BASS_TAG_HTTP as defined in bass.h:635
	BASS_TAG_HTTP = 3
	// BASS_TAG_ICY as defined in bass.h:636
	BASS_TAG_ICY = 4
	// BASS_TAG_META as defined in bass.h:637
	BASS_TAG_META = 5
	// BASS_TAG_APE as defined in bass.h:638
	BASS_TAG_APE = 6
	// BASS_TAG_MP4 as defined in bass.h:639
	BASS_TAG_MP4 = 7
	// BASS_TAG_WMA as defined in bass.h:640
	BASS_TAG_WMA = 8
	// BASS_TAG_VENDOR as defined in bass.h:641
	BASS_TAG_VENDOR = 9
	// BASS_TAG_LYRICS3 as defined in bass.h:642
	BASS_TAG_LYRICS3 = 10
	// BASS_TAG_CA_CODEC as defined in bass.h:643
	BASS_TAG_CA_CODEC = 11
	// BASS_TAG_MF as defined in bass.h:644
	BASS_TAG_MF = 13
	// BASS_TAG_WAVEFORMAT as defined in bass.h:645
	BASS_TAG_WAVEFORMAT = 14
	// BASS_TAG_RIFF_INFO as defined in bass.h:646
	BASS_TAG_RIFF_INFO = 0x100
	// BASS_TAG_RIFF_BEXT as defined in bass.h:647
	BASS_TAG_RIFF_BEXT = 0x101
	// BASS_TAG_RIFF_CART as defined in bass.h:648
	BASS_TAG_RIFF_CART = 0x102
	// BASS_TAG_RIFF_DISP as defined in bass.h:649
	BASS_TAG_RIFF_DISP = 0x103
	// BASS_TAG_APE_BINARY as defined in bass.h:650
	BASS_TAG_APE_BINARY = 0x1000
	// BASS_TAG_MUSIBASS_NAME as defined in bass.h:651
	BASS_TAG_MUSIBASS_NAME = 0x10000
	// BASS_TAG_MUSIBASS_MESSAGE as defined in bass.h:652
	BASS_TAG_MUSIBASS_MESSAGE = 0x10001
	// BASS_TAG_MUSIBASS_ORDERS as defined in bass.h:653
	BASS_TAG_MUSIBASS_ORDERS = 0x10002
	// BASS_TAG_MUSIBASS_AUTH as defined in bass.h:654
	BASS_TAG_MUSIBASS_AUTH = 0x10003
	// BASS_TAG_MUSIBASS_INST as defined in bass.h:655
	BASS_TAG_MUSIBASS_INST = 0x10100
	// BASS_TAG_MUSIBASS_SAMPLE as defined in bass.h:656
	BASS_TAG_MUSIBASS_SAMPLE = 0x10300
	// BASS_POS_BYTE as defined in bass.h:767
	POS_BYTE = 0
	// BASS_POS_MUSIBASS_ORDER as defined in bass.h:768
	POS_MUSIBASS_ORDER = 1
	// BASS_POS_OGG as defined in bass.h:769
	POS_OGG = 3
	// BASS_POS_INEXACT as defined in bass.h:770
	POS_INEXACT = 0x8000000
	// BASS_POS_DECODE as defined in bass.h:771
	POS_DECODE = 0x10000000
	// BASS_POS_DECODETO as defined in bass.h:772
	POS_DECODETO = 0x20000000
	// BASS_POS_SCAN as defined in bass.h:773
	POS_SCAN = 0x40000000
	// BASS_INPUT_OFF as defined in bass.h:776
	BASS_INPUT_OFF = 0x10000
	// BASS_INPUT_ON as defined in bass.h:777
	BASS_INPUT_ON = 0x20000
	// BASS_INPUT_TYPE_MASK as defined in bass.h:779
	BASS_INPUT_TYPE_MASK = 0xff000000
	// BASS_INPUT_TYPE_UNDEF as defined in bass.h:780
	BASS_INPUT_TYPE_UNDEF = 0x00000000
	// BASS_INPUT_TYPE_DIGITAL as defined in bass.h:781
	BASS_INPUT_TYPE_DIGITAL = 0x01000000
	// BASS_INPUT_TYPE_LINE as defined in bass.h:782
	BASS_INPUT_TYPE_LINE = 0x02000000
	// BASS_INPUT_TYPE_MIC as defined in bass.h:783
	BASS_INPUT_TYPE_MIC = 0x03000000
	// BASS_INPUT_TYPE_SYNTH as defined in bass.h:784
	BASS_INPUT_TYPE_SYNTH = 0x04000000
	// BASS_INPUT_TYPE_CD as defined in bass.h:785
	BASS_INPUT_TYPE_CD = 0x05000000
	// BASS_INPUT_TYPE_PHONE as defined in bass.h:786
	BASS_INPUT_TYPE_PHONE = 0x06000000
	// BASS_INPUT_TYPE_SPEAKER as defined in bass.h:787
	BASS_INPUT_TYPE_SPEAKER = 0x07000000
	// BASS_INPUT_TYPE_WAVE as defined in bass.h:788
	BASS_INPUT_TYPE_WAVE = 0x08000000
	// BASS_INPUT_TYPE_AUX as defined in bass.h:789
	BASS_INPUT_TYPE_AUX = 0x09000000
	// BASS_INPUT_TYPE_ANALOG as defined in bass.h:790
	BASS_INPUT_TYPE_ANALOG = 0x0a000000
	// BASS_DX8_PHASE_NEG_180 as defined in bass.h:884
	BASS_DX8_PHASE_NEG_180 = 0
	// BASS_DX8_PHASE_NEG_90 as defined in bass.h:885
	BASS_DX8_PHASE_NEG_90 = 1
	// BASS_DX8_PHASE_ZERO as defined in bass.h:886
	BASS_DX8_PHASE_ZERO = 2
	// BASS_DX8_PHASE_90 as defined in bass.h:887
	BASS_DX8_PHASE_90 = 3
	// BASS_DX8_PHASE_180 as defined in bass.h:888
	BASS_DX8_PHASE_180 = 4
	// BASS_IOSNOTIFY_INTERRUPT as defined in bass.h:894
	BASS_IOSNOTIFY_INTERRUPT = 1
	// BASS_IOSNOTIFY_INTERRUPT_END as defined in bass.h:895
	BASS_IOSNOTIFY_INTERRUPT_END = 2

	// BASS_FX_DX8_CHORUS as declared in bass.h:795
	BASS_FX_DX8_CHORUS = C.BASS_FX_DX8_CHORUS
	// BASS_FX_DX8_COMPRESSOR as declared in bass.h:796
	BASS_FX_DX8_COMPRESSOR = C.BASS_FX_DX8_COMPRESSOR
	// BASS_FX_DX8_DISTORTION as declared in bass.h:797
	BASS_FX_DX8_DISTORTION = C.BASS_FX_DX8_DISTORTION
	// BASS_FX_DX8_ECHO as declared in bass.h:798
	BASS_FX_DX8_ECHO = C.BASS_FX_DX8_ECHO
	// BASS_FX_DX8_FLANGER as declared in bass.h:799
	BASS_FX_DX8_FLANGER = C.BASS_FX_DX8_FLANGER
	// BASS_FX_DX8_GARGLE as declared in bass.h:800
	BASS_FX_DX8_GARGLE = C.BASS_FX_DX8_GARGLE
	// BASS_FX_DX8_I3DL2REVERB as declared in bass.h:801
	BASS_FX_DX8_I3DL2REVERB = C.BASS_FX_DX8_I3DL2REVERB
	// BASS_FX_DX8_PARAMEQ as declared in bass.h:802
	BASS_FX_DX8_PARAMEQ = C.BASS_FX_DX8_PARAMEQ
	// BASS_FX_DX8_REVERB as declared in bass.h:803
	FX_DX8_REVERB = C.BASS_FX_DX8_REVERB
	SYNC_DEV_FAIL=C.BASS_SYNC_DEV_FAIL
	SYNC_DEV_FORMAT=C.BASS_SYNC_DEV_FORMAT
	SYNC_DOWNLOAD=C.BASS_SYNC_DOWNLOAD
	SYNC_END=C.BASS_SYNC_END
	SYNC_FREE=C.BASS_SYNC_FREE
	SYNC_META=C.BASS_SYNC_META

	SYNC_MUSICFX=C.BASS_SYNC_MUSICFX
	SYNC_MUSICINST=C.BASS_SYNC_MUSICINST
	SYNC_MUSICPOS=C.BASS_SYNC_MUSICPOS
	SYNC_OGG_CHANGE=C.BASS_SYNC_OGG_CHANGE
	SYNC_POS=C.BASS_SYNC_POS
	SYNC_SETPOS=C.BASS_SYNC_SETPOS
	SYNC_SLIDE=C.BASS_SYNC_SLIDE
	SYNC_STALL=C.BASS_SYNC_STALL
	// BASS_SampleGetChannel flags
	SamchanNew = C.BASS_SAMCHAN_NEW
	SamchanStream = C.BASS_SAMCHAN_STREAM
	SAMPLE_LOOP = C.BASS_SAMPLE_LOOP
	SAMPLE_OVER_VOL = C.BASS_SAMPLE_OVER_VOL
	SAMPLE_OVER_POS = C.BASS_SAMPLE_OVER_POS
	SAMPLE_OVER_DIST = C.BASS_SAMPLE_OVER_DIST
	CONFIG_REC_WASAPI = C.BASS_CONFIG_REC_WASAPI
	CONFIG_REC_BUFFER = C.BASS_CONFIG_REC_BUFFER
	DATA_AVAILABLE = C.BASS_DATA_AVAILABLE
)